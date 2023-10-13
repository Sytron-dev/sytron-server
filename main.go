package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"

	middleware "sytron-server/middleware"
	routes "sytron-server/routes"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Failed to load environment variables")
	}

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

	routes.InitRoutes(router) // open routes
	router.Use(middleware.Authentication())
	routes.InitProtectedRoutes(router) // require authorization

	router.Run(":" + port)
}
