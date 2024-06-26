package storage

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"sytron-server/constants"
)

// DBinstance func
func DBinstance() *mongo.Client {
	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(constants.MONGODB_CONNECTION_STRING),
	)
	if err != nil {
		log.Printf("str %v", constants.MONGODB_CONNECTION_STRING)
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

// Client Database instance
var Client *mongo.Client = DBinstance()

// OpenCollection is a  function makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(constants.DATABASE_NAME).Collection(collectionName)

	return collection
}

func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database(constants.DATABASE_NAME).Collection(collectionName)
}
