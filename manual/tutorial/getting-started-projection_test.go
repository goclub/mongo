package tutorial_test

import (
	"context"
	"encoding/json"
	xerr "github.com/goclub/error"
	mo "github.com/goclub/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"testing"
)

func TestProjection(t *testing.T) {
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
	{
		filter := bson.M{}
		type partMovies struct {
			ID primitive.ObjectID `bson:"_id"`
			Title string `bson:"title"`
			Directors []string `bson:"directors"`
			Year int `bson:"year"`
		}
		list := []partMovies{}
		err = moviesColl.Find(ctx, filter, &list, mo.FindCommand{
			Projection: bson.M{
				"title":     1,
				"directors": 1,
				"year":      1,
			},
		}); if err != nil {
			return
		}
		log.Print("len(title:1,directors:1,year:1)", len(list))
		jsonb, err := json.MarshalIndent(list, "", "  ") ; if err != nil {
		return
	}
		log.Print("title:1,directors:1,year:1:", string(jsonb))
	}
	{
		filter := bson.M{}
		type partMovies struct {
			Title string `bson:"title"`
			Genres []string `bson:"genres"`
		}
		list := []partMovies{}
		err = moviesColl.Find(ctx, filter, &list, mo.FindCommand{
			Projection: bson.M{
				"_id":    0,
				"title":  1,
				"genres": 1,
			},
		}); if err != nil {
		return
	}
		log.Print("len(_id:0,title:1,genres:1)", len(list))
		jsonb, err := json.MarshalIndent(list, "", "  ") ; if err != nil {
		return
	}
		log.Print("_id:0,title:1,genres:1:", string(jsonb))
	}

}

