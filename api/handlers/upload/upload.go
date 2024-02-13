package upload

import (
	"fmt"
	"io"
	"net/http"

	cloud_storage "cloud.google.com/go/storage"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/appengine"

	"sytron-server/types"
)

func UploadFile(
	ctx *fiber.Ctx,
	key string,
	bucketHandle *cloud_storage.BucketHandle,
	filePath string,
) (*string, *types.ErrorResponse) {
	fileHeader, err := ctx.FormFile(key)
	if err != nil {
		return nil, &types.ErrorResponse{
			Message:  "Error reading files in form body",
			Error:    err,
			Metadata: err.Error(),
		}
	}

	f, err := fileHeader.Open()
	if err != nil {
		return nil, &types.ErrorResponse{
			Message:  "Error getting file from form",
			Error:    err,
			Metadata: err.Error(),
		}
	}

	appengineCtx := appengine.NewContext(&http.Request{})

	sw := bucketHandle.Object(filePath).NewWriter(appengineCtx)

	if _, err := io.Copy(sw, f); err != nil {
		return nil, &types.ErrorResponse{
			Message:  "Error copying file to storage",
			Error:    err,
			Metadata: err.Error(),
		}
	}

	if err := sw.Close(); err != nil {
		return nil, &types.ErrorResponse{
			Message:  "Error closing writer",
			Error:    err,
			Metadata: err.Error(),
		}
	}

	// if all is good :
	url := sw.Attrs().MediaLink

	fmt.Println(url)

	return &url, nil
}
