package config

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

	"sytron-server/constants"
	"sytron-server/storage"
	"sytron-server/storage/queries"
	"sytron-server/types"
)

func getCollection(collectionName string) *mongo.Collection {
	return storage.Client.Database(constants.CONFIGS_DATABASE_NAME).Collection(collectionName)
}

func GetCountries() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		if data, err := queries.GetCountries(); err != nil {
			ctx.Status(fiber.ErrInternalServerError.Code)
			return ctx.JSON(types.ErrorResponse{
				Message:  err.Error(),
				Metadata: err,
			})
		} else {
			return ctx.JSON(data)
		}
	}
}

func GetCities() types.HandlerFunc {
	return func(ctx *fiber.Ctx) (err error) {
		// Get the request body
		country := ctx.Query("country")

		cities, err := queries.GetCities(country)
		if err != nil {
			return ctx.Status(fiber.ErrInternalServerError.Code).JSON(types.ErrorResponse{
				Message:  "Oops",
				Metadata: err,
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(cities)
	}
}
