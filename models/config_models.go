package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Country struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id", omitempty`
	Name         string             `bson:"name" json:"name"`
	EnLabel      string             `bson:"en_label" json:"en_label"`
	Lat          float64            `bson:"lat" json:"lat"`
	Lon          float64            `bson:"lon" json:"lon"`
	CurrencyCode string             `bson:"currency_code" json:"currency_code"`
}

type City struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	CountryCode string             `bson:"country_code" json:"country_code"`
	Label       string             `bson:"label" json:"label"`
	EnLabel     string             `bson:"en_label" json:"en_label"`
	Value       string             `bson:"value" json:"value"`
	Lat         float64            `bson:"lat" json:"lat"`
	Lon         float64            `bson:"lon" json:"lon"`
}