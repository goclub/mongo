package mo

import (
	"context"
	"encoding/json"
	xerr "github.com/goclub/error"
	xtype "github.com/goclub/type"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)
type Collection struct {
	db   *Database
	Core *mongo.Collection
	name string
}

func NewCollection(db *Database, collectionName string, opts ...*options.CollectionOptions) *Collection {
	return &Collection{
		db:   db,
		Core: db.Core.Collection(collectionName, opts...),
		name: collectionName,
	}
}

type ResultInsertOne struct {
	*mongo.InsertOneResult
}

func (res ResultInsertOne) InsertedObjectID() (objectID primitive.ObjectID, err error) {
	switch id := res.InsertedID.(type) {
	case primitive.ObjectID:
		return id, nil
	default:
		err = xerr.New("goclub/mongo: ResultInsertOne{}.InsertedObjectID() id is not primitive.ObjectID")
		return
	}
}
func (c *Collection) InsertOne(ctx context.Context, document Document, cmd InsertOneCommand) (result ResultInsertOne, err error) {
	defer func() {
		if err != nil {
			err = xerr.WithStack(err)
		}
	}()
	err = document.BeforeInsert() ; if err != nil {
	    return
	}
	coreRes, err := c.Core.InsertOne(ctx, document, cmd.Options()...)
	if err != nil {
		return
	}
	result.InsertOneResult = coreRes

	err = document.AfterInsert(result)
	if err != nil {
		return
	}
	return
}

type ResultInsertMany struct {
	*mongo.InsertManyResult
}

func (res ResultInsertMany) InsertedObjectIDs() (insertedObjectIDs []primitive.ObjectID, err error) {
	for _, v := range res.InsertedIDs {
		switch id := v.(type) {
		case primitive.ObjectID:
		insertedObjectIDs = append(insertedObjectIDs, id)
		default:
			err = xerr.New("goclub/mongo: ResultInsertMany{}.InsertedObjectIDs() id is not primitive.ObjectID")
			return
		}
	}
	return
}
func (c *Collection) InsertMany(ctx context.Context, documents ManyDocument, cmd InsertManyCommand) (result ResultInsertMany, err error) {
	defer func() {
		if err != nil {
			err = xerr.WithStack(err)
		}
	}()
	data, err := documents.ManyD()
	if err != nil {
		return
	}

	coreRes, err := c.Core.InsertMany(ctx, data, cmd.Options()...)
	if err != nil {
		return
	}
	result.InsertManyResult = coreRes
	err = documents.AfterInsertMany(result)
	if err != nil {
		return
	}
	return
}

func (c *Collection) FindOneByObjectID(ctx context.Context, objectID primitive.ObjectID, document Document, cmd FindOneCommand) (has bool, err error) {
	return c.FindOne(ctx, bson.D{{"_id", objectID}},document, cmd)
}
func (c *Collection) FindOne(ctx context.Context, filter interface{}, document Document, cmd FindOneCommand) (has bool, err error) {
	defer func() {
		if err != nil {
			err = xerr.WithStack(err)
		}
	}()
	res := c.Core.FindOne(ctx, filter, cmd.Options()...)
	err = res.Err()
	has = true
	if xerr.Is(err, mongo.ErrNoDocuments) {
		has = false
		err = nil
		return
	}
	if xerr.Is(err, mongo.ErrNilDocument) {
		has = false
		err = nil
		return
	}
	if err != nil {
		return
	}
	err = res.Decode(document)
	if err != nil {
		return
	}
	return
}
func (c *Collection) FindCursor(ctx context.Context, filter interface{}, cmd FindCommand) (cursor *mongo.Cursor, err error) {
	defer func() {
		if err != nil {
			err = xerr.WithStack(err)
		}
	}()
	if cmd.DebugLookupFilter {
		filterBytes, jsonMarshalErr := json.MarshalIndent(filter, "", "  ") ; if jsonMarshalErr != nil {
			Logger.Printf("goclub/mongo:debug:\n%+v", jsonMarshalErr)
		} else {
			Logger.Print("goclub/mongo:debug lookup filter:\n", string(filterBytes))
		}
	}
	cursor, err = c.Core.Find(ctx, filter, cmd.Options()...) ; if err != nil {
		return
	}
	if cmd.DebugLookupResults {
		debugResults := []map[string]interface{}{}
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			debugItem := map[string]interface{}{}
			debugErr := cursor.Decode(&debugItem) ; if debugErr  != nil {
				Logger.Printf("goclub/mongo:debug:\n%+v", debugErr)
			}
			debugResults = append(debugResults, debugItem)
		}
		debugResultsBytes, jsonMarshalErr := json.Marshal(debugResults) ; if jsonMarshalErr != nil {
			Logger.Printf("goclub/mongo:debug:\n%+v", jsonMarshalErr)
		} else {
			Logger.Printf("goclub/mongo:debug:decode lookup results:\nlen:%d\nresults:\n%s", len(debugResults), string(debugResultsBytes))
		}
		// 如果将开启了 DebugLookupResults 的代码提交到线上,线上会以为 cursor.Next() 已经读完数据导致结果为空
		err = xerr.New("goclub/mongo:debug:decode lookup results")
		return
	}
	return
}
func (c *Collection) Find(ctx context.Context, filter interface{}, slicePtr interface{}, cmd FindCommand) (err error) {
	cursor, err := c.FindCursor(ctx, filter, cmd) ; if err != nil {
		return
	}
	return cursor.All(ctx, slicePtr)
}
func (c *Collection) UpdateOne(ctx context.Context, data FilterUpdate, cmd UpdateCommand) (updateResult *mongo.UpdateResult, err error) {
	defer func() {
		if err != nil {
			err = xerr.WithStack(err)
		}
	}()
	return c.Core.UpdateOne(ctx, data.Filter, data.Update, cmd.Options()...)
}
func (c *Collection) Aggregate(ctx context.Context, pipeline interface{}, cmd AggregateCommand, resultPtr interface{}) (err error) {
	rValue := reflect.ValueOf(resultPtr)
	rType := rValue.Type()
	if rType.Kind() != reflect.Ptr {
		return xerr.New("goclub/mongo: mo.Collection{}.Aggregate(ctx, pipeline, cmd, ptr) " + rType.String() + " must be ptr")
	}

	cursor, err := c.AggregateCursor(ctx, pipeline, cmd) ; if err != nil {
	    return
	}
	err = cursor.All(ctx, resultPtr) ; if err != nil {
	    return
	}
	return
}
func (c *Collection) AggregateCursor(ctx context.Context, pipeline interface{}, cmd AggregateCommand) (cursor *mongo.Cursor, err error) {
	defer func() {
		if err != nil {
			err = xerr.WithStack(err)
		}
	}()
	return c.Core.Aggregate(ctx, pipeline, cmd.Options()...)
}
func (c *Collection) DeleteOne(ctx context.Context, filter interface{}, cmd DeleteCommand) (result *mongo.DeleteResult, err error) {
	return c.Core.DeleteOne(ctx, filter, cmd.Options()...)
}
func (c *Collection) DeleteMany(ctx context.Context, filter interface{}, cmd DeleteCommand) (result *mongo.DeleteResult, err error) {
	return c.Core.DeleteMany(ctx, filter, cmd.Options()...)
}
func (c *Collection) Count(ctx context.Context, filter interface{}, cmd CountCommand) (total uint64, err error) {
	countResult, err := c.Core.CountDocuments(ctx, filter, cmd.Options()...) ; if err != nil {
	    return
	}
	total = uint64(countResult)
	return
}

type Paging struct {
	Filter interface{}
	FindCmd  FindCommand
	SlicePtr interface{}
	CountCmd CountCommand
	Page uint64
	PerPage uint64
}
func (c *Collection) Paging(ctx context.Context, p Paging) (total uint64, err error) {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.PerPage == 0 {
		p.PerPage = 10
	}
	p.FindCmd.Limit = xtype.Uint64(p.PerPage)
	p.FindCmd.Skip = xtype.Uint64((p.Page-1)*p.PerPage)
	err = c.Find(ctx, p.Filter, p.SlicePtr, p.FindCmd); if err != nil {
		return
	}
	countResult, err := c.Count(ctx, p.Filter, p.CountCmd) ; if err != nil {
		return
	}
	total = uint64(countResult)
	return
}
