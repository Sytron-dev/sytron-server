package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	CollectionDocument `bson:",inline" json:"collection_document"`

	Name    string `bson:"name,omitempty"     json:"name"`
	Email   string `bson:"email,omitempty"    json:"email"`
	LogoURL string `bson:"logo_url,omitempty" json:"logo_url"`

	// relations
	HQ primitive.ObjectID `bson:"hq,omitempty" json:"hq"`
}
