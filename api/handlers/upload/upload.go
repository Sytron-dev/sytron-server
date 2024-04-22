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
			Metadata: err.Error(),
		}
	}

	f, err := fileHeader.Open()
	if err != nil {
		return nil, &types.ErrorResponse{
			Message:  "Error getting file from form",
			Metadata: err.Error(),
		}
	}

	appengineCtx := appengine.NewContext(&http.Request{})

	sw := bucketHandle.Object(filePath).NewWriter(appengineCtx)

	if _, err := io.Copy(sw, f); err != nil {
		return nil, &types.ErrorResponse{
			Message:  "Error copying file to storage",
			Metadata: err.Error(),
		}
	}

	if err := sw.Close(); err != nil {
		return nil, &types.ErrorResponse{
			Message:  "Error closing writer",
			Metadata: err.Error(),
		}
	}

	// if all is good :
	url := sw.Attrs().MediaLink

	fmt.Println(url)

	return &url, nil
}

func DeleteFile(
	ctx *fiber.Ctx,
	bucketHandle *cloud_storage.BucketHandle,
	filePath string,
) (err types.ErrorResponse) {
	if delErr := bucketHandle.Object(filePath).Delete(ctx.Context()); delErr != nil {
		err.Message = "Failed to delete file"
		err.Metadata = delErr.Error()
	}

	return err
}

func DeleteFolder(
	ctx *fiber.Ctx,
	bucketHandle *cloud_storage.BucketHandle,
	folderPath string,
) (errRes types.ErrorResponse) {
	it := bucketHandle.Objects(ctx.Context(), &cloud_storage.Query{Prefix: folderPath})

	errs := []error{}

	// delete all individual object
	for {
		obj, err := it.Next()
		if err != nil {
			break
		}

		if err = bucketHandle.Object(obj.Name).Delete(ctx.Context()); err != nil {
			errs = append(errs, err)
		}
	}

	// collect errors
	if len(errs) > 0 {
		errRes.Message = "Failed to delete a few files. The object may contain useful metadata to debug this"
		errRes.Metadata = fmt.Sprintln(errs)
	}

	return
}
