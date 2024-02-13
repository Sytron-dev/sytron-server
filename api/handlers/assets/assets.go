package assets

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"sytron-server/api/handlers/upload"
	"sytron-server/storage/conn"
	"sytron-server/storage/queries"
	"sytron-server/types"
	"sytron-server/types/models"
)

/**
The following function are generics for asset management
each can be used to upload assets for different tables and therefore require a reference (table name)

_type ==> table name
_reference ==> row _id

*/

// type is the table name
func CreateAsset(tableRef string, rootDir string) types.HandlerFunc {
	return func(ctx *fiber.Ctx) error {
		// get _ref from ctx id
		refID := ctx.Params("id")

		asset := models.Asset{
			SqlDocument: models.SqlDocument{
				ID: uuid.New(),
			},
			Alt: ctx.FormValue("alt"),
			// FIXME implement better mimetypes
			Format:    ctx.FormValue("format"),
			Reference: refID,
			Type:      tableRef,
		}

		if asset.Url == "" {
			fileName := fmt.Sprintf("%s/%s/assets/%s", rootDir, refID, asset.ID.String())

			imageUrl, errResponse := upload.UploadFile(
				ctx,
				"file",
				conn.CMSBucketHandle,
				fileName,
			)

			if errResponse != nil {
				ctx.Status(http.StatusInternalServerError)
				return ctx.JSON(errResponse)
			}

			asset.Url = *imageUrl
		}

		if res, err := queries.CreateAsset(asset); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed saving asset",
				Error:    err,
				Metadata: err.Error(),
			})
		} else {
			return ctx.JSON(res)
		}
	}
}

// Update asset

func UpdateAsset(tableRef string, rootDir string) types.HandlerFunc {
	return func(ctx *fiber.Ctx) (err error) {
		// get asset id and ref id
		refID := ctx.Params("ref_id")
		assetID := ctx.Params("id")

		asset := models.Asset{
			Alt:    ctx.FormValue("alt"),
			Format: ctx.FormValue("format"),
			Type:   ctx.FormValue("type"),
			Url:    ctx.FormValue("url"),
		}

		asset.ID, _ = uuid.Parse(assetID)

		if asset.Url == "" {
			// Priority goes to url
			fileName := fmt.Sprintf("%s/%s/assets/%s", rootDir, refID, asset.ID.String())

			imageUrl, errResponse := upload.UploadFile(
				ctx,
				"file",
				conn.CMSBucketHandle,
				fileName,
			)

			if errResponse != nil {
				ctx.Status(http.StatusInternalServerError)
				return ctx.JSON(errResponse)
			}

			asset.Url = *imageUrl
		}

		if res, err := queries.UpdateAsset(assetID, asset); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed updating asset",
				Error:    err,
				Metadata: err.Error(),
			})
		} else {
			return ctx.JSON(res)
		}
	}
}
