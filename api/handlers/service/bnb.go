package service_controller

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"sytron-server/api/handlers/upload"
	"sytron-server/storage/conn"
	"sytron-server/storage/queries"
	"sytron-server/types"
	"sytron-server/types/models"
)

// CRUD for bed & breakfast services

func CreateBNB() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		// get json data
		var body models.BNB

		if err := ctx.BodyParser(&body); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(types.ErrorResponse{
				Message: "There's a problem with your request body",
				Error:   err,
			})
		}

		// add to database
		if res, err := queries.CreateBNB(body); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed writing to database",
				Error:    err,
				Metadata: err.Error(),
			})
		} else {
			return ctx.JSON(res)
		}
	}
}

func GetBNBs() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		if bnb, err := queries.GetBNBs(); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed while reading database",
				Error:    err,
				Metadata: err.Error(),
			})
		} else {
			return ctx.JSON(bnb)
		}
	}
}

func GetSingleBNB() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		// Get _id param
		id := ctx.Params("id")

		if bnb, err := queries.FindOneBNB(id); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed while reading database",
				Error:    err,
				Metadata: err.Error(),
			})
		} else {
			return ctx.JSON(bnb)
		}
	}
}

func UpdateBNB() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		// Get _id Params
		id := ctx.Params("id")

		// get json data
		var data models.BNB
		if err := ctx.BodyParser(&data); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed to parse JSON data",
				Error:    err,
				Metadata: err.Error(),
			})
		}

		data.UpdateTime()

		// update data
		if res, err := queries.UpdateBNB(id, data); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed while reading database",
				Error:    err,
				Metadata: err.Error(),
			})
		} else {
			return ctx.JSON(res)
		}
	}
}

// Asset management
func UploadBNBImage() types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		fileName := fmt.Sprintf("services/bnb/%v/image", id)

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

		if updatedBNB, err := queries.UpdateBnbImage(id, *imageUrl); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed updating bnb image",
				Error:    err,
				Metadata: err.Error(),
			})
		} else {
			return ctx.JSON(updatedBNB)
		}
	}
}

// Danger zone

func DeleteBNB() types.HandlerFunc {
	return func(ctx *fiber.Ctx) (err error) {
		id := ctx.Params("id")

		if err = queries.DeleteBNB(id); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.Status(fiber.StatusOK)
		return
	}
}
