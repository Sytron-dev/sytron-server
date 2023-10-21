package routes

import (
	"github.com/gin-gonic/gin"
)

// Open routes, do not require authorization
func InitRoutes(router *gin.Engine) {
	initAuthRoutes(router)
	initConfigRoutes(router)
}

// Require Authorization
func InitProtectedRoutes(router *gin.Engine) {
	initProtectedAuthRoutes(router)
	initUserRoutes(router)
	initDestinationRoutes(router)
	initProtectedEventRoutes(router)
}
