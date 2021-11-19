package mo

import (
	"context"
	xerr "github.com/goclub/error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	db *Database
	Core *mongo.Collection
	name string 
}
func NewCollection(db *Database, collectionName string, opts ...*options.CollectionOptions) *Collection {
	return &Collection{
		db: db,
		Core: db.core.Collection(collectionName, opts...),
		name:   collectionName,
	}
}
type ResultInsertOne struct {
	*mongo.InsertOneResult
}
func (res ResultInsertOne) InsertedObjectID() primitive.ObjectID {
	return res.InsertedID.(primitive.ObjectID)
}
func (c Collection) InsertOne(ctx context.Context, document Document, cmd InsertOneCommand) (result ResultInsertOne, err error) {
	defer func() { if err != nil { err = xerr.WithStack(err) } }()
	coreRes, err := c.Core.InsertOne(ctx, document, cmd.Options()...) ; if err != nil {
	    return
	}
	result.InsertOneResult = coreRes
	err = document.BeforeInsert(BeforeInsertData{
		ObjectID: result.InsertedObjectID(),
	}) ; if err != nil {
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
func (c Collection) InsertMany(ctx context.Context, documents DocumentMany, cmd InsertManyCommand) (result ResultInsertMany, err error) {
	defer func() { if err != nil { err = xerr.WithStack(err) } }()
	data, err := documents.ManyD() ; if err != nil {
	    return
	}
	coreRes, err := c.Core.InsertMany(ctx, data, cmd.Options()...) ; if err != nil {
		return
	}
	result.InsertManyResult = coreRes
	err = documents.BeforeInsertMany(BeforeInsertManyData{
		ObjectIDs: result.InsertedObjectIDs,
	}) ; if err != nil {
	    return
	}
	return
}

func (c Collection) FindByObjectID(ctx context.Context, objectID primitive.ObjectID, document Document, cmd FindOneCommand) (has bool, err error) {
	defer func() { if err != nil { err = xerr.WithStack(err) } }()
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
	err = res.Decode(document) ; if err != nil {
	    return
	}
	return
}

func (c Collection) FindOne(ctx context.Context, filter interface{}, document Document, cmd FindOneCommand) (has bool, err error) {
	defer func() { if err != nil { err = xerr.WithStack(err) } }()
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
	err = res.Decode(document) ; if err != nil {
		return
	}
	return
}