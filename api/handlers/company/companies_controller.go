package company

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"sytron-server/storage/queries"
	"sytron-server/types"
	"sytron-server/types/models"
)

func CreateCompany() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		// get json data
		var body models.Company

		if err := ctx.BodyParser(&body); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(types.ErrorResponse{
				Message:  "There's a problem with your request body",
				Metadata: err,
			})

		}

		if res, err := queries.CreateCompany(body); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed writing to database",
				Metadata: err.Error(),
			})
		} else {
			return ctx.JSON(res)
		}
	}
}

func GetCompanies() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		if companies, err := queries.GetCompanies(); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed while reading database",
				Metadata: err.Error(),
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

		if company, err := queries.FindOneCompany(id); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed while reading database",
				Metadata: err.Error(),
			})
		} else {
			return ctx.JSON(company)
		}
	}
}

func UpdateCompany() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		// Get updated company from request body
		var company models.Company
		if err := ctx.BodyParser(&company); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed parsing request body",
				Metadata: err,
			})
		}

		if updatedCompany, err := queries.UpdateCompany(id, company); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed updating company",
				Metadata: err,
			})
		} else {
			return ctx.JSON(updatedCompany)
		}
	}
}
