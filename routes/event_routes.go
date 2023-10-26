package routes

import (
	"github.com/gin-gonic/gin"

	controller "sytron-server/controllers/event_controllers"
)

func initEventRoutes(router *gin.Engine) {
	router.GET("/events/:id", controller.GetSingleEvent())
}

func initProtectedEventRoutes(router *gin.Engine) {
	router.POST("/events", controller.CreateEvent())
}
