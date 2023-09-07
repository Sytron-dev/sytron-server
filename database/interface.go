package database

import (
	"context"
	"log"
	"sytron-server/helpers/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertOne() {}

func FindOneByValue(key string, value interface{}, collectionName string) (document *interface{}, err error) {

	// Get database collection from collectionName
	collection := GetCollection(collectionName)
	filter := bson.D{{Key: key, Value: value}}

	// Write the result to an arbitrary interface
	err = collection.FindOne(context.TODO(), filter).Decode(document)

	return
}

func FindOneByID[ModelType any](collectionName string, _id primitive.ObjectID, model ModelType) (doc ModelType, err error) {

	log.SetFlags(log.Ldate | log.Lshortfile)

	log.Printf("Finding %s from %s", _id, collectionName)

	// Get database collection from collectionName
	collection := GetCollection(collectionName)
	filter := bson.D{{Key: "_id", Value: _id}}

	// Write the result to an arbitrary interface
	if err = collection.FindOne(context.TODO(), filter).Decode(model); err != nil {
		logger.Error(err, "")
	} else {
		doc = model
		logger.Log(doc)
	}

	return
}

func FindMany() {}

func FindManyPaginated() {}

func FindUpdateOne() {}

func FindDeleteOne() {}
