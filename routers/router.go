package routers

import (
	"sytron-server/routers/confs"
	"sytron-server/routers/users"

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
	confs.RegisterConfRoutes(router)
	users.RegisterUsersRoutes(router)

	return router
}
