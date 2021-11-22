package mo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client *mongo.Client
	Core *mongo.Database
	name string
}
func NewDatabase(client *mongo.Client, dbName string, opts ...*options.DatabaseOptions) *Database {
	return &Database{
		client: client,
		Core: client.Database(dbName, opts...),
		name: dbName,
	}
}