package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

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

	app := fiber.New()
	app.Use(cors.New())

	routes.InitRoutes(app)

	// TODO add endpoint-level protection
	routes.InitProtectedRoutes(app)

	log.Fatal(app.Listen(":" + port))
}
