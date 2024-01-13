package queries

import (
	"context"

	"github.com/jackc/pgx/v5"

	"sytron-server/constants"
	"sytron-server/types/models"
)

/*
allow backofficers to invite new users
backofficer logins should be updated

Create, Read, Update
*/

func createAuthCredential(
	table string,
	data models.AuthCredential,
) (credential models.AuthCredential, err error) {
	query := `
		INSERT INTO ` + table +
		` 
		(credential_value, credential_type, password)
		VALUES ($1, $2, $3)
	`

	row, err := pgxConn.Query(
		context.TODO(), query,
		data.Value,
		data.Type,
		data.Password,
	)

	credential, err = pgx.CollectOneRow(row, pgx.RowToStructByNameLax[models.AuthCredential])
	return
}

func getAuthCredentials(table string) (credentials []models.AuthCredential, err error) {
	query := `
		SELECT * FROM ` + table
	rows, err := pgxConn.Query(context.TODO(), query)
	credentials, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.AuthCredential])
	return
}

func getSingleEmailAndPasswordCredentials(table string, data models.AuthCredential) (credential models.AuthCredential, err error) {
	query := `
		SELECT * FROM ` + table + `
		WHERE credential_type = $1 AND credential_value = $2
		LIMIT 1
	`
	row, err := pgxConn.Query(context.TODO(), query, data.Type, data.Value)

	credential, err = pgx.CollectOneRow(row, pgx.RowToStructByNameLax[models.AuthCredential])
	return
}

func CreateBackofficerAuthCredentials(data models.AuthCredential) (credential models.AuthCredential, err error) {
	return createAuthCredential(constants.TABLE_AUTH_BACKOFFICERS, data)
}

func GetBackofficerAuthCredential(data models.AuthCredential) (credential models.AuthCredential, err error) {
	return getSingleEmailAndPasswordCredentials(constants.TABLE_AUTH_BACKOFFICERS, models.AuthCredential{
		Type:  constants.AUTH_TYPE_EMAIL_AND_PASSWORD,
		Value: data.Value,
	})
}
