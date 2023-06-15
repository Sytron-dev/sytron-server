package users	

import (
	"net/http"
	"sytron-server/database"
	"sytron-server/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUser(router *gin.Engine) {
	router.POST("users/users", func(ctx *gin.Context) {

		var newUser models.User
		if err := ctx.BindJSON(&newUser); err != nil {
			return
		}
		newUser.ID = uuid.New()

		var results map[string]interface{}
		err := database.GetClient().DB.From("users").Insert(newUser).Execute(&results)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err,
				"req": ctx.Request,
			})
			return
		}
	})

}
func UpdateUser(router *gin.Engine) {
	
}
func DeleteUser(router *gin.Engine) {
	
}