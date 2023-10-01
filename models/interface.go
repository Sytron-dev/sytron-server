package models

import (
	//"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Defines the database operation each struct is required to expose (CRUD)
type CollectionDocument struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	CreatedTime primitive.DateTime `bson:"created_time,omitempty"`
	UpdatedTime primitive.DateTime `bson:"updated_time,omitempty"`
}

type CollectionType interface {
	CollectionDocument | Location | Destination
}

func SetDocumentTime[T CollectionType](model T) {

	println(model)
	/*
		model.CreatedTime = primitive.NewDateTimeFromTime(time.Now())
		model.UpdatedTime = primitive.NewDateTimeFromTime(time.Now())
	*/
}
