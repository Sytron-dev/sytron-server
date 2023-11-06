package config

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"sytron-server/constants"
	"sytron-server/storage"
	"sytron-server/storage/queries"
	"sytron-server/types"
	"sytron-server/types/models"
)

func getCollection(collectionName string) *mongo.Collection {
	return storage.Client.Database(constants.CONFIGS_DATABASE_NAME).Collection(collectionName)
}

func GetCountries() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		if data, err := queries.GetCountries(); err != nil {
			ctx.Status(fiber.ErrInternalServerError.Code)
			return ctx.JSON(types.ErrorResponse{
				Message: "Oops",
				Error:   err,
			})
		} else {
			return ctx.JSON(data)
		}
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
			return ctx.JSON(types.ErrorResponse{
				Message: "Error fetching city",
				Error:   err,
			})
		}

		var cities []models.City
		cursor.All(context.TODO(), &cities)
		return ctx.JSON(cities)
	}
}
