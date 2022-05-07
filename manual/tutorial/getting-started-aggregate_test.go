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

func TestAggregate(t *testing.T) {
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
	var pipeline  []bson.M
	pipeline = append(pipeline, bson.M{
		"$unwind": "$genres",
	})
	pipeline = append(pipeline, bson.M{
		"$group": bson.M{
			"_id": "$genres",
			"genreCount": bson.M{
				"$sum": 1,
			},
		},
	})
	pipeline = append(pipeline, bson.M{
		"$sort": bson.M{
			"genreCount": -1,
		},
	})
	type Item struct {
		Name string `bson:"_id"`
		GenreCount int64 `bson:"genreCount"`
	}
	list := []Item{}
	err = moviesColl.Aggregate(ctx, pipeline, mo.AggregateCommand{}, &list) ; if err != nil {
		return
	}
	log.Print("len(list)", len(list))
	jsonb, err := json.MarshalIndent(list, "", "  ") ; if err != nil {
		return
	}
	log.Print("list:", string(jsonb))
}

