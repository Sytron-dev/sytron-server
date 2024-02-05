package queries

import (
	"context"
	"sytron-server/types/models"

	"github.com/jackc/pgx/v5"
)

// creates a new asset
func CreateAsset(data models.Asset) (asset models.Asset, err error) {
	query := `
		INSERT INTO assets
		(reference, type, url, format, alt)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *
	`

	row, err := pgxConn.Query(
		context.TODO(),
		query,
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
