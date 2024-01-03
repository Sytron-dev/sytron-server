package queries

import (
	"context"

	"github.com/jackc/pgx/v5"

	"sytron-server/types/models"
)

func CreateBNB(b models.BNB) (bnb models.BNB, err error) {
	query := `
		INSERT INTO bnb
		(name, description, lat, lon, image_url, _currency, price, _company, _country, _city)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING *
	`

	row, err := pgxConn.Query(
		context.TODO(),
		query,
		b.Name,
		b.Description,
		b.Coordinates.Latitude,
		b.Coordinates.Longitude,
		b.Currency,
		b.Price,
		b.ImageUrl, // TODO move this to its own query
		b.Company,
		b.Country,
		b.City,
	)
	if err != nil {
		return
	}

	bnb, err = pgx.CollectOneRow(row, pgx.RowToStructByNameLax[models.BNB])
	return
}

func GetBNBs() (bnbs []models.BNB, err error) {
	query := `
		SELECT * 
		FROM bnb
	`
	rows, err := pgxConn.Query(context.TODO(), query)
	if err != nil {
		return
	}
	bnbs, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.BNB])
	return
}

func FindOneBNB(id string) (bnb models.BNB, err error) {
	query := `
		SELECT * 
		FROM bnb
		WHERE _id = $1
	`
	row, err := pgxConn.Query(context.TODO(), query, id)
	if err != nil {
		return
	}

	bnb, err = pgx.CollectOneRow(row, pgx.RowToStructByNameLax[models.BNB])
	return
}

func UpdateBNB(id string, b models.BNB) (bnb models.BNB, err error) {
	query := `
		UPDATE bnb
		SET 
			name=$2, description=$3, image_url=$4, 
			_country=$5, _city=$6,_currency=$7,
			price=$8,lat=$9, lon=$10, updated_at=$11
		WHERE _id = $1
		RETURNING *
	`

	row, err := pgxConn.Query(
		context.TODO(),
		query,
		id,
		b.Name,
		b.Description,
		b.ImageUrl, // TODO move this to it's own query
		b.Country,
		b.City,
		b.Currency,
		b.Price,
		b.Coordinates.Latitude,
		b.Coordinates.Longitude,
		b.UpdatedAt,
	)
	if err != nil {
		return
	}

	bnb, err = pgx.CollectOneRow(row, pgx.RowToStructByNameLax[models.BNB])
	return
}
