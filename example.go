package mo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ObjectID primitive.ObjectID `bson:"_id,omitempty"`
	UserID uint64 `bson:"userID"`
	Message string `bson:"message"`
	Like uint64 `bson:"like"`
}


func (v *Comment) BeforeInsert(data BeforeInsertData) (err error) {
	if v.ObjectID.IsZero() { v.ObjectID = data.ObjectID }
	return
}
type ManyComment []Comment
func (many ManyComment) ManyD () (documents []interface{}, err error) {
	for _, v := range many {
		var b []byte
		b, err = bson.Marshal(v) ; if err != nil {
		    return
		}
		documents = append(documents, bson.Raw(b))
	}
	return
}

func (many ManyComment) BeforeInsertMany(data BeforeInsertManyData) (err error) {
	objectIDs := data.ObjectIDs()
	for i,_ := range many {
		many[i].ObjectID = objectIDs[i]
	}
	return
}