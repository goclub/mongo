package mo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Document interface {
	D() bson.D
	BeforeInsert(data BeforeInsertData) (err error)
}
type BeforeInsertData struct {
	ObjectID primitive.ObjectID
}
type DocumentMany interface {
	DS() []interface{}
	BeforeInsertMany(data BeforeInsertManyData) (err error)
}
type BeforeInsertManyData struct {
	ObjectIDs func() []primitive.ObjectID
}
