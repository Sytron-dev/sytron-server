package routes

import (
	controller "sytron-server/controllers/cms_controller"

	"github.com/gin-gonic/gin"
)

func initDestinationRoutes(incomingRoutes *gin.Engine) {
	// destinations CRUD
	incomingRoutes.POST("/destinations", controller.CreateDestination())
	incomingRoutes.GET("/destinations", controller.GetDestinations())
	incomingRoutes.GET("/destinations/:id", controller.GetSingleDestination())
	incomingRoutes.PUT("/destinations/:id", controller.UpdateDestination())
	incomingRoutes.PUT("/destinations/:id/upload/image", controller.UploadDestinationImage())
	incomingRoutes.DELETE("/destinations/:id", controller.DeleteDestination())
}
