package cms_controller

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"sytron-server/controllers/uploads_controller"
	"sytron-server/database"
	"sytron-server/helpers/logger"
	"sytron-server/models"
	"sytron-server/resolvers"
	"sytron-server/storage"
	"sytron-server/types"
)

func CreateDestination() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		// get json data
		var body models.Destination

		if err := ctx.BodyParser(&body); err != nil {
			resErr := models.ErrorResponse{
				Message: "There's a problem with your request body",
				Error:   err,
			}
			logger.Handle(err, "Decoding destination json")
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(resErr)

		}

		body.ID = primitive.NewObjectID()

		if res, err := resolvers.DestinationResolver.InsertOne(body); err != nil {
			logger.Handle(err, "Here is where the read fails")
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

func GetDestinations() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		if destinations, err := resolvers.DestinationResolver.FindMany(database.PaginationOptions{}); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed while reading database",
				Error:   err,
			})
		} else {
			return ctx.JSON(destinations)
		}
	}
}

func GetSingleDestination() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		// Get _id param
		id := ctx.Params("id")

		if destination, err := resolvers.DestinationResolver.FindOneByID(id); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Error finding document",
				Error:   err,
			})
		} else {
			return ctx.JSON(destination)
		}
	}
}

func updateOneDestination(id string, data models.Destination) (*models.Destination, error) {
	data.UpdateTime()
	return resolvers.DestinationResolver.UpdateOne(id, data)
}

func UpdateDestination() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		// Get updated destination from request body
		var dest models.Destination

		if err := ctx.BodyParser(&dest); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed parsing request body",
				Error:   err,
			})
		}

		if updatedDest, err := updateOneDestination(id, dest); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed updating destination",
				Error:   err,
			})
		} else {
			return ctx.JSON(updatedDest)
		}
	}
}

func UploadDestinationImage() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		fileName := fmt.Sprintf("destinations/%v/image", id)

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

		var newDest models.Destination
		newDest.ImageURL = *imageUrl

		if updatedDest, err := updateOneDestination(id, newDest); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed updating destination",
				Error:   err,
			})
		} else {
			return ctx.JSON(updatedDest)
		}
	}
}

func DeleteDestination() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		if err := resolvers.DestinationResolver.DeleteOne(id); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(models.ErrorResponse{
				Message: "Failed deleting destination",
				Error:   err,
			})
		}

		return ctx.JSON(types.EmptyResponse{})
	}
}
