package confs

import (
	"fmt"
	"net/http"
	"sytron-server/database"
	"sytron-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetCountries(router *gin.Engine) {
	router.GET("/confs/countries", func(ctx *gin.Context) {

		var results map[string]interface{}
		err := database.GetClient().DB.From("users").Select("*").Execute(&results)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, results)

	})
}

func CreateCountry(router *gin.Engine) {
	router.POST("/confs/countries", func(ctx *gin.Context) {

		var newCountry models.Country
		if err := ctx.BindJSON(&newCountry); err != nil {
			return
		}
		newCountry.ID = uuid.New()

		var results map[string]interface{}
		err := database.GetClient().DB.From("countries").Insert(newCountry).Execute(&results)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err,
				"req": ctx.Request,
			})
			return
		}
		fmt.Println(results)

		ctx.JSON(http.StatusOK, results)

	})

}
