package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientMongo *mongo.Client
var dbName *string

type IDBinstance interface {
	InitDB(host string, port string, dbname string) *mongo.Client
	OpenCollection(name string) *mongo.Collection
}

type DBinstance struct{}

func (db *DBinstance) InitDB(host string, port string, dbname string) *mongo.Client {

	MongoDb := "mongodb://" + host + ":" + port

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if clientMongo == nil {
		clientMongo = client
	}

	if dbName == nil {
		dbName = &dbname
	}

	return client
}

//OpenCollection is a  function makes a connection with a collection in the database
func (db *DBinstance) OpenCollection(collectionName string) *mongo.Collection {
	var collection *mongo.Collection = clientMongo.Database(*dbName).Collection(collectionName)
	return collection
}
