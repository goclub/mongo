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
	"time"
)

func TestFilterData(t *testing.T) {
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
	// $lt 查询
	{
		filter := bson.M{
			"released": bson.M{
				"$lt": time.Date(2000,1,1,0,0,0,0, time.UTC), // 中国时区用 time.FixedZone("CST", 8*3600) 代替 time.UTC
			},
		}
		list := mo.ManyExampleMovie{}
		err = moviesColl.Find(ctx, filter, &list, mo.FindCommand{}); if err != nil {
			return
		}
		log.Print("len($lt)", len(list))
		jsonb, err := json.MarshalIndent(list, "", "  ") ; if err != nil {
			return
		}
		log.Print("$lt:", string(jsonb))
	}
	// $gt 查询
	{
		filter := bson.M{
			"awards.wins": bson.M{
				"$gt": 100,
			},
		}
		list := mo.ManyExampleMovie{}
		err = moviesColl.Find(ctx, filter, &list, mo.FindCommand{}); if err != nil {
			return
		}
		log.Print("len($gt)", len(list))
		jsonb, err := json.MarshalIndent(list, "", "  ") ; if err != nil {
		return
	}
		log.Print("$gt:", string(jsonb))
	}
	// $in 查询
	{
		filter := bson.M{
			"languages": bson.M{
				"$in": []string{"Japanese", "Mandarin"},
			},
		}
		list := mo.ManyExampleMovie{}
		err = moviesColl.Find(ctx, filter, &list, mo.FindCommand{}); if err != nil {
		return
	}
		log.Print("len($in)", len(list))
		jsonb, err := json.MarshalIndent(list, "", "  ") ; if err != nil {
		return
	}
		log.Print("$in:", string(jsonb))
	}

}

