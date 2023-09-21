package resolvers

import (
	"sytron-server/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *collectionResolver[T]) FindOneByID(id string) (*T, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return database.FindOneByID(r.GetCollectionName(), _id, &r.model)
}
