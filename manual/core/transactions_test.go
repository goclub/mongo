package core_test

import (
	"context"
	xerr "github.com/goclub/error"
	mo "github.com/goclub/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"testing"
)

func TestTransactions(t *testing.T) {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	ctx := context.Background()

	client, err := mongo.Connect(ctx, mongoOptions.Client().ApplyURI(mo.ExampleReplicaSetURI)) ; if err != nil {
		return
	}
	err = client.Ping(ctx, readpref.Primary()) ; if err != nil {
		return
	}
	db := mo.NewDatabase(client, "goclub")
	// 提前准备好测试用的集合 foo bar
	err = db.Core.CreateCollection(ctx, "foo") ; if err != nil {
		// log.Print(err)
	}
	err = db.Core.CreateCollection(ctx, "bar") ; if err != nil {
		// log.Print(err)
	}
	fooColl := mo.NewCollection(db, "foo")
	barColl := mo.NewCollection(db, "bar")
	session, err := client.StartSession() ; if err != nil {
	    return
	}
	defer session.EndSession(ctx)
	// cbResult 最终会作为 WithTransaction 的出参 result
	result, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (cbResult interface{}, err error) {
		ctx := 0;_=ctx // redefine ctx avert bug
		_, err = fooColl.Core.InsertOne(sessCtx, bson.D{{"abc", 1}}) ; if err != nil {
		    return
		}
		_, err = barColl.Core.InsertOne(sessCtx, bson.D{{"xyz", 999}}) ; if err != nil {
			return
		}
		return
	}) ; if err != nil {
	    return
	}
	log.Printf("result: %v", result)

}