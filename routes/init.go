package routes

import "github.com/gin-gonic/gin"

func InitRoutes(router *gin.Engine) {
	initConfigRoutes(router)
	initUserRoutes(router)
	initDestinationRoutes(router)
	InitUploadRoutes(router)
}
