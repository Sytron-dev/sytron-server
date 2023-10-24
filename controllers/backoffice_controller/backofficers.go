package backoffice_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sytron-server/database"
	"sytron-server/models"
	"sytron-server/resolvers"
)

// get a list of all backoffice credentials info
func GetBackOfficers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if backofficers, err := resolvers.BackOfficersResolver.FindMany(database.PaginationOptions{}); err != nil {
			// failure
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed while reading database",
				Error:   err,
			})
		} else {
			// success
			ctx.JSON(http.StatusOK, backofficers)
		}
	}
}
