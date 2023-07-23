package routes

import (
	controller "sytron-server/controllers"

	"github.com/gin-gonic/gin"
)

// UserRoutes function
func initUserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}
