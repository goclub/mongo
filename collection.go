package mo

import (
	"context"
	xerr "github.com/goclub/error"
	xtype "github.com/goclub/type"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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

func (res ResultInsertOne) InsertedObjectID() primitive.ObjectID {
	log.Printf("---------- %T %+v", res.InsertedID, res.InsertedID)
	return res.InsertedID.(primitive.ObjectID)
}
func (c *Collection) InsertOne(ctx context.Context, document Document, cmd InsertOneCommand) (result ResultInsertOne, err error) {
	defer func() {
		if err != nil {
			err = xerr.WithStack(err)
		}
	}()
	coreRes, err := c.Core.InsertOne(ctx, document, cmd.Options()...)
	if err != nil {
		return
	}
	result.InsertOneResult = coreRes
	err = document.BeforeInsert(BeforeInsertData{
		ObjectID: result.InsertedObjectID(),
	})
	if err != nil {
		return
	}
	return
}

type ResultInsertMany struct {
	*mongo.InsertManyResult
}

func (res ResultInsertMany) InsertedObjectIDs() (insertedObjectIDs []primitive.ObjectID) {
	for _, id := range res.InsertedIDs {
		insertedObjectIDs = append(insertedObjectIDs, id.(primitive.ObjectID))
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
	err = documents.BeforeInsertMany(BeforeInsertManyData{
		ObjectIDs: result.InsertedObjectIDs,
	})
	if err != nil {
		return
	}
	return
}

func (c *Collection) FindByObjectID(ctx context.Context, objectID primitive.ObjectID, document Document, cmd FindOneCommand) (has bool, err error) {
	defer func() {
		if err != nil {
			err = xerr.WithStack(err)
		}
	}()
	res := c.Core.FindOne(ctx, bson.D{{"_id", objectID}}, cmd.Options()...)
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
func (c *Collection) Find(ctx context.Context, filter interface{}, cmd FindCommand, resultPtr interface{}) (err error) {
	cursor, err := c.FindCursor(ctx, filter, cmd) ; if err != nil {
		return
	}
	return cursor.All(ctx, resultPtr)
}
func (c *Collection) FindCursor(ctx context.Context, filter interface{}, cmd FindCommand) (cursor *mongo.Cursor, err error) {
	defer func() {
		if err != nil {
			err = xerr.WithStack(err)
		}
	}()
	cursor, err = c.Core.Find(ctx, filter, cmd.Options()...) ; if err != nil {
		return
	}
	return
}
func (c *Collection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, cmd UpdateCommand) (updateResult *mongo.UpdateResult, err error) {
	defer func() {
		if err != nil {
			err = xerr.WithStack(err)
		}
	}()
	return c.Core.UpdateOne(ctx, filter, update, cmd.Options()...)
}
func (c *Collection) Aggregate(ctx context.Context, pipeline interface{}, cmd AggregateCommand) (cursor *mongo.Cursor, err error) {
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
	FindCmd FindCommand
	ResultPtr interface{}
	CountCmd CountCommand
	Page uint64
	PerPage uint64
}
func (c *Collection) Paging(ctx context.Context, p Paging) (total uint64, err error) {
	if p.Page == 0 {
		log.Print("goclub/mongo: mo.Collection{}.Paging(ctx, p), p.Page can not be 0")
		p.Page = 1
	}
	p.FindCmd.Limit = xtype.Uint64(p.PerPage)
	p.FindCmd.Skip = xtype.Uint64((p.Page-1)*p.PerPage)
	err = c.Find(ctx, p.Filter, p.FindCmd, p.ResultPtr) ; if err != nil {
		return
	}
	countResult, err := c.Count(ctx, p.Filter, p.CountCmd) ; if err != nil {
		return
	}
	total = uint64(countResult)
	return
}