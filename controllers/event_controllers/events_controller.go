package event_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"sytron-server/models"
	"sytron-server/resolvers"
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
