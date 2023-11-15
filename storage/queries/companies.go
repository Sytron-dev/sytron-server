package queries

import (
	"context"
	"sytron-server/types/models"

	"github.com/jackc/pgx/v5"
)

func CreateCompany(b models.Company) (company models.Company, err error) {
	query := `
    INSERT INTO companies
    (name, email, phone, logo_url, _hq)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
  `
	row, err := pgxConn.Query(
		context.TODO(),
		query,
		b.Name,
		b.Email,
		b.Phone,
		b.LogoURL,
		b.HQ,
	)
	if err != nil {
		return
	}

	company, err = pgx.CollectOneRow(row, pgx.RowToStructByNameLax[models.Company])
	return
}

func GetCompanies() (companies []models.Company, err error) {
	query := `
    SELECT * 
    FROM companies
  `
	rows, err := pgxConn.Query(context.TODO(), query)
	if err != nil {
		return
	}

	companies, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Company])
	return
}

func FindOneCompany(id string) (company models.Company, err error) {
	query := `
    SELECT * 
    FROM companies
    WHERE _id = $1
  `
	row, err := pgxConn.Query(context.TODO(), query, id)
	if err != nil {
		return
	}

	company, err = pgx.CollectOneRow(row, pgx.RowToStructByNameLax[models.Company])
	return
}

func UpdateCompany(id string, b models.Company) (company models.Company, err error) {
	query := `
    UPDATE companies
    SET name = $1, email = $2, phone = $3, logo_url = $4, _hq = $5
    WHERE _id = $6
    RETURNING *
  `
	row, err := pgxConn.Query(
		context.TODO(),
		query,
		b.Name,
		b.Email,
		b.Phone,
		b.LogoURL,
		b.HQ,
		id,
	)
	if err != nil {
		return
	}

	company, err = pgx.CollectOneRow(row, pgx.RowToStructByNameLax[models.Company])
	return
}
