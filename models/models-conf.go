package models

import "github.com/google/uuid"

type Country struct {
	ID           uuid.UUID `json:"id"`
	Value        string    `json:"value"`
	Label        string    `json:"label"`
	Lat          float32   `json:"lat"`
	Lon          float32   `json:"lon"`
	Emoji        string    `json:"emoji"`
	CurrencyID   string    `json:"currency_id"`
	CurrencyName string    `json:"currency_name"`
}

type Location struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Lat       float32   `json:"lat"`
	Lon       float32   `json:"lon"`
	CountryID string    `json:"country_id"`
}

