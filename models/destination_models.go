package models

// Destination represents a tourist destination
// Tourist destinations can be a country, city or a location

type Destination struct {
	*CollectionDocument

	Name        string  `json:"name" bson:"name,omitempty"`
	Latitude    float64 `json:"lat" bson:"lat,omitempty"`
	Longitude   float64 `json:"lon" bson:"lon,omitempty"`
	Description string  `json:"description" bson:"description,omitempty"`
	ImageURL    string  `json:"image_url" bson:"image_url,omitempty"`
	CountryCode string  `json:"country_code" bson:"country_code,omitempty"`
}
