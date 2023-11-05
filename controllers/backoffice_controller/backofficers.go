package backoffice_controller

import (
	"net/http"
	"sytron-server/database"
	"sytron-server/models"
	"sytron-server/resolvers"
	"sytron-server/types"

	"github.com/gofiber/fiber/v2"
)

// get a list of all backoffice credentials info
func GetBackOfficers() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		if backofficers, err := resolvers.BackOfficersResolver.FindMany(database.PaginationOptions{}); err != nil {
			// failure
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed while reading database",
				Error:   err,
			})
		} else {
			// success
			return ctx.JSON(backofficers)
		}
	}
}
