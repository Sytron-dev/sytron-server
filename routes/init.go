package routes

import (
	"github.com/gin-gonic/gin"
)

// Open routes, do not require authorization
func InitRoutes(router *gin.Engine) {
	initAuthRoutes(router)
	initUserRoutes(router)
	initConfigRoutes(router)
	initCMSRoutes(router)
	initEventRoutes(router)
}

// Require Authorization
func InitProtectedRoutes(router *gin.Engine) {
	initProtectedAuthRoutes(router)
	initProtectedCMSRoutes(router)
	initProtectedCompanyRoutes(router)
	initProtectedEventRoutes(router)
}
