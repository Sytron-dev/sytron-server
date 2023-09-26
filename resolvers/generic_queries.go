package resolvers

import (
	"sytron-server/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Reading

func (r *collectionResolver[T]) FindOneByID(id string) (*T, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return database.FindOneByID(r.GetCollectionName(), _id, &r.model)
}

func (r *collectionResolver[T]) FindMany(opts database.PaginationOptions) ([]T, error) {
	return database.FindMany(r.GetCollectionName(), r.model, opts)
}

// Writing

func (r *collectionResolver[T]) InsertOne(model T) (T, error) {
	return database.InsertOne[T](r.GetCollectionName(), model)
}

func (r *collectionResolver[T]) UpdateOne(id string, model T) (*T, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return database.UpdateOne(r.GetCollectionName(), _id, &model)
}

// Destructive

func (r *collectionResolver[T]) DeleteOne(id string) (err error) {

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	return database.DeleteOne[T](r.GetCollectionName(), _id)
}
