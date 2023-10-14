package routes

import (
	"github.com/gin-gonic/gin"

	controller "sytron-server/controllers/auth_controller"
)

func initAuthRoutes(router *gin.Engine) {
	router.POST("/auth/users", controller.LoginUser())

	// backoffice
	router.POST("/auth/backofficer/login", controller.LoginBackOfficer)

	// merchant
	router.POST("/auth/merchant/create", controller.CreateMerchantAuth)
	router.POST("/auth/merchant/login", controller.LoginMerchant)
}

// requires user to be logged in
func initProtectedAuthRoutes(router *gin.Engine) {
	// only logged in backofficers can create other backofficers
	router.POST("/auth/backofficer/create", controller.CreateBackOfficeAuth)
}
