package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"sytron-server/api/middleware"
	"sytron-server/api/routes"
	"sytron-server/constants"
)

func main() {
	// database
	// defer conn.Close()

	// Server
	port := "8000"
	if constants.PORT != "" {
		port = constants.PORT
	}

	app := fiber.New()
	app.Use(cors.New())

	if constants.ENVIRONMENT == "development" {
		app.Use(func(c *fiber.Ctx) error {
			log.Printf("Fetching %v\n", c.Request().URI())
			return c.Next()
		})
	}

	routes.InitRoutes(app)

	// backofficers are only allowed access here
	app.Use(middleware.InitJWTAuth())

	// TODO define edpoint level validation for other roles
	routes.InitProtectedRoutes(app)

	log.Fatal(app.Listen(":" + port))
}
