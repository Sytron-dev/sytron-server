package cms

import (
	"fmt"
	"net/http"
	"sytron-server/api/handlers/upload"
	"sytron-server/storage/conn"
	"sytron-server/storage/queries"
	"sytron-server/types"
	"sytron-server/types/models"

	"github.com/gofiber/fiber/v2"
)

func CreateDestination() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		// get json data
		var body models.Destination

		if err := ctx.BodyParser(&body); err != nil {
			resErr := types.ErrorResponse{
				Message:  "There's a problem with your request body",
				Error:    err,
				Metadata: err.Error(),
			}
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(resErr)

		}

		if res, err := queries.CreateDestination(body); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed reading/writing to database",
				Error:    err,
				Metadata: err.Error(),
			})
		} else {
			return ctx.JSON(res)
		}
	}
}

func GetDestinations() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		if destinations, err := queries.GetDestinations(); err != nil {
			return ctx.Status(fiber.ErrInternalServerError.Code).JSON(types.ErrorResponse{
				Message:  "Failed reading destinations",
				Error:    err,
				Metadata: err.Error(),
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

		if destination, err := queries.FindOneDestination(id); err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(types.ErrorResponse{
				Message:  "Error finding document",
				Error:    err,
				Metadata: err.Error(),
			})
		} else {
			return ctx.JSON(destination)
		}
	}
}

func UpdateDestination() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		// Get updated destination from request body
		var dest models.Destination

		if err := ctx.BodyParser(&dest); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(types.ErrorResponse{
				Message: "Failed parsing request body",
				Error:   err,
			})
		}

		if updatedDest, err := queries.UpdateDestination(id, dest); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed updating destination",
				Error:    err,
				Metadata: err.Error(),
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

		imageUrl, errResponse := upload.UploadFile(
			ctx,
			"image",
			conn.CMSBucketHandle,
			fileName,
		)
		if errResponse != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(errResponse)
		}

		if updatedDest, err := queries.UpdateDestinationImage(id, *imageUrl); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed updating destination image",
				Error:    err,
				Metadata: err.Error(),
			})
		} else {
			return ctx.JSON(updatedDest)
		}
	}
}

func DeleteDestination() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		if err := queries.DeleteDestination(id); err != nil {
			ctx.Status(http.StatusInternalServerError).JSON(types.ErrorResponse{
				Message: "Failed deleting destination",
				Error:   err,
			})
		}

		return ctx.JSON(types.EmptyResponse{})
	}
}
