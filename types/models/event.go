package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"sytron-server/types"
)

type Event struct {
	CollectionDocument `bson:",inline"     json:",inline"`
	types.Coordinates  `bson:"coordinates" json:"coordinates"`

	CompanyID   primitive.ObjectID `bson:"company_id,omitempty"   json:"company_id"`
	LocationID  primitive.ObjectID `bson:"location_id,omitempty"  json:"location_id"`
	CountryCode string             `bson:"country_code,omitempty" json:"country_code"`

	Title        string             `bson:"title,omitempty"       json:"title"`
	StartTime    primitive.DateTime `bson:"start_time,omitempty"  json:"start_time"`
	EndTime      primitive.DateTime `bson:"end_time,omitempty"    json:"end_time"`
	OneLiner     string             `bson:"one_liner,omitempty"   json:"one_liner"`
	Description  string             `bson:"description,omitempty" json:"description"`
	HeroImageUrl string             `bson:"image_url,omitempty"   json:"image_url"`
	Prices       []EventPrice       `bson:"prices"                json:"prices"`
	Assets       []Asset            `bson:"assets"                json:"assets"`
}

type EventPrice struct {
	Currency    string  `bson:"currency,omitempty"    json:"currency"`
	Amount      float32 `bson:"amount,omitempty"      json:"amount"`
	Title       string  `bson:"title,omitempty"       json:"title"`
	Description string  `bson:"description,omitempty" json:"description"`
}
