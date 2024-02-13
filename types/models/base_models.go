package models

// images of different formats, videos, etc
type Asset struct {
	SqlDocument `       json:",inline,omitempty"    db:",inline"`
	Type        string `json:"_type,omitempty"      db:"_type"`
	Reference   string `json:"_reference,omitempty" db:"_reference"`
	Url         string `json:"url,omitempty"        db:"url"`
	Format      string `json:"format,omitempty"     db:"format"`
	Alt         string `json:"alt,omitempty"        db:"alt"`
}

// Pricing information
type Price struct {
	SqlDocument
	Currency float32 `json:"_currency,omitempty"`

	Title       int     `json:"title,omitempty"`
	Amount      float32 `json:"amount,omitempty"`
	Description string  `json:"description,omitempty"`
}
