package destinations_controller

import (
	"context"
	"fmt"
	"net/http"
	"sytron-server/database"
	"sytron-server/helpers/logger"
	"sytron-server/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateDestination() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// get json data
		var body models.Destination
		if err := ctx.BindJSON(&body); err != nil {
			logger.Handle(err, "Decoding destination json")
			resErr := models.ErrorResponse{
				Message: "There's a problem with your request body",
				Error:   err,
			}
			ctx.JSON(http.StatusBadRequest, resErr)
			return
		}

		collection := database.GetCollection(DESTINATIONS_COLLECTION)

		body.ID = primitive.NewObjectID()
		_, err := collection.InsertOne(context.TODO(), body)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Error writing to database",
				Error:   err,
			})
			return
		}

		ctx.JSON(http.StatusOK, body)
	}
}

func GetDestinations() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get destinations collection
		collection := database.GetCollection(DESTINATIONS_COLLECTION)

		// Find all destinations
		cursor, err := collection.Find(context.TODO(), bson.D{{}})
		fmt.Println(err)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed while reading database",
				Error:   err,
			})
		}

		// Decode results into slice
		var destinations []models.Destination
		if err := cursor.All(context.TODO(), &destinations); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed transforming data",
				Error:   err,
			})
		}

		ctx.JSON(http.StatusOK, destinations)
	}
}

func GetSingleDestination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get _id param
		id := ctx.Params.ByName("id")
		_id, _ := primitive.ObjectIDFromHex(id)

		// Get destinations collection
		collection := database.GetCollection(DESTINATIONS_COLLECTION)
		filter := bson.D{{Key: "_id", Value: _id}}

		// Decode results into slice
		var destination models.Destination
		if err := collection.FindOne(context.TODO(), filter).Decode(&destination); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed to read from database",
				Error:   err,
			})
			return
		}
		ctx.JSON(http.StatusOK, destination)

	}
}

func UpdateDestination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Update one destination")
	}
}

func DeleteDestination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		collection := database.GetCollection(DESTINATIONS_COLLECTION)

		if _, err := collection.DeleteOne(ctx.Request.Context(), bson.M{"_id": id}); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed deleting destination",
				Error:   err,
			})
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
