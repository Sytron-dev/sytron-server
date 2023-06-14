package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterDestinationRoutes(router *gin.Engine) {

	router.GET("/destinations", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"live": "ok"})
	},
	)
}
