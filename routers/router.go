package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitRouters() *gin.Engine {

	// load .env

	if err := godotenv.Load(); err != nil {
		println(err)
	}

	router := gin.New()

	// Register routes
	RegisterDestinationRoutes(router)

	return router
}
