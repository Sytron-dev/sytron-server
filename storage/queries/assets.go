package queries

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"sytron-server/types/models"
)

// creates a new asset
func CreateAsset(data models.Asset) (asset models.Asset, err error) {
	query := `
		INSERT INTO assets
		(_id, _reference, _type, url, format, alt)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *
	`

	fmt.Println(data)

	row, err := pgxConn.Query(
		context.TODO(),
		query,
		data.ID,
		data.Reference,
		data.Type,
		data.Url,
		data.Format,
		data.Alt,
	)
	if err != nil {
		return
	}

	asset, err = pgx.CollectOneRow(row, pgx.RowToStructByNameLax[models.Asset])
	return
}

// update the asset
func UpdateAsset(id string, data models.Asset) (asset models.Asset, err error) {
	query := `
		UPDATE assets
		SET url = $1, format = $2, alt = $3
		WHERE _id = $4
		RETURNING *
	`

	row, err := pgxConn.Query(
		context.TODO(),
		query,
		data.Url,
		data.Format,
		data.Alt,
		id,
	)
	if err != nil {
		return
	}

	asset, err = pgx.CollectOneRow(row, pgx.RowToStructByNameLax[models.Asset])

	return
}

// deletes all assets that match the reference
func DeleteAssets(_reference string, _type string) (err error) {
	query := `
		DELETE FROM assets
		WHERE
			_reference = $1 AND _type = $2
	`
	_, err = pgxConn.Query(context.TODO(), query, _reference, _type)
	return
}
