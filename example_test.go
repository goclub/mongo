package mo_test

import (
	"context"
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
var newsStatDaily *mo.Collection

func init () {
	ExampleNewDatabase()
	ExampleNewCollection()
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
func ExampleNewCollection() {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	commentColl = mo.NewCollection(db, "comment")
	newsStatDaily = mo.NewCollection(db, "newsStatDaily")
	// create indexes
	{
		f := mo.NewsStatDaily{}.Field()
		_, err = newsStatDaily.Core.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
			Keys: bson.D{{f.Date, 1}, {f.NewsID, 1}},
			Options: mongoOptions.Index().SetUnique(true),
		}) ; if err != nil {
			return
		}
	}
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
	_=ctx
	newsID := primitive.NewObjectID()
	stat := mo.NewsStatDaily{Date: "1949-10-01", NewsID: newsID}
	field := stat.Field()
	for i:=0;i<2;i++ {
		result, err := newsStatDaily.UpdateOne(ctx, bson.D{
			{field.Date, "2008-08-08"},
			{field.NewsID, newsID},
		}, bson.D{
			{
				"$set", bson.D{
				{field.UV, 0},
				{field.PV, 0},
			},
			},
		}, mo.UpdateCommand{
			// If true, a new document will be inserted if the filter does not match any documents in the collection. The
			// default value is false.
			Upsert: xtype.Bool(true),
		}) ; if err != nil {
			return
		}
		log.Printf("$set UpdateResult%+v", *result)
	}
	for i:=0;i<2;i++ {
		result, err := newsStatDaily.UpdateOne(ctx, bson.D{
			{field.Date, "2008-08-08"},
			{field.NewsID, newsID},
		}, bson.D{
			{
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
	}
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
			{UserID: 2, Message: "x"},
			{UserID: 2, Message: "y"},
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

