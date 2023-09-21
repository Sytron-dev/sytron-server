package models

import (
	"context"
	"sytron-server/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Defines the database operation each struct is required to expose (CRUD)
type CollectionDocument struct {
	ID          primitive.ObjectID  `bson:"_id" json:"_id"`
	CreatedTime *primitive.DateTime `bson:"created_time,omitempty" json:"created_time"`
	UpdatedTime *primitive.DateTime `bson:"updated_time,omitempty" json:"updated_time"`

	collectionName string
}

func NewCollectionDocument() (c *CollectionDocument) {
	return &CollectionDocument{
		collectionName: "test",
	}
}

func (c *CollectionDocument) GetCollection() *mongo.Collection {
	return database.GetCollection(c.collectionName)
}

func (c *CollectionDocument) SetID(id string) (_id primitive.ObjectID) {
	_id, _ = primitive.ObjectIDFromHex(id)
	c.ID = _id
	return
}

func (c *CollectionDocument) FindOneByID() (doc *CollectionDocument, err error) {
	return database.FindOneByID(c.collectionName, c.ID, c)
}

func (c *CollectionDocument) UpdateOne(model CollectionDocument) (doc *CollectionDocument, err error) {

	// Time
	now := time.Now()
	*model.UpdatedTime = primitive.NewDateTimeFromTime(now)
	if model.CreatedTime == nil {
		*model.CreatedTime = primitive.NewDateTimeFromTime(now)
	}

	// Filters
	filter := bson.D{{Key: "_id", Value: model.ID}}
	update := bson.M{"$set": model}

	if _, err := model.GetCollection().UpdateOne(context.TODO(), filter, update); err != nil {
		return nil, err
	}

	// Return value
	return model.FindOneByID()
}
