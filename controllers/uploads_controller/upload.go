package uploads_controller

import (
	"io"
	"sytron-server/models"

	cloud_storage "cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
)

func UploadFile(

	ctx *gin.Context,
	key string,
	bucketHandle *cloud_storage.BucketHandle,
	filePath string,

) (*string, *models.ErrorResponse) {

	f, fHeader, err := ctx.Request.FormFile(key)

	if err != nil {
		return nil, &models.ErrorResponse{
			Message: "Error getting file from form",
			Error:   err,
		}
	}
	defer f.Close()

	appengineCtx := appengine.NewContext(ctx.Request)

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
