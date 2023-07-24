package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Destination represents a tourist destination
type Destination struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Name        string             `json:"name" bson:"name"`
	Latitude    float64            `json:"lat" bson:"lat"`
	Longitude   float64            `json:"lon" bson:"lon"`
	Description string             `json:"description" bson:"description"`
	ImageURL    string             `json:"image_url" bson:"image_url"`
	CountryCode string             `json:"country_code" bson:"country_code"`
}
