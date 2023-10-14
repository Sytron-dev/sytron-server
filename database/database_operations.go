package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"sytron-server/helpers/logger"
)

// ---- Create operations ----------------------------------------------------------------
func InsertOne[ModelType any](collectionName string, model ModelType) (doc ModelType, err error) {
	collection := GetCollection(collectionName)

	_, err = collection.InsertOne(context.TODO(), model)

	if err != nil {
		return
	}

	doc = model
	return
}

// ---- Read Operations ------------------------------------------------------------------

func FindOne[ModelType any](
	collectionName string,
	filter interface{},
	model ModelType,
) (doc ModelType, err error) {
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Printf("Finding %v from %s", filter, collectionName)

	// Get database collection from collectionName
	collection := GetCollection(collectionName)

	// Write the result to an arbitrary interface
	if err = collection.FindOne(context.TODO(), filter).Decode(model); err != nil {
		logger.Error(err, "")
	} else {
		doc = model
		logger.Log(doc)
	}

	return
}

func FindOneByID[ModelType any](
	collectionName string,
	_id primitive.ObjectID,
	model ModelType,
) (doc ModelType, err error) {
	filter := bson.D{{Key: "_id", Value: _id}}
	return FindOne[ModelType](collectionName, filter, model)
}

type PaginationOptions struct {
	pageSize   int64
	pageNumber int64
	filter     interface{}
}

// TODO use valid generics for return type
func FindMany[ModelType any](
	collectionName string,
	model ModelType,
	paginationOptions PaginationOptions,
) (docs []ModelType, err error) {
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
	log.Printf("                : %v", paginationOptions)

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
		err := cursor.Decode(&model)
		if err != nil {
			break
		}
		docs = append(docs, model)
	}

	return docs, err
}

// ---- Update operations ----------------------------------------------------------------

func UpdateOne[ModelType any](
	collectionName string,
	_id primitive.ObjectID,
	model ModelType,
) (doc ModelType, err error) {
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Printf("Updating %s from %s", _id, collectionName)

	// Get database collection from collectionName
	collection := GetCollection(collectionName)
	filter := bson.D{{Key: "_id", Value: _id}}
	update := bson.D{{Key: "$set", Value: model}}

	// Write the result to an arbitrary interface
	if _, err = collection.UpdateOne(context.TODO(), filter, update); err != nil {
		logger.Error(err, "write error")
	} else {
		doc = model
	}

	return
}

// ---- Delete operations ----------------------------------------------------------------

func DeleteOne[ModelType any](collectionName string, _id primitive.ObjectID) (err error) {
	collection := GetCollection(collectionName)

	filter := bson.D{{Key: "_id", Value: _id}}
	opts := options.DeleteOptions{}

	_, err = collection.DeleteOne(context.TODO(), filter, &opts)

	return
}
