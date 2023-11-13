package resolvers

import (
	"sytron-server/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

type CollectionResolver[T any] struct {
	collectionName string
	model          T
}

func (c *CollectionResolver[T]) TODO() T {
	return c.model
}

func (c *CollectionResolver[T]) GetCollection() *mongo.Collection {
	return storage.GetCollection(c.GetCollectionName())
}

func (c *CollectionResolver[T]) GetCollectionName() string {
	return c.collectionName
}

type PqlResolver[T any] struct {
	tableName      string
	model          T
	createQuery    string
	readOneQuery   string
	readManyQuery  string
	updateOneQuery string
	deleteOneQuery string
}
