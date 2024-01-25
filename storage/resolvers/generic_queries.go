package resolvers

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"sytron-server/storage"
)

// Reading

func (r *CollectionResolver[T]) FindOneByID(id string) (*T, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return storage.FindOneByID(r.GetCollectionName(), _id, &r.model)
}

func (r *CollectionResolver[T]) FindOne(filter interface{}) (*T, error) {
	return storage.FindOne(r.GetCollectionName(), filter, &r.model)
}

func (r *CollectionResolver[T]) FindMany(opts storage.PaginationOptions) ([]T, error) {
	return storage.FindMany(r.GetCollectionName(), r.model, opts)
}

func (r *CollectionResolver[T]) CountDocuments(filter interface{}) (count int64, err error) {
	count, err = r.GetCollection().CountDocuments(context.TODO(), filter)
	return
}

// Writing

func (r *CollectionResolver[T]) InsertOne(model T) (T, error) {
	return storage.InsertOne(r.GetCollectionName(), model)
}

func (r *CollectionResolver[T]) UpdateOne(id string, model T) (*T, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return storage.UpdateOne(r.GetCollectionName(), _id, &model)
}

// Destructive

func (r *CollectionResolver[T]) DeleteOne(id string) (err error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	return storage.DeleteOne[T](r.GetCollectionName(), _id)
}
