package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Defines the database operation each struct is required to expose (CRUD)
type CollectionDocument struct {
	ID          primitive.ObjectID  `bson:"_id" json:"_id"`
	CreatedTime *primitive.DateTime `bson:"created_time,omitempty" json:"created_time,omitempty"`
	UpdatedTime *primitive.DateTime `bson:"updated_time,omitempty" json:"updated_time,omitempty"`
}
