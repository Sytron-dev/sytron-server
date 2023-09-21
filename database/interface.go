package database

import (
	"context"
	"log"
	"sytron-server/helpers/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ---- Create operations ----------------------------------------------------------------
func InsertOne() {}

// ---- Read Operations ------------------------------------------------------------------

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

type PaginationOptions struct {
	pageSize   int64
	pageNumber int64
	filter     interface{}
}

// TODO use valid generics for return type
func FindMany[ModelType any](collectionName string, model ModelType, paginationOptions PaginationOptions) (docs []map[string]interface{}, err error) {
	log.SetFlags(log.Ldate | log.Lshortfile)

	if paginationOptions.pageSize > 20 || paginationOptions.pageSize < 1 {
		paginationOptions.pageSize = 20
	}
	if paginationOptions.pageNumber < 1 {
		paginationOptions.pageNumber = 1
	}
	if paginationOptions.filter == nil {
		paginationOptions.filter = bson.D{}
	}

	log.Printf("Finding many from %s of type %T\n", collectionName, docs)
	log.Printf("					: %v", paginationOptions)

	// get  database collection from collectionName
	collection := GetCollection(collectionName)

	// Get many values with pagination
	offset := (paginationOptions.pageNumber - 1) * paginationOptions.pageSize
	opts := options.FindOptions{
		Limit: &paginationOptions.pageSize,
		Skip:  &offset,
	}

	cursor, err := collection.Find(context.TODO(), paginationOptions.filter, &opts)
	if err != nil {
		return
	}

	for cursor.Next(context.TODO()) {
		var doc map[string]interface{}
		err := cursor.Decode(&doc)
		if err != nil {
			break
		}
		docs = append(docs, doc)
	}

	return docs, err
}

// ---- Update operations ----------------------------------------------------------------

// ---- Delete operations ----------------------------------------------------------------

func DeleteOne() {}
