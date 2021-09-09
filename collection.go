package mo

import (
	"context"
	xerr "github.com/goclub/error"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	db *Database
	core *mongo.Collection
	name string 
}
func NewCollection(db *Database, collectionName string, opts ...*options.CollectionOptions) *Collection {
	return &Collection{
		db: db,
		core: db.core.Collection(collectionName, opts...),
		name:   collectionName,
	}
}
type ResultInsertOne struct {
	*mongo.InsertOneResult
}
func (res ResultInsertOne) InsertedObjectID() primitive.ObjectID {
	return res.InsertedID.(primitive.ObjectID)
}
func (c Collection) InsertOne(ctx context.Context, document Document, opts ...*options.InsertOneOptions) (result ResultInsertOne, err error) {
	defer func() { if err != nil { err = xerr.WithStack(err) } }()
	coreRes, err := c.core.InsertOne(ctx, document.D(), opts...) ; if err != nil {
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
func (c Collection) InsertMany(ctx context.Context, documents DocumentMany, opts ...*options.InsertManyOptions) (result ResultInsertMany, err error) {
	defer func() { if err != nil { err = xerr.WithStack(err) } }()
	coreRes, err := c.core.InsertMany(ctx, documents.DS(), opts...) ; if err != nil {
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
