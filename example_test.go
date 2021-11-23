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
var commentColl *mo.Collection
var newsStatDailyCool *mo.Collection

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
	db = mo.NewDatabase(client, "goclub_mongo")
}
type MigrateActions struct {

}
func (MigrateActions) Migrate_2021_11_23__09_52_CreateNewsStatDailyIndexs(db *mo.Database) (err error) {
	// create indexes
	f := mo.NewsStatDaily{}.Field()
	_, err = newsStatDailyCool.Core.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
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
func ExampleNewCollection() {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	commentColl = mo.NewCollection(db, "comment")
	newsStatDailyCool = mo.NewCollection(db, "newsStatDaily")
}

func (suite TestExampleSuite) TestCollection_InsertOne() {
	ExampleCollection_InsertOne()
}
func ExampleCollection_InsertOne() {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	ctx := context.Background()
	comment := mo.Comment{
		UserID: 1,
		Message: "goclub/mongo",
	}
	_, err = commentColl.InsertOne(ctx, &comment, mo.InsertOneCommand{}) ; if err != nil {
		return
	}
	log.Printf("ExampleCollection_InsertOne: %+v", comment)
}

func (suite TestExampleSuite) TestCollection_InsertMany() {
	ExampleCollection_InsertMany()
}
func ExampleCollection_InsertMany() {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	ctx := context.Background()
	commentList := mo.ManyComment{
		{UserID: 1, Message: "a"},
		{UserID: 1, Message: "b"},
	}
	_, err = commentColl.InsertMany(ctx, &commentList, mo.InsertManyCommand{}) ; if err != nil {
		return
	}
	log.Printf("ExampleCollection_InsertMany: %+v", commentList)
}

func (suite TestExampleSuite) TestCollection_InsertIgnore() {
	ExampleCollection_InsertIgnore()
}
func ExampleCollection_InsertIgnore() {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	ctx := context.Background()
	newsID := primitive.NewObjectID()
	stat := mo.NewsStatDaily{Date: "1949-10-01", NewsID: newsID}
	field := stat.Field()
	wg := sync.WaitGroup{}
	// Simulation of concurrent
	for i:=0;i<2;i++ {
		wg.Add(1)
		go func() {
			/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
			result, err := newsStatDailyCool.UpdateOne(ctx, bson.D{
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
		comment := mo.Comment{
			UserID:   1,
			Message:  "test find",
			DateTime: time.Now(),
		}
		_, err = commentColl.InsertOne(ctx, &comment, mo.InsertOneCommand{}) ; if err != nil {
		return
	}
		findComment := mo.Comment{}
		// comment.ObjectID = primitive.NewObjectID() // if uncomment this line of code, hasComment will be false
		hasComment, err := commentColl.FindByObjectID(ctx, comment.ID, &findComment, mo.FindOneCommand{}) ; if err != nil {
			return
		}
		log.Print("ExampleCollection_Find FindByObjectID: ", findComment, hasComment)
	}
	// FindOne
	{
		commentList := mo.ManyComment{
			{UserID: 2, Message: "x",DateTime: time.Now(),},
			{UserID: 2, Message: "y",DateTime: time.Now(),},
		}
		_, err = commentColl.InsertMany(ctx, &commentList, mo.InsertManyCommand{}) ; if err != nil {
			return
		}
		comment := mo.Comment{}
		field := comment.Field()
		has, err := commentColl.FindOne(ctx, bson.D{
			{field.UserID, 2},
		}, &comment, mo.FindOneCommand{}) ; if err != nil {
		    return
		}
		log.Print("ExampleCollection_Find FindOne: ", has, comment)
	}
}

// Aggregate map[string]uint64
func TestAggregateMapStringUint64(t *testing.T) {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	ctx := context.Background()
	newsID := primitive.NewObjectID()
	list := mo.ManyNewsStatDaily{
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
	_, err = newsStatDailyCool.InsertMany(ctx, &list, mo.InsertManyCommand{}) ; if err != nil {
	    return
	}
	field := mo.NewsStatDaily{}.Field()
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

	cursor, err := newsStatDailyCool.Core.Aggregate(ctx, pipeline) ; if err != nil {
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
