package queries

import (
	"context"
	"sytron-server/types/models"

	"github.com/jackc/pgx/v5"
)

func GetDestinations() (destinations []models.Destination, err error) {
	query := `
    SELECT * FROM tourist_destinations
  `
	rows, err := pgxConn.Query(context.TODO(), query)
	destinations, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Destination])
	return
}

func FindOneDestination(id string) (destination models.Destination, err error) {
	query := `
    SELECT * FROM tourist_destinations
    WHERE _id = $1
  `
	row, err := pgxConn.Query(context.TODO(), query, id)
	destination, err = pgx.CollectOneRow(row, pgx.RowToStructByNameLax[models.Destination])
	return
}

func CreateDestination(d models.Destination) (destination models.Destination, err error) {
	query := `
    INSERT INTO tourist_destinations
    (name, one_liner, description, image_url, _country, _city, lat, lon)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING *
  `

	row, err := pgxConn.Query(
		context.TODO(),
		query,
		d.Name,
		d.OneLiner,
		d.Description,
		d.ImageURL,
		d.CountryCode,
		d.CityID,
		d.Coordinates.Latitude,
		d.Coordinates.Longitude,
	)
	destination, err = pgx.CollectOneRow(row, pgx.RowToStructByNameLax[models.Destination])

	return
}

func UpdateDestination(
	id string,
	d models.Destination,
) (destination models.Destination, err error) {
	query := `
    UPDATE tourist_destinations
    SET name = $1, one_liner = $2, description = $3, image_url = $4, _country = $5, _city = $6, lat = $7, lon = $8, updated_at=$9
    WHERE _id = $10
    RETURNING *
  `
	d.UpdateTime()

	row, err := pgxConn.Query(
		context.TODO(),
		query,
		d.Name,
		d.OneLiner,
		d.Description,
		d.ImageURL,
		d.CountryCode,
		d.CityID,
		d.Coordinates.Latitude,
		d.Coordinates.Longitude,
		d.UpdatedAt,
		id,
	)
	destination, err = pgx.CollectOneRow(row, pgx.RowToStructByNameLax[models.Destination])

	return
}

func DeleteDestination(id string) (err error) {
	query := `
    DELETE FROM tourist_destinations
    WHERE _id = $1
  `
	_, err = pgxConn.Exec(context.TODO(), query, id)
	return
}
