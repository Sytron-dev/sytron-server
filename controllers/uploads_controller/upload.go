package uploads_controller

import (
	"io"
	"net/http"
	"sytron-server/models"

	cloud_storage "cloud.google.com/go/storage"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/appengine"
)

func UploadFile(

	ctx *fiber.Ctx,
	key string,
	bucketHandle *cloud_storage.BucketHandle,
	filePath string,

) (*string, *models.ErrorResponse) {

	fileHeader, _ := ctx.FormFile(key)
	f, err := fileHeader.Open()

	if err != nil {
		return nil, &models.ErrorResponse{
			Message: "Error getting file from form",
			Error:   err,
		}
	}

	appengineCtx := appengine.NewContext(&http.Request{})

	sw := bucketHandle.Object(filePath).NewWriter(appengineCtx)

	if _, err := io.Copy(sw, f); err != nil {
		return nil, &models.ErrorResponse{
			Message: "Error copying file to storage",
			Error:   err,
		}
	}

	if err := sw.Close(); err != nil {
		return nil, &models.ErrorResponse{
			Message: "Error closing writer",
			Error:   err,
		}
	}

	// if all is good :
	url := sw.Attrs().MediaLink
	return &url, nil
}
