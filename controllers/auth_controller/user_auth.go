package auth_controller

import (
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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

// Login is the api used to get a single user
func LoginUser(c *fiber.Ctx) error {
	collection := database.GetCollection(constants.USERS_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user models.User
	var foundUser models.User

	if err := c.BodyParser(&user); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(types.ErrorResponse{
			Message:  "Something's wrong with the body",
			Error:    err,
			Metadata: err.Error(),
		})
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
		constants.USER_ROLE_MERCHANT,
	)

	helpers.UpdateAllTokens(token, refreshToken, foundUser.ID.String())

	return c.JSON(foundUser)
}
