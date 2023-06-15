package users

import "github.com/gin-gonic/gin"

func RegisterUsersRoutes(router *gin.Engine) {
	CreateUser(router)
}
