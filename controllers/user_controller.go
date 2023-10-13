package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// CreateUser is the api used to get a single user
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := database.GetCollection(constants.USERS_COLLECTION)

		// initialise variables for this scope
		var user models.User

		// Create DB context
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Bind user data
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate data integrity
		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// check if email or phone number exists
		filter := bson.M{
			"$or": bson.A{
				bson.M{"email": user.Email},
				bson.M{"phone": user.Phone},
			},
		}
		if count, err := collection.CountDocuments(ctx, filter); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "Error occurred while checking for the email"},
			)
			return
		} else if count != 0 {
			c.JSON(
				http.StatusForbidden,
				gin.H{"error": "This email/phone number already exists. Please contact the support!"},
			)
			return
		}

		// Load info
		var err error
		user.ID = primitive.NewObjectID()
		if user.Password, err = helper.HashPassword(user.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Era"})
		}
		if user.Token, user.RefreshToken, err = helper.GenerateAllTokens(
			user.Email,
			user.ID.String(),
			constants.USER_ROLE_CONSUMER,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Era"})
			return
		}

		user.InsertTime()

		resultInsertionNumber, err := collection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User item was not created"})
			return
		}
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

// Login is the api used to get a single user
func Login() gin.HandlerFunc {
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
			constants.USER_ROLE_CONSUMER,
		)

		helper.UpdateAllTokens(token, refreshToken, foundUser.ID.String())

		c.JSON(http.StatusOK, foundUser)
	}
}
