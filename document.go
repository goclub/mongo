package mo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Document interface {
	AfterInsert(data AfterInsertData) (err error)
	BeforeInsert()(err error)
	BeforeUpdate() error
	AfterUpdate() error
}
type DefaultLifeCycle struct {

}
func (v *DefaultLifeCycle) BeforeInsert() error {return nil}
func (v *DefaultLifeCycle) AfterInsert(data AfterInsertData) error {return nil}
func (v *DefaultLifeCycle) BeforeUpdate() error {return nil}
func (v *DefaultLifeCycle) AfterUpdate() error {return nil}

type AfterInsertData struct {
	ObjectID primitive.ObjectID
}
type ManyDocument interface {
	ManyD() (documents []interface{}, err error)
	AfterInsertMany(data AfterInsertManyData) (err error)
}
type AfterInsertManyData struct {
	ObjectIDs func() []primitive.ObjectID
}
