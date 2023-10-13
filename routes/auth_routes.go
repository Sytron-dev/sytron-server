package routes

import (
	"github.com/gin-gonic/gin"

	controller "sytron-server/controllers/auth_controller"
)

func initAuthRoutes(router *gin.Engine) {
	router.POST("/auth/users", controller.LoginUser())

	// backoffice
	router.POST("/auth/backofficer/create", controller.CreateBackOfficer)
	router.POST("/auth/backofficer/login", controller.LoginBackOfficer)
}

// requires user to be logged in
func initProtectedAuthRoutes(router *gin.Engine) {
}
