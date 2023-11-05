package company_controllers

import (
	"net/http"

	"sytron-server/database"
	"sytron-server/models"
	"sytron-server/resolvers"
	"sytron-server/types"

	"github.com/gofiber/fiber/v2"
)

func CreateCompany() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		// get json data
		var body models.Company

		if err := ctx.BodyParser(&body); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(models.ErrorResponse{
				Message: "There's a problem with your request body",
				Error:   err,
			})

		}

		body.SetID()
		body.InsertTime()

		if res, err := resolvers.CompanyResolver.InsertOne(body); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed writing to database",
				Error:   err,
			})
		} else {
			return ctx.JSON(res)
		}
	}
}

func GetCompanies() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		if companies, err := resolvers.CompanyResolver.FindMany(database.PaginationOptions{}); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed while reading database",
				Error:   err,
			})
		} else {
			return ctx.JSON(companies)
		}
	}
}

func GetSingleCompany() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		// Get _id param
		id := ctx.Params("id")

		if company, err := resolvers.CompanyResolver.FindOneByID(id); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed while reading database",
				Error:   err,
			})
		} else {
			return ctx.JSON(company)
		}
	}
}

func updateOneCompany(id string, data models.Company) (*models.Company, error) {
	data.UpdateTime()
	return resolvers.CompanyResolver.UpdateOne(id, data)
}

func UpdateCompany() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		// Get updated company from request body
		var company models.Company
		if err := ctx.BodyParser(&company); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed parsing request body",
				Error:   err,
			})
		}

		if updatedCompany, err := updateOneCompany(id, company); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed updating company",
				Error:   err,
			})
		} else {
			return ctx.JSON(updatedCompany)
		}
	}
}
