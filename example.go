package mo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// 演示用连接uri,正式项目请使用公司环境对于的uri,因为示例的副本集环境是docker所以需要使用connect=direct
const ExampleReplicaSetURI = "mongodb://goclub:goclub@localhost:27017/?authSource=goclub&readPreference=primary&directConnection=true&ssl=false"

type ExampleMovie struct {
	ID          primitive.ObjectID     `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Year        int                `bson:"year"`
	Genres      []string           `bson:"genres"`
	Rated       string             `bson:"rated"`
	Languages   []string           `bson:"languages"`
	Released    time.Time          `bson:"released"`
	Awards      ExampleMovieAwards `bson:"awards"`
	Cast        []string           `bson:"cast"`
	Directors   []string           `bson:"directors"`
	LastUpdated *time.Time         `bson:"lastupdated"`
}
type ExampleMovieAwards struct {
	Wins        int    `bson:"wins"`
	Nominations int    `bson:"nominations"`
	Text        string `bson:"text"`
}

func (v *ExampleMovie) AfterInsert(result ResultInsertOne) (err error) {
	if v.ID.IsZero() {
		v.ID, err = result.InsertedObjectID() ; if err != nil {
			return
		}
	}
	return
}

type ManyExampleMovie []ExampleMovie

func (many ManyExampleMovie) ManyD() (documents []interface{}, err error) {
	for _, v := range many {
		var b []byte
		b, err = bson.Marshal(v)
		if err != nil {
			return
		}
		documents = append(documents, bson.Raw(b))
	}
	return
}

func (many ManyExampleMovie) AfterInsertMany(result ResultInsertMany) (err error) {
	objectIDs, err := result.InsertedObjectIDs() ; if err != nil {
		return
	}
	for i, _ := range many {
		many[i].ID = objectIDs[i]
	}
	return
}

type ExampleComment struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     uint64             `bson:"userID"`
	NewsID     primitive.ObjectID `bson:"newsID"`
	Message    string             `bson:"message"`
	Like       uint64             `bson:"like"`
	DefaultLifeCycle
}

func (d ExampleComment) Field() (f struct {
	ID         string
	UserID     string
	NewsID     string
	Message    string
	Like       string
}) {
	f.ID = "_id"
	f.UserID = "userID"
	f.NewsID = "newsID"
	f.Message = "message"
	f.Like = "like"
	return
}

func (v *ExampleComment) AfterInsert(result ResultInsertOne) (err error) {
	if v.ID.IsZero() {
		v.ID, err = result.InsertedObjectID() ; if err != nil {
			return
		}
	}
	return
}

type ManyExampleComment []ExampleComment

func (many ManyExampleComment) ManyD() (documents []interface{}, err error) {
	for _, v := range many {
		var b []byte
		b, err = bson.Marshal(v)
		if err != nil {
			return
		}
		documents = append(documents, bson.Raw(b))
	}
	return
}

func (many ManyExampleComment) AfterInsertMany(result ResultInsertMany) (err error) {
	objectIDs, err := result.InsertedObjectIDs() ; if err != nil {
		return
	}
	for i, _ := range many {
		many[i].ID = objectIDs[i]
	}
	return
}

type ExampleNewsStatDaily struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Date       string             `bson:"date"`
	NewsID     primitive.ObjectID `bson:"newsID"`
	UV         uint64             `bson:"uv"`
	PV         uint64             `bson:"pv"`
	PlatformUV map[string]uint64  `bson:"platformUV"`
}

func (v *ExampleNewsStatDaily) AfterInsert(result ResultInsertOne) (err error) {
	if v.ID.IsZero() {
		v.ID, err = result.InsertedObjectID() ; if err != nil {
			return
		}
	}
	return
}

func (d ExampleNewsStatDaily) Field() (f struct {
	ID         string
	Date       string
	NewsID     string
	UV         string
	PV         string
	PlatformUV string
}) {
	f.ID = "_id"
	f.Date = "date"
	f.NewsID = "newsID"
	f.UV = "uv"
	f.PV = "pv"
	f.PlatformUV = "platformUV"
	return
}

type ManyExampleNewsStatDaily []ExampleNewsStatDaily

func (many ManyExampleNewsStatDaily) ManyD() (documents []interface{}, err error) {
	for _, v := range many {
		var b []byte
		b, err = bson.Marshal(v)
		if err != nil {
			return
		}
		documents = append(documents, bson.Raw(b))
	}
	return
}

func (many ManyExampleNewsStatDaily) AfterInsertMany(result ResultInsertMany) (err error) {
	objectIDs, err := result.InsertedObjectIDs() ; if err != nil {
		return
	}
	for i, _ := range many {
		many[i].ID = objectIDs[i]
	}
	return
}

type ExampleLocation struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Location Point              `bson:"location"`
}

func (v *ExampleLocation) AfterInsert(result ResultInsertOne) (err error) {
	if v.ID.IsZero() {
		v.ID, err = result.InsertedObjectID() ; if err != nil {
			return
		}
	}
	return
}

func (d ExampleLocation) Field() (f struct {
	ID       string
	Location string
}) {
	f.ID = "_id"
	f.Location = "location"
	return
}

type ManyExampleLocation []ExampleLocation

func (many ManyExampleLocation) ManyD() (documents []interface{}, err error) {
	for _, v := range many {
		var b []byte
		b, err = bson.Marshal(v)
		if err != nil {
			return
		}
		documents = append(documents, bson.Raw(b))
	}
	return
}

func (many ManyExampleLocation) AfterInsertMany(result ResultInsertMany) (err error) {
	objectIDs, err := result.InsertedObjectIDs() ; if err != nil {
		return
	}
	for i, _ := range many {
		many[i].ID = objectIDs[i]
	}
	return
}