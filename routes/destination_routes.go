package routes

import (
	destinations "sytron-server/controllers/destinations_controller"

	"github.com/gin-gonic/gin"
)

func initDestinationRoutes(incomingRoutes *gin.Engine) {
	// destinations CRUD
	incomingRoutes.POST("/destinations", destinations.CreateDestination())
	incomingRoutes.GET("/destinations", destinations.GetDestinations())
	incomingRoutes.GET("/destinations/:id", destinations.GetSingleDestination())
	incomingRoutes.PUT("/destinations/:id", destinations.UpdateDestination())
	incomingRoutes.DELETE("/destinations/:id", destinations.DeleteDestination())
}
