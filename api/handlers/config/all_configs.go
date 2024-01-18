package config

import (
	"github.com/gofiber/fiber/v2"

	"sytron-server/storage/queries"
	"sytron-server/types"
)

func GetAllConfigs() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		// get countries
		countries, err := queries.GetCountries()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{Error: err})
		}

		// get all cities
		cities, err := queries.GetCities("all")
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
				Error:    err,
				Metadata: err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
			"countries": countries,
			"cities":    cities,
		})
	}
}
