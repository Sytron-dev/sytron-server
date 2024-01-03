package models

import (
	"time"

	"github.com/google/uuid"
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

type SqlDocument struct {
	ID        uuid.UUID `json:"_id,omitempty"        db:"_id"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (doc *SqlDocument) InsertTime() {
	doc.CreatedAt = time.Now().UTC()
	doc.UpdatedAt = time.Now().UTC()
}

func (doc *SqlDocument) UpdateTime() {
	doc.UpdatedAt = time.Now().UTC()
}
