package routes

import (
	configs "sytron-server/controllers/configs_controller"

	"github.com/gin-gonic/gin"
)

func initConfigRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/configs/countries", configs.GetCountries())
}
