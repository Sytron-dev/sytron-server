package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"sytron-server/constants"
	"sytron-server/database"
	"sytron-server/helpers"
	"sytron-server/models"
	"sytron-server/types"
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
func SignUp() types.HandlerFunc {
	return func(c *fiber.Ctx) error {
		collection := database.GetCollection(constants.USERS_COLLECTION)

		// initialise variables for this scope
		var user models.User

		// Create DB context
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Bind user data
		if err := c.BodyParser(&user); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(bson.M{"error": err.Error()})

		}

		// Validate data integrity
		if err := validate.Struct(user); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(bson.M{"error": err.Error()})
		}

		// check if email or phone number exists
		filter := bson.M{
			"$or": bson.A{
				bson.M{"email": user.Email},
				bson.M{"phone": user.Phone},
			},
		}
		if count, err := collection.CountDocuments(ctx, filter); err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(bson.M{"error": "Error occurred while checking for the email"})

		} else if count != 0 {
			c.Status(http.StatusForbidden)
			return c.JSON(bson.M{"error": "This email/phone number already exists. Please contact the support!"})
		}

		// Load info
		var err error
		user.ID = primitive.NewObjectID()
		if user.Password, err = helpers.HashPassword(user.Password); err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(bson.M{"error": "Era"})
		}

		if user.Token, user.RefreshToken, err = helpers.GenerateAllTokens(
			user.Email,
			user.ID.String(),
			constants.USER_ROLE_CONSUMER,
		); err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(bson.M{"error": "Era"})
		}

		user.InsertTime()

		resultInsertionNumber, err := collection.InsertOne(ctx, user)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(bson.M{"error": "User item was not created"})

		}
		return c.JSON(resultInsertionNumber)
	}
}

// Login is the api used to get a single user
func Login() types.HandlerFunc {
	return func(c *fiber.Ctx) error {
		collection := database.GetCollection(constants.USERS_COLLECTION)

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		var foundUser models.User

		if err := c.BodyParser(&user); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(bson.M{"error": err.Error()})

		}

		err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(bson.M{"error": "Invalid login credentials"})

		}

		passwordIsValid, msg := VerifyPassword(user.Password, foundUser.Password)
		defer cancel()
		if !passwordIsValid {
			c.Status(http.StatusInternalServerError)
			return c.JSON(bson.M{"error": msg})

		}

		token, refreshToken, _ := helpers.GenerateAllTokens(
			foundUser.Email,
			foundUser.ID.String(),
			constants.USER_ROLE_CONSUMER,
		)

		helpers.UpdateAllTokens(token, refreshToken, foundUser.ID.String())

		return c.JSON(foundUser)
	}
}
