package tutorial_test

import (
	"context"
	xerr "github.com/goclub/error"
	mo "github.com/goclub/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
	"time"
)

var db *mo.Database
var moviesColl  *mo.Collection

func init () {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { panic(err) } }()
	ctx := context.Background()
	client, err := mongo.Connect(ctx, mongoOptions.Client().ApplyURI("mongodb://goclub:goclub@localhost:27017/goclub?authSource=goclub")) ; if err != nil {
		return
	}
	err = client.Ping(ctx, readpref.Primary()) ; if err != nil {
		return
	}
	db = mo.NewDatabase(client, "goclub")
	moviesColl = mo.NewCollection(db, "movies")
}

func TestInsert(t *testing.T) {
	/* In a formal environment ignore defer code */var err error;defer func() { if err != nil { xerr.PrintStack(err) } }()
	ctx := context.Background()
	documents := mo.ManyExampleMovie{
		{
			Title: "Titanic",
			Year: 1997,
			Genres: []string{"Drama", "Romance"},
			Rated: "PG-13",
			Languages: []string{"English", "French", "German", "Swedish", "Italian", "Russian" },
			Released: time.Date(1997, 12, 19, 0,0,0,0,time.UTC),
			Awards: mo.ExampleMovieAwards{
				Wins:        127,
				Nominations: 63,
				Text:        "Won 11 Oscars. Another 116 wins & 63 nominations.",
			},
			Cast: []string{"Leonardo DiCaprio", "Kate Winslet", "Billy Zane", "Kathy Bates"},
			Directors: []string{"James Cameron"},
		},
		{
			Title:     "The Dark Knight",
			Year:      2008,
			Genres:    []string{"Action", "Crime", "Drama" },
			Rated:     "PG-13",
			Languages: []string{"English", "Mandarin" },
			Released: time.Date(2008, 07, 18, 0,0,0,0,time.UTC),
			Awards:    mo.ExampleMovieAwards{
				Wins:        144,
				Nominations: 106,
				Text:        "Won 2 Oscars. Another 142 wins & 106 nominations.",
			},
			Cast:      []string{"Christian Bale", "Heath Ledger", "Aaron Eckhart", "Michael Caine"},
			Directors: []string{"Christopher Nolan"},
		},
		{
			Title:     "Spirited Away",
			Year:      2001,
			Genres:    []string{ "Animation", "Adventure", "Family" },
			Rated:     "PG",
			Languages: []string{"Japanese"},
			Released:  time.Date(2003, 3, 28, 0,0,0,0,time.UTC),
			Awards:    mo.ExampleMovieAwards{
				Wins:        52,
				Nominations: 22,
				Text:        "Won 1 Oscar. Another 51 wins & 22 nominations.",
			},
			Cast:      []string{"Rumi Hiiragi", "Miyu Irino", "Mari Natsuki", "Takashi Naitè"},
			Directors: []string{"Hayao Miyazaki"},
		},
		{
			Title:     "Casablanca",
			Year:      1942,
			Genres:    []string{"Drama", "Romance", "War" },
			Rated:     "PG",
			Languages: []string{"Humphrey Bogart", "Ingrid Bergman", "Paul Henreid", "Claude Rains" },
			Released:  time.Date(1943, 1, 23, 0,0,0, 0,time.UTC),
			Awards:    mo.ExampleMovieAwards{
				Wins:        9,
				Nominations: 6,
				Text:        "Won 3 Oscars. Another 6 wins & 6 nominations.",
			},
			LastUpdated: time.Date(2015, 9, 4, 0,22,54, 0,time.UTC),
			Directors: []string{"Michael Curtiz"},
		},
	}

	_, err = moviesColl.InsertMany(ctx, documents, mo.InsertManyCommand{}) ; if err != nil {
	    return
	}
}

