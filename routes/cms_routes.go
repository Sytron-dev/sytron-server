package routes

import (
	"github.com/gin-gonic/gin"

	controller "sytron-server/controllers/cms_controller"
)

func initCMSRoutes(router *gin.Engine) {
	// destinations
	router.GET("/destinations", controller.GetDestinations())
	router.GET("/destinations/:id", controller.GetSingleDestination())
}

func initProtectedCMSRoutes(router *gin.Engine) {
	// destinations
	router.POST("/destinations", controller.CreateDestination())
	router.PUT("/destinations/:id", controller.UpdateDestination())
	router.PUT("/destinations/:id/upload/image", controller.UploadDestinationImage())
	router.DELETE("/destinations/:id", controller.DeleteDestination())
}
