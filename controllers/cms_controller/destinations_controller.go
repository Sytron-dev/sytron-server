package cms_controller

import (
	"context"
	"fmt"
	"net/http"
	"sytron-server/controllers/uploads_controller"
	"sytron-server/database"
	"sytron-server/helpers/logger"
	"sytron-server/models"
	"sytron-server/resolvers"
	"sytron-server/storage"

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

		collection := database.GetCollection(database.DESTINATIONS_COLLECTION)

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
		if destinations, err := resolvers.DestinationResolver.FindMany(database.PaginationOptions{}); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed while reading database",
				Error:   err,
			})
			return
		} else {
			ctx.JSON(http.StatusOK, destinations)
		}
	}
}

func GetSingleDestination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get _id param
		id := ctx.Params.ByName("id")

		if destination, err := resolvers.DestinationResolver.FindOneByID(id); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Error finding document",
				Error:   err,
			})
		} else {
			ctx.JSON(http.StatusOK, destination)
		}
	}
}

func updateOneDestination(id string, data models.Destination) (*models.Destination, error) {
	_id, _ := primitive.ObjectIDFromHex(id)

	tmpDest := resolvers.DestinationResolver.TODO()
	tmpDest.SetID(id)

	collection := database.GetCollection(database.DESTINATIONS_COLLECTION)
	filter := bson.D{{Key: "_id", Value: _id}}

	update := bson.M{"$set": data}

	if _, err := collection.UpdateOne(context.TODO(), filter, update); err != nil {
		return nil, err
	}

	return resolvers.DestinationResolver.FindOneByID(id)
}

func UpdateDestination() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Params.ByName("id")

		// Get updated destination from request body
		var dest models.Destination

		if err := ctx.ShouldBindJSON(&dest); err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
				Message: "Failed parsing request body",
				Error:   err,
			})
			return
		}

		updatedDest, err := updateOneDestination(id, dest)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed updating destination :(",
				Error:   err,
			})
			return
		}

		ctx.JSON(http.StatusOK, updatedDest)
	}
}

func UploadDestinationImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Params.ByName("id")

		fileName := fmt.Sprintf("destinations/%v/image", id)

		imageUrl, errResponse := uploads_controller.UploadFile(ctx, "image", storage.CMSBucketHandle, fileName)
		if errResponse != nil {
			ctx.JSON(http.StatusInternalServerError, errResponse)
			return
		}

		if updatedDest, err := updateOneDestination(id, models.Destination{Location: &models.Location{ImageURL: *imageUrl}}); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed updating destination",
				Error:   err,
			})
		} else {
			ctx.JSON(http.StatusOK, updatedDest)
		}
	}
}

func DeleteDestination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		collection := database.GetCollection(database.DESTINATIONS_COLLECTION)

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