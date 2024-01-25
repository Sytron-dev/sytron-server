package models

import (
	"sytron-server/types"
)

// bed and breakfast

type BNB struct {
	// keys
	SqlDocument `json:",inline" db:",inline"`
	City        *int64      `json:"_city"    db:"_city,omitempty"`
	Country     *string     `json:"_country" db:"_country,omitempty"`
	Company     *types.UUID `json:"_company" db:"_company,omitempty"`
	Currency    *string     `json:"currency" db:"_currency,omitempty"`

	Name        string   `json:"name"        db:"name"`
	Description *string  `json:"description" db:"description"`
	ImageUrl    *string  `json:"image_url"   db:"image_url"`
	Price       *float32 `json:"price"       db:"price,omitempty"`
	Assets      []Asset  `json:"assets"      db:"assets"`
	Archived    bool     `json:"archived"    db:"archived"`

	types.Coordinates `json:"coordinates" db:",inline"`
}
