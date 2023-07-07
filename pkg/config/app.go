package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var ctx = context.TODO()

func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database("new_db")
	// collection = client.Database("new_db").Collection("new_collection")
	//A collection is analogous to a table in a relational database

}

func GetDB() *mongo.Database {
	return db
}
