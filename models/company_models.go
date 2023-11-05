package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	CollectionDocument `bson:",inline"`

	Name    string `bson:"name,omitempty"     json:"name"`
	Email   string `bson:"email,omitempty"    json:"email"`
	Phone   string `bson:"phone,omitempty"    json:"phone"`
	LogoURL string `bson:"logo_url,omitempty" json:"logo_url"`

	// relations
	HQ primitive.ObjectID `bson:"hq,omitempty" json:"hq"`
}
