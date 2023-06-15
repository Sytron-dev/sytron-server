package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAccomodationRoutes(router *gin.Engine) {

	router.GET("/accomodations", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"live": "ok"})
	},
	)
}
