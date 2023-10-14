package resolvers

import (
	"go.mongodb.org/mongo-driver/mongo"

	"sytron-server/database"
)

type CollectionResolver[T any] struct {
	collectionName string
	model          T
}

func (c *CollectionResolver[T]) TODO() T {
	return c.model
}

func (c *CollectionResolver[T]) GetCollection() *mongo.Collection {
	return database.GetCollection(c.GetCollectionName())
}

func (c *CollectionResolver[T]) GetCollectionName() string {
	return c.collectionName
}
