package resolvers

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"sytron-server/database"
)

// Reading

func (r *CollectionResolver[T]) FindOneByID(id string) (*T, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return database.FindOneByID(r.GetCollectionName(), _id, &r.model)
}

func (r *CollectionResolver[T]) FindOne(filter interface{}) (*T, error) {
	return database.FindOne(r.GetCollectionName(), filter, &r.model)
}

func (r *CollectionResolver[T]) FindMany(opts database.PaginationOptions) ([]T, error) {
	return database.FindMany(r.GetCollectionName(), r.model, opts)
}

func (r *CollectionResolver[T]) CountDocuments(filter interface{}) (count int64, err error) {
	count, err = r.GetCollection().CountDocuments(context.TODO(), filter)
	return
}

// Writing

func (r *CollectionResolver[T]) InsertOne(model T) (T, error) {
	return database.InsertOne(r.GetCollectionName(), model)
}

func (r *CollectionResolver[T]) UpdateOne(id string, model T) (*T, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return database.UpdateOne(r.GetCollectionName(), _id, &model)
}

// Destructive

func (r *CollectionResolver[T]) DeleteOne(id string) (err error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	return database.DeleteOne[T](r.GetCollectionName(), _id)
}
