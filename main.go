package main

import (
	"os"

	middleware "sytron-server/middleware"
	routes "sytron-server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.InitRoutes(router)

	// Allow cors
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	router.Use(middleware.Authentication())

	router.Run(":" + port)
}
