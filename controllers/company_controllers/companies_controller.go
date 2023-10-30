package company_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sytron-server/database"
	"sytron-server/models"
	"sytron-server/resolvers"
)

func CreateCompany() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get json data
		var body models.Company

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
				Message: "There's a problem with your request body",
				Error:   err,
			})
			return
		}

		body.SetID()
		body.InsertTime()

		if res, err := resolvers.CompanyResolver.InsertOne(body); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed writing to database",
				Error:   err,
			})
			return
		} else {
			ctx.JSON(http.StatusOK, res)
		}
	}
}

func GetCompanies() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if companies, err := resolvers.CompanyResolver.FindMany(database.PaginationOptions{}); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed while reading database",
				Error:   err,
			})
		} else {
			ctx.JSON(http.StatusOK, companies)
		}
	}
}

func GetSingleCompany() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get _id param
		id := ctx.Params.ByName("id")

		if company, err := resolvers.CompanyResolver.FindOneByID(id); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed while reading database",
				Error:   err,
			})
		} else {
			ctx.JSON(http.StatusOK, company)
		}
	}
}

func updateOneCompany(id string, data models.Company) (*models.Company, error) {
	data.UpdateTime()
	return resolvers.CompanyResolver.UpdateOne(id, data)
}

func UpdateCompany() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")

		// Get updated company from request body
		var company models.Company
		if err := ctx.ShouldBindJSON(&company); err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
				Message: "Failed parsing request body",
				Error:   err,
			})
			return
		}

		if updatedCompany, err := updateOneCompany(id, company); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed updating company",
				Error:   err,
			})
			return
		} else {
			ctx.JSON(http.StatusOK, updatedCompany)
		}
	}
}
