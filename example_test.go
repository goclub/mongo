package mo_test

import (
	"context"
	xerr "github.com/goclub/error"
	mo "github.com/goclub/mongo"
	"github.com/stretchr/testify/suite"
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
	commentColl = mo.NewCollection(db, "comment")
}

func (suite TestExampleSuite) TestCollection_InsertOne() {
	ExampleCollection_InsertOne()
}
func ExampleCollection_InsertOne() {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	ctx := context.Background()
	comment := mo.EComment{
		UserID: 1,
		Message: "goclub/mongo",
	}
	_, err = commentColl.InsertOne(ctx, &comment) ; if err != nil {
		return
	}
	log.Printf("ExampleCollection_InsertOne: %+v", comment)
	// ExampleCollection_InsertOne: {ObjectID:ObjectID("613a2571f3526f555cdd39a0") UserID:1 Message:goclub/mongo}
}

func (suite TestExampleSuite) TestCollection_InsertMany() {
	ExampleCollection_InsertMany()
}
func ExampleCollection_InsertMany() {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	ctx := context.Background()
	commentList := mo.ManyEComment{
		{UserID: 1, Message: "a"},
		{UserID: 1, Message: "b"},
	}
	_, err = commentColl.InsertMany(ctx, &commentList) ; if err != nil {
		return
	}
	log.Printf("ExampleCollection_InsertMany: %+v", commentList)
	// ExampleCollection_InsertMany: [{ObjectID:ObjectID("613a2571f3526f555cdd399e") UserID:1 Message:a} {ObjectID:ObjectID("613a2571f3526f555cdd399f") UserID:1 Message:b}]
}
