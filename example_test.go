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
)

func TestExample(t *testing.T) {
	suite.Run(t, new(TestExampleSuite))
}

type TestExampleSuite struct {
	suite.Suite
}

var db *mo.Database
var commentColl *mo.Collection
var newsStatDailyColl *mo.Collection
var locationCool *mo.Collection

func init() {
	ExampleNewDatabase()
	ExampleNewCollection()
	ExampleMigrate()
}
func ExampleNewDatabase() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, mongoOptions.Client().ApplyURI(mo.ExampleReplicaSetURI))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	db = mo.NewDatabase(client, "goclub")
}
func ExampleNewCollection() {
	/* In a formal environment ignore defer code */ var err error
	defer func() {
		if err != nil {
			xerr.PrintStack(err)
		}
	}()
	commentColl = mo.NewCollection(db, "exampleComment")
	newsStatDailyColl = mo.NewCollection(db, "exampleNewsStatDaily")
	locationCool = mo.NewCollection(db, "exampleLocation")
}

type MigrateActions struct {
}

func (MigrateActions) Migrate_2021_11_23__09_52_CreateExmapleCommentJSONSchema(db *mo.Database) (err error) {
	f := mo.ExampleComment{}.Field()
	var jsonSchema = bson.M{
		"bsonType":             "object",
		"additionalProperties": false,
		"properties": bson.M{
			f.ID: bson.M{
				"bsonType": "objectId",
			},
			f.UserID: bson.M{
				"bsonType": "number",
			},
			f.NewsID: bson.M{
				"bsonType": "objectId",
			},
			f.Message: bson.M{
				"bsonType": "string",
			},
			f.Like: bson.M{
				"bsonType": "number",
			},
		},
	}
	var validator = bson.M{
		"$jsonSchema": jsonSchema,
	}
	opts := mongoOptions.CreateCollection().SetValidator(validator)
	err = db.Core.CreateCollection(context.TODO(), "exampleComment", opts)
	if err != nil {
		return
	}
	return
}
func (MigrateActions) Migrate_2021_11_23__09_52_CreateExampleNewsStatDailyIndexs(db *mo.Database) (err error) {
	// create indexes
	f := mo.ExampleNewsStatDaily{}.Field()
	_, err = newsStatDailyColl.Core.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    bson.D{{f.Date, 1}, {f.NewsID, 1}},
		Options: mongoOptions.Index().SetUnique(true),
	})
	if err != nil {
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
func TestExampleCollection_InsertOne (t *testing.T) {
	ExampleCollection_InsertOne()
}
func ExampleCollection_InsertOne()  {
    ctx := context.Background()
	err := func() (err error){
		exampleComment := mo.ExampleComment{
			UserID:     1,
			NewsID:     primitive.NewObjectID(),
			Message:    "goclub/mongo",
		}
		_, err = commentColl.InsertOne(ctx, &exampleComment, mo.InsertOneCommand{})
		if err != nil {
			return
		}
		log.Printf("ExampleCollection_InsertOne: %+v", exampleComment)
		return
	}() ; if err != nil {
	    log.Printf("%+v",err)
	}
}


func (suite TestExampleSuite) TestCollection_InsertMany() {
	ExampleCollection_InsertMany()
}
func ExampleCollection_InsertMany() {
	/* In a formal environment ignore defer code */ var err error
	defer func() {
		if err != nil {
			xerr.PrintStack(err)
		}
	}()
	ctx := context.Background()
	exampleCommentList := mo.ManyExampleComment{
		{UserID: 1, Message: "a", NewsID: primitive.NewObjectID()},
		{UserID: 1, Message: "b", NewsID: primitive.NewObjectID()},
	}
	_, err = commentColl.InsertMany(ctx, &exampleCommentList, mo.InsertManyCommand{})
	if err != nil {
		return
	}
	log.Printf("ExampleCollection_InsertMany: %+v", exampleCommentList)
}

func (suite TestExampleSuite) TestCollection_InsertIgnore() {
	ExampleUpdateCommand_Options_Upsert_InsertIgnore()
}
func ExampleUpdateCommand_Options_Upsert_InsertIgnore() {
	/* In a formal environment ignore defer code */ var err error
	defer func() {
		if err != nil {
			xerr.PrintStack(err)
		}
	}()
	ctx := context.Background()
	newsID := primitive.NewObjectID()
	stat := mo.ExampleNewsStatDaily{Date: "1949-10-01", NewsID: newsID}
	field := stat.Field()
	wg := sync.WaitGroup{}
	// Simulation of concurrent
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			/* In a formal environment ignore defer code */ var err error
			defer func() {
				if err != nil {
					xerr.PrintStack(err)
				}
			}()
			result, err := newsStatDailyColl.UpdateOne(ctx,mo.FilterUpdate{
				Filter: bson.D{
					{field.Date, "2008-08-08"},
					{field.NewsID, newsID},
				},
				Update: bson.D{{
					// You can also change it to $set
					// Upsert is the key
					"$inc", bson.D{
						{field.UV, 1},
						{field.PV, 1},
					},
				}},
			}, mo.UpdateCommand{
				// If true, a new document will be inserted if the filter does not match any documents in the collection. The
				// default value is false.
				Upsert: xtype.Bool(true),
			})
			if err != nil {
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
	/* In a formal environment ignore defer code */ var err error
	defer func() {
		if err != nil {
			xerr.PrintStack(err)
		}
	}()
	ctx := context.Background()
	// FindOneByObjectID
	{
		ExampleComment := mo.ExampleComment{
			UserID:     1,
			NewsID:     primitive.NewObjectID(),
			Message:    "test find",
		}
		_, err = commentColl.InsertOne(ctx, &ExampleComment, mo.InsertOneCommand{})
		if err != nil {
			return
		}
		findExampleComment := mo.ExampleComment{}
		// ExampleComment.AliasObjectID = primitive.NewObjectID() // if unExampleComment this line of code, hasExampleComment will be false
		hasExampleComment, err := commentColl.FindOneByObjectID(ctx, ExampleComment.ID, &findExampleComment, mo.FindOneCommand{}) ; if err != nil {
			return
		}
		log.Print("ExampleCollection_Find FindOneByObjectID: ", findExampleComment, hasExampleComment)
	}
	// FindOne
	{
		exampleCommentList := mo.ManyExampleComment{
			{UserID: 2, Message: "x", NewsID: primitive.NewObjectID()},
			{UserID: 2, Message: "y", NewsID: primitive.NewObjectID()},
		}
		_, err = commentColl.InsertMany(ctx, &exampleCommentList, mo.InsertManyCommand{})
		if err != nil {
			return
		}
		exampleComment := mo.ExampleComment{}
		field := exampleComment.Field()
		has, err := commentColl.FindOne(ctx, bson.D{
			{field.UserID, 2},
		}, &exampleComment, mo.FindOneCommand{
			LookupQuery:  true,
			LookupResult: true,
		}) ; if err != nil {
			return
		}
		log.Print("ExampleCollection_Find FindOne: ", has, exampleComment)
	}
}

func (suite TestExampleSuite) TestAggregateMapStringUint64() {
	ExampleAggregateMapStringUint64()
}

// Aggregate map[string]uint64
func ExampleAggregateMapStringUint64() {
	/* In a formal environment ignore defer code */ var err error
	defer func() {
		if err != nil {
			xerr.PrintStack(err)
		}
	}()
	ctx := context.Background()
	newsID := primitive.NewObjectID()
	list := mo.ManyExampleNewsStatDaily{
		{
			Date:   "2011-01-01",
			NewsID: newsID,
			PlatformUV: map[string]uint64{
				"ios":     8,
				"android": 2,
				"web":     24,
			},
		},
		{
			Date:   "2011-01-02",
			NewsID: newsID,
			PlatformUV: map[string]uint64{
				"ios":     13,
				"android": 14,
				"web":     31,
			},
		},
	}
	_, err = newsStatDailyColl.InsertMany(ctx, &list, mo.InsertManyCommand{})
	if err != nil {
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
						{"id", "$" + field.Date},
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

	cursor, err := newsStatDailyColl.Core.Aggregate(ctx, pipeline)
	if err != nil {
		return
	}
	results := []bson.M{}
	err = cursor.All(ctx, &results)
	if err != nil {
		return
	}
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return
	}
	log.Printf("TestAggregateMapStringUint64: results: %s", data)
}

func (suite TestExampleSuite) TestPointGeoJSON() {
	ExamplePointGeoJSON()
}

// Aggregate map[string]uint64
func ExamplePointGeoJSON() {
	/* In a formal environment ignore defer code */ var err error
	defer func() {
		if err != nil {
			xerr.PrintStack(err)
		}
	}()
	ctx := context.Background()
	_, err = locationCool.InsertMany(ctx, &hanzhouWestLakeList, mo.InsertManyCommand{
		Ordered: xtype.Bool(true),
	})
	if err != nil {
		return
	}
	var targetList mo.ManyExampleLocation
	filter := bson.M{
		"location": bson.M{
			"$geoWithin": bson.M{
				"$center": []interface{}{
					// 杭州西湖中心
					[]float64{120.11947930184483, 30.235950232037645},
					10,
				},
			},
		},
	}
	err = locationCool.Find(ctx, filter, &targetList, mo.FindCommand{Limit: xtype.Uint64(100)}); if err != nil {
		return
	}
	log.Print("len(targetList)", len(targetList))
	var jsBD09Data [][2]float64
	for _, location := range targetList {
		bd09Data := location.Location.WGS84().BD09()
		jsBD09Data = append(jsBD09Data, [2]float64{bd09Data.Longitude, bd09Data.Latitude})
	}
	jsonb, err := json.Marshal(jsBD09Data)
	if err != nil {
		return
	}
	log.Print("jsBD09Data:\n", string(jsonb))
}

func TestPaging(t *testing.T) {
	ExamplePaging()
}
func ExamplePaging() {
	/* In a formal environment ignore defer code */ var err error
	defer func() {
		if err != nil {
			xerr.PrintStack(err)
		}
	}()
	ctx := context.Background()
	field := mo.ExampleComment{}.Field()
	delResult, err := commentColl.DeleteMany(ctx, bson.M{
		field.Message: "paging",
	}, mo.DeleteCommand{})
	if err != nil {
		return
	}
	log.Print("delResult.DeletedCount:", delResult.DeletedCount)
	var insertList mo.ManyExampleComment
	for i := uint64(0); i < 111; i++ {
		insertList = append(insertList, mo.ExampleComment{
			UserID:     i + 1,
			NewsID:     primitive.NewObjectID(),
			Message:    "paging",
			Like:       0,
		})
	}
	_, err = commentColl.InsertMany(ctx, &insertList, mo.InsertManyCommand{})
	if err != nil {
		return
	}
	var pagingList mo.ManyExampleComment
	total, err := commentColl.Paging(ctx, mo.Paging{
		Filter:    bson.M{
			field.Message: "paging",
		},
		FindCmd:  mo.FindCommand{},
		SlicePtr: &pagingList,
		CountCmd: mo.CountCommand{},
		Page:     1,
		PerPage:  10,
	}) ; if err != nil {
		return
	}
	log.Print("len(pagingList):", len(pagingList))
	log.Print("pagingList:", pagingList)
	log.Print("total:", total)
}