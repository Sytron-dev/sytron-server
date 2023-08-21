package uploads_controller

import (
	"io"
	"net/http"
	"net/url"
	"sytron-server/models"
	"sytron-server/storage"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
)

func TestUploads() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		f, uploadFile, err := ctx.Request.FormFile("file")

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Error getting file from form",
				Error:   err,
			})
			return
		}

		defer f.Close()

		appengineCtx := appengine.NewContext(ctx.Request)

		sw := storage.CMSBucket.Object(uploadFile.Filename).NewWriter(appengineCtx)

		if _, err := io.Copy(sw, f); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Error copying file to storage",
				Error:   err,
			})

			return
		}

		if err := sw.Close(); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Error closing writer",
				Error:   err,
			})
		}

		u, err := url.Parse("/" + "stride-gcp" + "/" + sw.Attrs().Name)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"Error":   true,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":  "file uploaded successfully",
			"pathname": u.EscapedPath(),
		})

	}
}
