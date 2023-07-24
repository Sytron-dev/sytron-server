package configs_controller

import (
	"context"
	"net/http"
	"sytron-server/database"
	"sytron-server/helpers/logger"
	"sytron-server/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCollection(collectionName string) *mongo.Collection {
	return database.Client.Database(DATABASE_NAME).Collection(collectionName)
}

func GetCountries() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		collection := getCollection(COUNTRIES_COLLECTION)
		filter := bson.D{{}}
		option := options.Find()

		// Find all countries
		cursor, err := collection.Find(context.TODO(), filter, option)
		if err != nil {
			logger.Handle(err, "Fetching countries array")
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Internal server error",
				Error:   err,
			})
		} else {
			var countries []models.Country
			cursor.All(context.TODO(), &countries)
			ctx.JSON(http.StatusOK, countries)
		}
	}
}

type getCitiesRequestBody struct {
	CountryCode string `json:"country_code"`
}

func GetCities() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the request body
		var body getCitiesRequestBody
		if err := ctx.BindJSON(&body); err != nil {
			logger.Handle(err, "Fetching cities array")
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
				Message: "Check request body country_code",
				Error:   err,
			})
			return
		}

		collection := getCollection(CITIES_COLLECTION)
		filter := bson.D{{Key: "country_code", Value: body.CountryCode}}
		options := options.Find()
		options.SetLimit(200)

		// Find all cities
		cursor, err := collection.Find(context.TODO(), filter, options)
		if err != nil {
			logger.Handle(err, "Fetching cities array")
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Internal server error",
				Error:   err,
			})
		} else {
			var cities []models.City
			cursor.All(context.TODO(), &cities)
			ctx.JSON(http.StatusOK, cities)
		}
	}
}
