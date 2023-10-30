package routes

import (
	"github.com/gin-gonic/gin"

	controller "sytron-server/controllers/company_controllers"
)

func initProtectedCompanyRoutes(router *gin.Engine) {
	router.GET("/companies", controller.GetCompanies())
	router.GET("/companies/:id", controller.GetSingleCompany())
	router.POST("/companies", controller.CreateCompany())
	router.PUT("/companies/:id", controller.UpdateCompany())
}
