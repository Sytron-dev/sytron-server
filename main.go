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

	// Allow cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowHeaders:     []string{"*"},
		AllowWildcard:    true,
		AllowFiles:       true,
	}))
	routes.InitRoutes(router)

	// @vin, I will need to move this up when I implement authorization
	router.Use(middleware.Authentication())

	router.Run(":" + port)

	} 