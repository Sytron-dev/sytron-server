package queries

import (
	"sytron-server/storage/models"
	"sytron-server/storage/tables"
)

func GetCountries() (data []models.Country, err error) {
	err = db.From(tables.COUNTRIES).Select("*").Execute(&data)
	return
}

func GetCities() (data []models.City, err error) {
	err = db.From(tables.CITIES).Select("*").Execute(&data)
	return
}
