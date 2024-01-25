package backoffice

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"sytron-server/storage"
	"sytron-server/storage/resolvers"
	"sytron-server/types"
)

// get a list of all backoffice credentials info
func GetBackOfficers() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		if backofficers, err := resolvers.BackOfficersResolver.FindMany(storage.PaginationOptions{}); err != nil {
			// failure
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message: "Failed while reading database",
				Error:   err,
			})
		} else {
			// success
			return ctx.JSON(backofficers)
		}
	}
}
