package mo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Document interface {
	BeforeInsert(data BeforeInsertData) (err error)
}
type BeforeInsertData struct {
	ObjectID primitive.ObjectID
}
type ManyDocument interface {
	ManyD() (documents []interface{}, err error)
	BeforeInsertMany(data BeforeInsertManyData) (err error)
}
type BeforeInsertManyData struct {
	ObjectIDs func() []primitive.ObjectID
}
