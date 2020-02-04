package models

import (
	"context"
	//"fmt"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_"gopkg.in/mgo.v2/bson"
	"log"
)

func MongoConnection(database string, coll string) *mongo.Collection {
	//// We need this because incoming requests to a server should create a context and outgoing calls to servers
	//// should accept a Context

	// set client options but why -
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil{
		log.Fatal(err)
	}

	collection := client.Database(database).Collection(coll)

	return collection
}