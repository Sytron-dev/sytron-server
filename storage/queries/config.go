package queries

import (
	"context"

	"github.com/jackc/pgx/v5"

	"sytron-server/types/models"
)

func GetCountries() (data []models.Country, err error) {
	query := `
    SELECT name, iso2, en_label, lat, lon, currency_code FROM countries 
    WHERE iso2 IS NOT NULL
  `
	rows, err := pgxConn.Query(context.TODO(), query)
	data, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Country])
	return
}

func GetCities(country string) (data []models.City, err error) {
	query := `
    SELECT _country_iso2, label, value, lon, lat FROM cities 
		`

	if country == "" {
		data = []models.City{}
		return
	}

	if country != "all" {
		query += " WHERE _country_iso2 = '" + country + "'"
	}

	rows, err := pgxConn.Query(context.TODO(), query)
	data, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.City])
	return
}
