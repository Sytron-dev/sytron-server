package uploads_controller

import (
	"net/http"
	"sytron-server/storage"

	"github.com/gin-gonic/gin"
)

func TestUploads() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if u, errResponse := uploadFile(ctx, "file", storage.CMSBucketHandle, "test/test.png"); errResponse != nil {
			ctx.JSON(http.StatusInternalServerError, errResponse)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "File uploaded successfully",
				"data":    u,
			})
		}
	}
}
