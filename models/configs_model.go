package models

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Country struct {
	ID    primitive.ObjectID `bson:"_id"`
	Value string             `json:"value"`
	Label string             `json:"label"`
	Lat   float32            `json:"lat"`
	Lon   float32            `json:"lon"`
	// Emoji        string             `json:"emoji"`
	// CurrencyID   string             `json:"currency_id"`
	// CurrencyName string             `json:"currency_name"`
}

type Location struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Lat       float32   `json:"lat"`
	Lon       float32   `json:"lon"`
	CountryID string    `json:"country_id"`
}

type Accomodation struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Price       float32   `json:"price"`
	LocationID  string    `json:"location_id"`
	Rating      float32   `json:"rating"`
	Rooms       int       `json:"rooms"`
	Description string    `json:"description"`
}
