package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"sytron-server/storage/resolvers"
	"sytron-server/types"
	"sytron-server/types/models"
	"sytron-server/util"
)

func GetGenericAuthCredentials(
	role string,
	resolver resolvers.CollectionResolver[models.AuthCredential],
) (CreateAuthCredentials types.HandlerFunc, Login types.HandlerFunc) {
	// creates user credentials ----------------------------------------------------
	CreateAuthCredentials = func(ctx *fiber.Ctx) error {
		// get data from request body
		type Credentials struct {
			Email string `json:"email"`
		}

		var credentials Credentials

		if err := ctx.BodyParser(&credentials); err != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(types.ErrorResponse{
				Message: "Failed to parse body",
				Error:   err,
			})
		}

		// ensure uniqueness
		filter := bson.D{{Key: "value", Value: credentials.Email}}
		if count, err := resolver.CountDocuments(filter); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message: "Failed checking credential validity",
				Error:   err,
			})
		} else if count > 0 {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message: "This user exists",
				Error:   nil,
			})
		}

		// initialize defaults
		authCredentials := models.AuthCredential{
			Value: credentials.Email,
			Type:  "email",
			CollectionDocument: models.CollectionDocument{
				ID: primitive.NewObjectID(),
			},
		}
		authCredentials.InsertTime()

		// set default password
		if pwd, err := util.HashPassword("super.secret.shhh!"); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message: "Failed generating password",
				Error:   err,
			})
		} else {
			authCredentials.Password = pwd
		}

		// add credentials to database
		res, err := resolver.InsertOne(authCredentials)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message: "Failed creating backofficer",
				Error:   err,
			})

		}
		// return credentials
		return ctx.JSON(res)
	}

	// returns user credentials ------------------------------------------------------
	Login = func(ctx *fiber.Ctx) error {
		// get data from request body

		type LoginCredentials struct {
			Email    string `json:"email"    validate:"required"`
			Password string `json:"password" validate:"required"`
		}

		var credentials LoginCredentials
		if err := ctx.BodyParser(&credentials); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(types.ErrorResponse{
				Message: "Failed to parse body",
				Error:   err,
			})

		}

		// validate Password
		filter := bson.D{{
			Key:   "value",
			Value: credentials.Email,
		}}

		// check if credentials are correct
		res, err := resolver.FindOne(filter)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message: "Invalid credentials",
				Error:   err,
			})

		}

		// validate password
		if !util.VerifyPassword(res.Password, credentials.Password) {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(types.ErrorResponse{
				Message: "Invalid credentials fr",
				Error:   err,
			})
		}
		// update credentials
		res.UpdateLastLogin()
		res.BearerToken, res.RefreshToken, err = util.GenerateAllTokens(
			res.ID.String(),
			res.Value,
			role,
		)

		// update DB
		if res, err = resolver.UpdateOne(res.ID.Hex(), *res); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message: "Failed updating record on db",
				Error:   err,
			})

		}
		return ctx.JSON(bson.M{
			"bearer_token":  res.BearerToken,
			"refresh_token": res.RefreshToken,
		})
	}

	return
}
