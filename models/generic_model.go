package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Defines the database operation each struct is required to expose (CRUD)
type CollectionDocument struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	CreatedTime primitive.DateTime `bson:"created_time"  json:"created_time,omitempty"`
	UpdatedTime primitive.DateTime `bson:"updated_time"  json:"updated_time,omitempty"`
}

func (doc *CollectionDocument) InsertTime() {
	doc.CreatedTime = primitive.NewDateTimeFromTime(time.Now())
	doc.UpdatedTime = primitive.NewDateTimeFromTime(time.Now())
}

func (doc *CollectionDocument) UpdateTime() {
	doc.UpdatedTime = primitive.NewDateTimeFromTime(time.Now())
}

func (doc *CollectionDocument) SetID() {
	doc.ID = primitive.NewObjectID()
}
