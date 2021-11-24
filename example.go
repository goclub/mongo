package mo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ExampleComment struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	UserID uint64 `bson:"userID"`
	NewsID primitive.ObjectID `bson:"newsID"`
	Message string `bson:"message"`
	Like uint64 `bson:"like"`
	CreateTime time.Time `bson:"createTime"`
}

func (d ExampleComment) Field() (f struct {
	ID      string
	UserID  string
	NewsID string
	Message string
	Like    string
	CreateTime string
} ) {
	f.ID = "_id"
	f.UserID = "userID"
	f.NewsID = "newsID"
	f.Message = "message"
	f.Like = "like"
	f.CreateTime = "createTime"
	return
}


func (v *ExampleComment) BeforeInsert(data BeforeInsertData) (err error) {
	if v.ID.IsZero() { v.ID = data.ObjectID }
	return
}
type ManyExampleComment []ExampleComment
func (many ManyExampleComment) ManyD () (documents []interface{}, err error) {
	for _, v := range many {
		var b []byte
		b, err = bson.Marshal(v) ; if err != nil {
		    return
		}
		documents = append(documents, bson.Raw(b))
	}
	return
}

func (many ManyExampleComment) BeforeInsertMany(data BeforeInsertManyData) (err error) {
	IDs := data.ObjectIDs()
	for i,_ := range many {
		many[i].ID = IDs[i]
	}
	return
}

type ExampleNewsStatDaily struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Date string `bson:"date"`
	NewsID primitive.ObjectID `bson:"newsID"`
	UV uint64 `bson:"uv"`
	PV uint64 `bson:"pv"`
	PlatformUV map[string]uint64 `bson:"platformUV"`
}

func (v *ExampleNewsStatDaily) BeforeInsert(data BeforeInsertData) (err error) {
	if v.ID.IsZero() { v.ID = data.ObjectID }
	return
}

func (d ExampleNewsStatDaily) Field() (f struct {
	ID      string
	Date  string
	NewsID string
	UV    string
	PV string
	PlatformUV string
} ) {
	f.ID = "_id"
	f.Date = "date"
	f.NewsID = "newsID"
	f.UV = "uv"
	f.PV = "pv"
	f.PlatformUV = "platformUV"
	return
}

type ManyExampleNewsStatDaily []ExampleNewsStatDaily
func (many ManyExampleNewsStatDaily) ManyD () (documents []interface{}, err error) {
	for _, v := range many {
		var b []byte
		b, err = bson.Marshal(v) ; if err != nil {
			return
		}
		documents = append(documents, bson.Raw(b))
	}
	return
}

func (many ManyExampleNewsStatDaily) BeforeInsertMany(data BeforeInsertManyData) (err error) {
	IDs := data.ObjectIDs()
	for i,_ := range many {
		many[i].ID = IDs[i]
	}
	return
}