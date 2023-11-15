package models

type Country struct {
	ID      int      `json:"_id,omitempty"`
	Name    string   `json:"name,omitempty"     db:"name"`
	EnLabel string   `json:"en_label,omitempty"`
	ISO2    *string  `json:"value,omitempty"    db:"iso2"`
	ISO3    *string  `json:"iso3,omitempty"`
	Lat     *float64 `json:"lat,omitempty"`
	Lon     *float64 `json:"lon,omitempty"`
}

type City struct {
	ID          int     `json:"_id,omitempty"`
	CountryISO2 string  `json:"_country,omitempty" db:"_country_iso2"`
	Label       string  `json:"label,omitempty"`
	Value       string  `json:"value,omitempty"`
	Lat         float64 `json:"lat,omitempty"`
	Lon         float64 `json:"lon,omitempty"`
}
