package resolvers

import (
	"sytron-server/database"

	"go.mongodb.org/mongo-driver/mongo"
)

type collectionResolver[T any] struct {
	collectionName string
	model          T
}

func (c *collectionResolver[T]) TODO() T {
	return c.model
}

func (c *collectionResolver[T]) GetCollection() *mongo.Collection {
	return database.GetCollection(c.GetCollectionName())
}

func (c *collectionResolver[T]) GetCollectionName() string {
	return c.collectionName
}
