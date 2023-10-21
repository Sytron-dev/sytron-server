package event_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEvent() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, "endpoint is open")

	}
}
