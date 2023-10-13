package auth_controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"sytron-server/constants"
	"sytron-server/database"
	helper "sytron-server/helpers"
	"sytron-server/models"
)

var validate = validator.New()

// VerifyPassword checks the input password while verifying it with the password in the DB.
func VerifyPassword(userPassword string, providedPassword string) (check bool, msg string) {
	if err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword)); err != nil {
		msg = "Login or password is incorrect"
		check = false
	} else {
		msg = "Valid credentials"
		check = true
	}
	return
}

// Login is the api used to get a single user
func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := database.GetCollection(constants.USERS_COLLECTION)

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid login credentials"})
			return
		}

		passwordIsValid, msg := VerifyPassword(user.Password, foundUser.Password)
		defer cancel()
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(
			foundUser.Email,
			foundUser.ID.String(),
			constants.USER_ROLE_MERCHANT,
		)

		helper.UpdateAllTokens(token, refreshToken, foundUser.ID.String())

		c.JSON(http.StatusOK, foundUser)
	}
}
