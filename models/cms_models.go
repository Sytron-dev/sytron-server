package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"sytron-server/types"
)

// Location is a geographical/geopolitical area like a city

type Location struct {
	CollectionDocument `bson:",inline"`
	types.Coordinates  `bson:",omitempty" json:"coordinates"`

	Name        string             `bson:"name,omitempty"         json:"name"`
	OneLiner    string             `bson:"one_liner,omitempty"    json:"one_liner"`
	Description string             `bson:"description,omitempty"  json:"description"`
	ImageURL    string             `bson:"image_url,omitempty"    json:"image_url"`
	CountryCode string             `bson:"country_code,omitempty" json:"country_code"`
	CityID      primitive.ObjectID `bson:"city_id,omitempty"      json:"city_id"`
}

// Destination represents a tourist destination
// Tourist destinations can be a country, city or a place

type Destination struct {
	Location `bson:",inline"`
}
