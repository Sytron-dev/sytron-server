package models

// images of different formats, videos, etc
type Asset struct {
	SqlDocument `       json:",inline" db:",inline"`
	Url         string `json:"url"                  bson:"url,omitempty"`
	Format      string `json:"format"               bson:"format,omitempty"`
	Alt         string `json:"alt"                  bson:"alt,omitempty"`
}

// Pricing information
type Price struct {
	SqlDocument
	Currency float32 `json:"_currency,omitempty"`

	Title       int     `json:"title,omitempty"`
	Amount      float32 `json:"amount,omitempty"`
	Description string  `json:"description,omitempty"`
}
