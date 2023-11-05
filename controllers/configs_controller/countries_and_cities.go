package configs_controller

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"sytron-server/constants"
	"sytron-server/database"
	"sytron-server/models"
	"sytron-server/types"
)

func getCollection(collectionName string) *mongo.Collection {
	return database.Client.Database(constants.CONFIGS_DATABASE_NAME).Collection(collectionName)
}

func GetCountries() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		collection := getCollection(constants.COUNTRIES_COLLECTION_NAME)
		filter := bson.D{{}}
		option := options.Find()

		// Find all countries
		cursor, err := collection.Find(context.TODO(), filter, option)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Internal server error",
				Error:   err,
			})
		}

		var countries []models.Country
		cursor.All(context.TODO(), &countries)
		return ctx.JSON(countries)
	}
}

func GetCities() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		// Get the request body

		countryCode := ctx.Query("country_code", "KE")

		collection := getCollection(constants.CITIES_COLLECTION_NAME)
		filter := bson.D{{Key: "country_iso2", Value: countryCode}}
		options := options.Find()
		options.SetLimit(200)

		// Find all cities
		cursor, err := collection.Find(context.TODO(), filter, options)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Error fetching city",
				Error:   err,
			})
		}

		var cities []models.City
		cursor.All(context.TODO(), &cities)
		return ctx.JSON(cities)
	}
}
