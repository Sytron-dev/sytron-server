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
				Message: err.Error(),
				Error:   err,
			})
		} else {
			return ctx.JSON(data)
		}
	}
}

func GetCities() types.HandlerFunc {
	return func(ctx *fiber.Ctx) (err error) {
		// Get the request body
		cities, err := queries.GetCities()
		if err != nil {
			return ctx.Status(fiber.ErrInternalServerError.Code).JSON(types.ErrorResponse{
				Message: "Oops",
				Error:   err,
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(cities)
	}
}
