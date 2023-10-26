package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"sytron-server/types"
)

type Company struct {
	CollectionDocument `bson:",inline"`
	types.Coordinates

	Name        string             `json:"name"         bson:"name,omitempty"`
	OneLiner    string             `json:"one_liner"    bson:"one_liner,omitempty"`
	Description string             `json:"description"  bson:"description,omitempty"`
	ImageURL    string             `json:"image_url"    bson:"image_url,omitempty"`
	CountryCode string             `json:"country_code" bson:"country_code,omitempty"`
	CityID      primitive.ObjectID `json:"city_id"      bson:"city_id,omitempty"`
}
