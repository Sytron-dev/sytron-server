package cms_controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"sytron-server/constants"
	"sytron-server/controllers/uploads_controller"
	"sytron-server/database"
	"sytron-server/helpers/logger"
	"sytron-server/models"
	"sytron-server/resolvers"
	"sytron-server/storage"
)

func CreateDestination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get json data
		var body models.Destination

		if err := ctx.ShouldBindJSON(&body); err != nil {
			resErr := models.ErrorResponse{
				Message: "There's a problem with your request body",
				Error:   err,
			}
			logger.Handle(err, "Decoding destination json")
			ctx.JSON(http.StatusBadRequest, resErr)
			return
		}

		body.ID = primitive.NewObjectID()

		if res, err := resolvers.DestinationResolver.InsertOne(body); err != nil {
			logger.Handle(err, "Here is where the read fails")
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed reading/writing to database",
				Error:   err,
			})
		} else {
			ctx.JSON(http.StatusOK, res)
		}
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
	data.UpdateTime()
	return resolvers.DestinationResolver.UpdateOne(id, data)
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

		imageUrl, errResponse := uploads_controller.UploadFile(
			ctx,
			"image",
			storage.CMSBucketHandle,
			fileName,
		)
		if errResponse != nil {
			ctx.JSON(http.StatusInternalServerError, errResponse)
			return
		}

		var newDest models.Destination
		newDest.ImageURL = *imageUrl

		if updatedDest, err := updateOneDestination(id, newDest); err != nil {
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
		collection := database.GetCollection(constants.CMS_COLLECTION_DESTINATIONS)

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
