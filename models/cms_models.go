package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Location is a geographical/geopolitical area like a city

type Location struct {
	CollectionDocument `bson:",inline"`

	Name        string             `json:"name"         bson:"name,omitempty"`
	Latitude    float64            `json:"lat"          bson:"lat,omitempty"`
	Longitude   float64            `json:"lon"          bson:"lon,omitempty"`
	Description string             `json:"description"  bson:"description,omitempty"`
	ImageURL    string             `json:"image_url"    bson:"image_url,omitempty"`
	CountryCode string             `json:"country_code" bson:"country_code,omitempty"`
	CityID      primitive.ObjectID `json:"city_id"      bson:"city_id,omitempty"`
}

// Destination represents a tourist destination
// Tourist destinations can be a country, city or a place

type Destination struct {
	Location `bson:",inline"`
}
