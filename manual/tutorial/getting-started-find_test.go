package tutorial_test

import (
	"context"
	"encoding/json"
	xerr "github.com/goclub/error"
	mo "github.com/goclub/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"testing"
)

func TestFind(t *testing.T) {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	ctx := context.Background()
	client, err := mongo.Connect(ctx, mongoOptions.Client().ApplyURI(mo.ExampleReplicaSetURI)) ; if err != nil {
		return
	}
	err = client.Ping(ctx, readpref.Primary()) ; if err != nil {
		return
	}
	db := mo.NewDatabase(client, "goclub")
	moviesColl := mo.NewCollection(db, "movies")
	// 正文开始
	filter := bson.M{}
	list := mo.ManyExampleMovie{}
	err = moviesColl.Find(ctx, filter, &list, mo.FindCommand{}); if err != nil {
	    return
	}
	log.Print("len(list)", len(list))
	jsonb, err := json.MarshalIndent(list, "", "  ") ; if err != nil {
	    return
	}
	log.Print("list:", string(jsonb))

}

