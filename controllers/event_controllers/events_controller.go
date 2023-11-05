package event_controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"sytron-server/controllers/uploads_controller"
	"sytron-server/database"
	"sytron-server/models"
	"sytron-server/resolvers"
	"sytron-server/storage"
	"sytron-server/types"
)

func CreateEvent() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		// get json data
		var body models.Event

		if err := ctx.BodyParser(&body); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "There's a problem with your request body",
				Error:   err,
			})
		}

		// init default info
		body.ID = primitive.NewObjectID()
		body.InsertTime()

		if res, err := resolvers.EventsResolver.InsertOne(body); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed reading/writing to database",
				Error:   err,
			})
		} else {
			return ctx.JSON(res)
		}
	}
}

func GetSingleEvent() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		// Get _id param
		id := ctx.Params("id")

		// variables
		// var event models.Event
		// var company models.Company
		// var location models.Location

		if event, err := resolvers.EventsResolver.FindOneByID(id); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Error finding event",
				Error:   err,
			})
		} else {
			return ctx.JSON(event)
		}
	}
}

func GetEvents() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		if events, err := resolvers.EventsResolver.FindMany(database.PaginationOptions{}); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed while reading events",
				Error:   err,
			})
		} else {
			return ctx.JSON(events)
		}
	}
}

func updateOneEvent(id string, data models.Event) (*models.Event, error) {
	data.UpdateTime()
	return resolvers.EventsResolver.UpdateOne(id, data)
}

func UpdateEvent() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		// bind json data
		var body models.Event

		if err := ctx.BodyParser(body); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed reading request body",
				Error:   err,
			})
		}

		body.UpdateTime()
		if res, err := resolvers.EventsResolver.UpdateOne(id, body); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed updating event",
				Error:   err,
			})
		} else {
			return ctx.JSON(res)
		}
	}
}

func UploadEventHeroImage() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		fileName := fmt.Sprintf("events/%v/image", id)

		imageUrl, errResponse := uploads_controller.UploadFile(
			ctx,
			"image",
			storage.CMSBucketHandle,
			fileName,
		)
		if errResponse != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(errResponse)
		}

		var newEvent models.Event
		newEvent.HeroImageUrl = *imageUrl

		if updatedEvent, err := updateOneEvent(id, newEvent); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed updating event",
				Error:   err,
			})
		} else {
			return ctx.JSON(updatedEvent)
		}
	}
}

func DeleteEvent() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		if err := resolvers.EventsResolver.DeleteOne(id); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed deleting event",
				Error:   err,
			})
		}
		ctx.Status(http.StatusNoContent)
		return nil
	}
}
