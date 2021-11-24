package mo_test

import (
	"context"
	"encoding/json"
	xerr "github.com/goclub/error"
	mo "github.com/goclub/mongo"
	xtype "github.com/goclub/type"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"sync"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	suite.Run(t, new(TestExampleSuite))
}

type TestExampleSuite struct {
	suite.Suite
}

var db *mo.Database
var exampleCommentColl *mo.Collection
var exampleNewsStatDailyColl *mo.Collection
var exampleLocationCool *mo.Collection

func init () {
	ExampleNewDatabase()
	ExampleNewCollection()
	ExampleMigrate()
}
func ExampleNewDatabase() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, mongoOptions.Client().ApplyURI("mongodb://localhost:27017")) ; if err != nil {
		return
	}
	err = client.Ping(ctx, readpref.Primary()) ; if err != nil {
		return
	}
	db = mo.NewDatabase(client, "goclubMongo")
}
func ExampleNewCollection() {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	exampleCommentColl = mo.NewCollection(db, "exampleComment")
	exampleNewsStatDailyColl = mo.NewCollection(db, "exampleNewsStatDaily")
	exampleLocationCool = mo.NewCollection(db, "exampleLocation")
}
type MigrateActions struct {

}
func (MigrateActions) Migrate_2021_11_23__09_52_CreateExmapleCommentJSONSchema(db *mo.Database) (err error) {
	f := mo.ExampleComment{}.Field()
	var jsonSchema = bson.M{
		"bsonType":             "object",
		"required":             []string{f.UserID, f.NewsID, f.CreateTime, f.Message},
		"additionalProperties": false,
		"properties": bson.M{
			"_id": bson.M{
				"bsonType": "objectId",
			},
			"userID": bson.M{
				"bsonType":    "number",
			},
			"newsID": bson.M{
				"bsonType":    "objectId",
			},
			"message": bson.M{
				"bsonType":    "string",
			},
			"like": bson.M{
				"bsonType":    "number",
			},
			"createTime": bson.M{
				"bsonType":    "date",
			},
		},
	}
	var validator = bson.M{
		"$jsonSchema": jsonSchema,
	}
	opts := mongoOptions.CreateCollection().SetValidator(validator)
	err = db.Core.CreateCollection(context.TODO(), "exampleComment", opts) ; if err != nil {
	    return
	}
	return
}
func (MigrateActions) Migrate_2021_11_23__09_52_CreateExampleNewsStatDailyIndexs(db *mo.Database) (err error) {
	// create indexes
	f := mo.ExampleNewsStatDaily{}.Field()
	_, err = exampleNewsStatDailyColl.Core.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.D{{f.Date, 1}, {f.NewsID, 1}},
		Options: mongoOptions.Index().SetUnique(true),
	}) ; if err != nil {
		return
	}
	return
}
// In a formal project you should use `go run cmd/migrate/main.go`, not running in init function
func ExampleMigrate() {
	mo.Migrate(db, &MigrateActions{})
}

func (suite TestExampleSuite) TestCollection_InsertOne() {
	ExampleCollection_InsertOne()
}
func ExampleCollection_InsertOne() {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	ctx := context.Background()
	exampleComment := mo.ExampleComment{
		UserID: 1,
		NewsID: primitive.NewObjectID(),
		Message: "goclub/mongo",
		CreateTime: time.Now(),
	}
	_, err = exampleCommentColl.InsertOne(ctx, &exampleComment, mo.InsertOneCommand{
		ByPassDocumentValidation: xtype.Bool(true),
	}) ; if err != nil {
		return
	}
	log.Printf("ExampleCollection_InsertOne: %+v", exampleComment)
}

func (suite TestExampleSuite) TestCollection_InsertMany() {
	ExampleCollection_InsertMany()
}
func ExampleCollection_InsertMany() {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	ctx := context.Background()
	exampleCommentList := mo.ManyExampleComment{
		{UserID: 1, Message: "a", NewsID: primitive.NewObjectID(), CreateTime: time.Now()},
		{UserID: 1, Message: "b", NewsID: primitive.NewObjectID(), CreateTime: time.Now()},
	}
	_, err = exampleCommentColl.InsertMany(ctx, &exampleCommentList, mo.InsertManyCommand{}) ; if err != nil {
		return
	}
	log.Printf("ExampleCollection_InsertMany: %+v", exampleCommentList)
}

func (suite TestExampleSuite) TestCollection_InsertIgnore() {
	ExampleCollection_InsertIgnore()
}
func ExampleCollection_InsertIgnore() {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	ctx := context.Background()
	newsID := primitive.NewObjectID()
	stat := mo.ExampleNewsStatDaily{Date: "1949-10-01", NewsID: newsID}
	field := stat.Field()
	wg := sync.WaitGroup{}
	// Simulation of concurrent
	for i:=0;i<2;i++ {
		wg.Add(1)
		go func() {
			/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
			result, err := exampleNewsStatDailyColl.UpdateOne(ctx, bson.D{
				{field.Date, "2008-08-08"},
				{field.NewsID, newsID},
			}, bson.D{
				{
					// You can also change it to $set
					// Upsert is the key
					"$inc", bson.D{
					{field.UV, 1},
					{field.PV, 1},
				},
				},
			}, mo.UpdateCommand{
				// If true, a new document will be inserted if the filter does not match any documents in the collection. The
				// default value is false.
				Upsert: xtype.Bool(true),
			}) ; if err != nil {
				return
			}
			log.Printf("$inc UpdateResult%+v", *result)
			wg.Done()
		}()
	}
	wg.Wait()
}

func (suite TestExampleSuite) TestCollection_Find() {
	ExampleCollection_Find()
}
func ExampleCollection_Find() {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	ctx := context.Background()
	// FindByObjectID
	{
		ExampleComment := mo.ExampleComment{
			UserID:   1,
			NewsID: primitive.NewObjectID(),
			Message:  "test find",
			CreateTime: time.Now(),
		}
		_, err = exampleCommentColl.InsertOne(ctx, &ExampleComment, mo.InsertOneCommand{}) ; if err != nil {
		return
	}
		findExampleComment := mo.ExampleComment{}
		// ExampleComment.ObjectID = primitive.NewObjectID() // if unExampleComment this line of code, hasExampleComment will be false
		hasExampleComment, err := exampleCommentColl.FindByObjectID(ctx, ExampleComment.ID, &findExampleComment, mo.FindOneCommand{}) ; if err != nil {
			return
		}
		log.Print("ExampleCollection_Find FindByObjectID: ", findExampleComment, hasExampleComment)
	}
	// FindOne
	{
		exampleCommentList := mo.ManyExampleComment{
			{UserID: 2, Message: "x",NewsID: primitive.NewObjectID(),CreateTime: time.Now(),},
			{UserID: 2, Message: "y",NewsID: primitive.NewObjectID(),CreateTime: time.Now(),},
		}
		_, err = exampleCommentColl.InsertMany(ctx, &exampleCommentList, mo.InsertManyCommand{}) ; if err != nil {
			return
		}
		ExampleComment := mo.ExampleComment{}
		field := ExampleComment.Field()
		has, err := exampleCommentColl.FindOne(ctx, bson.D{
			{field.UserID, 2},
		}, &ExampleComment, mo.FindOneCommand{}) ; if err != nil {
		    return
		}
		log.Print("ExampleCollection_Find FindOne: ", has, ExampleComment)
	}
}

func (suite TestExampleSuite) TestAggregateMapStringUint64() {
	ExampleAggregateMapStringUint64()
}
// Aggregate map[string]uint64
func ExampleAggregateMapStringUint64() {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	ctx := context.Background()
	newsID := primitive.NewObjectID()
	list := mo.ManyExampleNewsStatDaily{
		{
			Date: "2011-01-01",
			NewsID: newsID,
			PlatformUV: map[string]uint64{
				"ios": 8,
				"android": 2,
				"web": 24,
			},
		},
		{
			Date: "2011-01-02",
			NewsID: newsID,
			PlatformUV: map[string]uint64{
				"ios": 13,
				"android": 14,
				"web": 31,
			},
		},
	}
	_, err = exampleNewsStatDailyColl.InsertMany(ctx, &list, mo.InsertManyCommand{}) ; if err != nil {
	    return
	}
	field := mo.ExampleNewsStatDaily{}.Field()
	var pipeline []bson.D
	pipeline = append(pipeline, bson.D{
		{"$match", bson.D{
			{field.NewsID, newsID},
			{field.Date, bson.D{
				{"$gte", "2011-01-01"},
				{"$lte", "2011-01-02"},
			}},
		}},
	})

	pipeline = append(pipeline, bson.D{
		{"$addFields", bson.D{
			{"keys", bson.D{{"$objectToArray", "$" + field.PlatformUV}}},
		},
		},
	})
	pipeline = append(pipeline, bson.D{
		{
			"$unwind", "$keys",
		},
	})
	pipeline = append(pipeline, bson.D{
		{
			"$group", bson.D{
				{
					"_id", bson.D{
					{"id", "$" + field.Date,},
					{"key", "$keys.k"},
				},
				},
				{
					"sumUV", bson.D{
					{"$sum", "$keys.v"},
				},
				},
			},
		},
	})
	pipeline = append(pipeline, bson.D{
		{
			"$group", bson.D{
				{"_id", "$_id.key"},
				{"total", bson.D{{"$sum", "$sumUV"}}},
			},
		},
	})

	cursor, err := exampleNewsStatDailyColl.Core.Aggregate(ctx, pipeline) ; if err != nil {
	    return
	}
	results := []bson.M{}
	err = cursor.All(ctx, &results) ; if err != nil {
	    return
	}
	data, err := json.MarshalIndent(results, "", "  ") ; if err != nil {
	    return
	}
	log.Printf("TestAggregateMapStringUint64: results: %s", data)
}

func (suite TestExampleSuite) TestGeoJSONPoint() {
	ExampleGeoJSONPoint()
}
// Aggregate map[string]uint64
func ExampleGeoJSONPoint() {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	ctx := context.Background()
	_, err = exampleLocationCool.InsertMany(ctx, &hanzhouWestLakeList, mo.InsertManyCommand{
		Ordered: xtype.Bool(true),
	}) ; if err != nil {
	    return
	}
	var targetList mo.ManyExampleLocation
	filter := bson.M{
		"location": bson.M{
			"$geoWithin": bson.M{
				"$center": []interface{}{
					[]float64{120.11947930184483, 30.235950232037645},
					1,
				},
			},
		},
	}
	cursor, err := exampleLocationCool.Core.Find(ctx, filter, mongoOptions.Find().SetLimit(500)) ; if err != nil {
	    return
	}
	err = cursor.All(ctx, &targetList) ; if err != nil {
	    return
	}
	log.Print("len(targetList)", len(targetList))
	var jsBD09Data [][2]float64
	for _, location := range targetList {
		bd09Data := location.Location.BD09()
		jsBD09Data = append(jsBD09Data, [2]float64{bd09Data.Longitude, bd09Data.Latitude})
	}
	jsonb , err := json.Marshal(jsBD09Data) ; if err != nil {
	    return
	}
	log.Print("jsBD09Data:\n", string(jsonb))
}