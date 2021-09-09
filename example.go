package mo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EComment struct {
	ObjectID primitive.ObjectID
	UserID uint64
	Message string
}


func (v *EComment) BeforeInsert(data BeforeInsertData) (err error) {
	if v.ObjectID.IsZero() { v.ObjectID = data.ObjectID }
	return
}
func (v EComment) D() bson.D {
	return bson.D{
		{"userID", v.UserID},
		{"message", v.Message,},
	}
}

type ManyEComment []EComment
func (many ManyEComment) DS () (ds []interface{}) {
	for _, v := range many {
		ds = append(ds, v.D())
	}
	return
}

func (many ManyEComment) BeforeInsertMany(data BeforeInsertManyData) (err error) {
	objectIDs := data.ObjectIDs()
	for i,_ := range many {
		many[i].ObjectID = objectIDs[i]
	}
	return
}