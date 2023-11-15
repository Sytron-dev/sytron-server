package queries

import (
	"context"
	"sytron-server/types/models"

	"github.com/jackc/pgx/v5"
)

func GetCountries() (data []models.Country, err error) {
	query := `
    SELECT name,iso2 FROM countries 
    WHERE iso2 IS NOT NULL
  `
	rows, err := pgxConn.Query(context.TODO(), query)
	data, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Country])
	return
}

func GetCities() (data []models.City, err error) {
	query := `
    SELECT _country_iso2, label, value FROM cities 
    WHERE _country_iso2 = 'KE'
  `
	rows, err := pgxConn.Query(context.TODO(), query)
	data, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.City])
	return
}
