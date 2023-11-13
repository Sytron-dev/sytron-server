package models

import (
	"sytron-server/types"
)

type Asset struct {
	Url    string `bson:"url,omitempty"    json:"url"`
	Format string `bson:"format,omitempty" json:"format"`
	Alt    string `bson:"alt,omitempty"    json:"alt"`
}

// Location is a geographical/geopolitical area like a city

type Location struct {
	SqlDocument       `json:",inline"     db:",inline"`
	types.Coordinates `json:"coordinates" db:",inline"`

	// foreign keys
	CountryCode string `json:"_country" db:"_country"`
	CityID      string `json:"_city"    db:"_city"`

	Name        string `json:"name"        db:"name"`
	OneLiner    string `json:"one_liner"   db:"one_liner"`
	Description string `json:"description" db:"description"`
	ImageURL    string `json:"image_url"   db:"image_url"`
}

// Destination represents a tourist destination
// Tourist destinations can be a country, city or a place

type Destination struct {
	Location `json:",inline" db:",inline"`
}
