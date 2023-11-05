package event_controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"sytron-server/controllers/uploads_controller"
	"sytron-server/database"
	"sytron-server/models"
	"sytron-server/resolvers"
	"sytron-server/storage"
)

func CreateEvent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get json data
		var body models.Event

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "There's a problem with your request body",
				Error:   err,
			})
			return
		}

		// init default info
		body.ID = primitive.NewObjectID()
		body.InsertTime()

		if res, err := resolvers.EventsResolver.InsertOne(body); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed reading/writing to database",
				Error:   err,
			})
		} else {
			ctx.JSON(http.StatusOK, res)
		}
	}
}

func GetSingleEvent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get _id param
		id := ctx.Params.ByName("id")

		// variables
		// var event models.Event
		// var company models.Company
		// var location models.Location

		if event, err := resolvers.EventsResolver.FindOneByID(id); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Error finding event",
				Error:   err,
			})
			return
		} else {
			ctx.JSON(http.StatusOK, event)
		}
	}
}

func GetEvents() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if events, err := resolvers.EventsResolver.FindMany(database.PaginationOptions{}); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed while reading events",
				Error:   err,
			})
		} else {
			ctx.JSON(http.StatusOK, events)
		}
	}
}

func updateOneEvent(id string, data models.Event) (*models.Event, error) {
	data.UpdateTime()
	return resolvers.EventsResolver.UpdateOne(id, data)
}

func UpdateEvent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")

		// bind json data
		var body models.Event

		if err := ctx.ShouldBindJSON(body); err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
				Message: "Failed reading request body",
				Error:   err,
			})
			return
		}

		body.UpdateTime()
		if res, err := resolvers.EventsResolver.UpdateOne(id, body); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed updating event",
				Error:   err,
			})
		} else {
			ctx.JSON(http.StatusOK, res)
		}
	}
}

func UploadEventHeroImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")

		fileName := fmt.Sprintf("events/%v/image", id)

		imageUrl, errResponse := uploads_controller.UploadFile(
			ctx,
			"image",
			storage.CMSBucketHandle,
			fileName,
		)
		if errResponse != nil {
			ctx.JSON(http.StatusInternalServerError, errResponse)
		}

		var newEvent models.Event
		newEvent.HeroImageUrl = *imageUrl

		if updatedEvent, err := updateOneEvent(id, newEvent); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed updating event",
				Error:   err,
			})
		} else {
			ctx.JSON(http.StatusOK, updatedEvent)
		}
	}
}

func DeleteEvent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")

		if err := resolvers.EventsResolver.DeleteOne(id); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed deleting event",
				Error:   err,
			})
			return
		}
		ctx.Status(http.StatusNoContent)
	}
}
