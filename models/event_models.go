package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"sytron-server/types"
)

type Event struct {
	CollectionDocument `bson:",inline"`
	types.Coordinates  `bson:"coordinates" json:"coordinates"`

	CompanyID   primitive.ObjectID `bson:"company_id"   json:"company_id"`
	LocationID  primitive.ObjectID `bson:"location_id"  json:"location_id"`
	CountryCode string             `bson:"country_code" json:"country_code"`

	Title        string             `bson:"title"       json:"title"`
	StartTime    primitive.DateTime `bson:"start_time"  json:"start_time"`
	EndTime      primitive.DateTime `bson:"end_time"    json:"end_time"`
	OneLiner     string             `bson:"one_liner"   json:"one_liner"`
	Description  string             `bson:"description" json:"description"`
	HeroImageUrl string             `bson:"image_url"   json:"image_url"`
	Prices       []EventPrice       `bson:"prices"      json:"prices"`
	Assets       []Asset            `bson:"assets"      json:"assets"`
}

type EventPrice struct {
	Currency    string  `bson:"currency"    json:"currency"`
	Amount      float32 `bson:"amount"      json:"amount"`
	Title       string  `bson:"title"       json:"title"`
	Description string  `bson:"description" json:"description"`
}
