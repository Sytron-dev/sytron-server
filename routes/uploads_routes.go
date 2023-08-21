package routes

import (
	"github.com/gin-gonic/gin"

	controller "sytron-server/controllers/uploads_controller"
)

func InitUploadRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/uploads/test", controller.TestUploads())
}
