package configs_controller

import (
	"context"
	"fmt"
	"net/http"
	"sytron-server/database"
	"sytron-server/helpers/logger"
	"sytron-server/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCollection(collectionName string) *mongo.Collection {
	return database.Client.Database(DATABASE_NAME).Collection(collectionName)
}

func GetCountries() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		collection := getCollection(COUNTRIES_COLLECTION)
		filter := bson.D{{}}

		// Find all countries
		cursor, err := collection.Find(context.TODO(), filter)
		if err != nil {
			logger.Handle(err, "Fetching countries array")
			var resErr models.ErrorResponse
			resErr.Message = "Internal server error"
			ctx.JSON(http.StatusInternalServerError, resErr)
		} else {
			var countries []models.Country
			cursor.All(context.TODO(), &countries)
			fmt.Println("Countries countries")
			ctx.JSON(http.StatusOK, countries)
		}
	}
}
