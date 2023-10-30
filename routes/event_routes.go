package routes

import (
	"github.com/gin-gonic/gin"

	controller "sytron-server/controllers/event_controllers"
)

func initEventRoutes(router *gin.Engine) {
	router.GET("/events/:id", controller.GetSingleEvent())
	router.GET("/events", controller.GetEvents())
}

func initProtectedEventRoutes(router *gin.Engine) {
	router.POST("/events", controller.CreateEvent())
	router.POST("/events/:id/image", controller.UploadEventHeroImage())
	router.PUT("/events/:id", controller.UpdateEvent())
	router.DELETE("/events/:id", controller.DeleteEvent())
}
