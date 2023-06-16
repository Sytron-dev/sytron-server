package confs

import "github.com/gin-gonic/gin"

func RegisterConfRoutes(router *gin.Engine) {
	GetCountries(router)
	CreateCountry(router)
	CreateLocation(router)
	
}
