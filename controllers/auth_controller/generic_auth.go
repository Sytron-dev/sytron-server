package auth_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"sytron-server/helpers"
	"sytron-server/models"
	"sytron-server/resolvers"
)

func GetGenericAuthCredentials(
	role string,
	resolver resolvers.CollectionResolver[models.AuthCredential],
) (CreateAuthCredentials gin.HandlerFunc, Login gin.HandlerFunc) {

	// creates user credentials ----------------------------------------------------
	CreateAuthCredentials = func(ctx *gin.Context) {
		// get data from request body
		type Credentials struct {
			Email string `json:"email"`
		}

		var credentials Credentials

		if err := ctx.BindJSON(&credentials); err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
				Message: "Failed to parse body",
				Error:   err,
			})
			return
		}

		// ensure uniqueness
		filter := bson.D{{Key: "value", Value: credentials.Email}}
		if count, err := resolver.CountDocuments(filter); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed checking credential validity",
				Error:   err,
			})
			return
		} else if count > 0 {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
				Message: "This user exists",
				Error:   nil,
			})
			return
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
		if pwd, err := helpers.HashPassword("super.secret.shhh!"); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed generating password",
				Error:   err,
			})
			return
		} else {
			authCredentials.Password = pwd
		}

		// add credentials to database
		res, err := resolver.InsertOne(authCredentials)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed creating backofficer",
				Error:   err,
			})
			return
		}
		// return credentials
		ctx.JSON(http.StatusOK, res)
	}

	// returns user credentials ------------------------------------------------------
	Login = func(ctx *gin.Context) {
		// get data from request body

		type LoginCredentials struct {
			Email    string `json:"email"    validate:"required"`
			Password string `json:"password" validate:"required"`
		}

		var credentials LoginCredentials
		if err := ctx.BindJSON(&credentials); err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
				Message: "Failed to parse body",
				Error:   err,
			})
			return
		}

		// validate Password
		filter := bson.D{{
			Key:   "value",
			Value: credentials.Email,
		}}

		// check if credentials are correct
		res, err := resolver.FindOne(filter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Invalid credentials",
				Error:   err,
			})
			return
		}

		// validate password
		if !helpers.VerifyPassword(res.Password, credentials.Password) {
			ctx.JSON(http.StatusForbidden, models.ErrorResponse{
				Message: "Invalid credentials fr",
				Error:   err,
			})
			return
		}
		// update credentials
		res.UpdateLastLogin()
		res.BearerToken, res.RefreshToken, err = helpers.GenerateAllTokens(
			res.ID.String(),
			res.Value,
			role,
		)

		// update DB
		if res, err = resolver.UpdateOne(res.ID.Hex(), *res); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed updating record on db",
				Error:   err,
			})
			return
		}
		ctx.JSON(http.StatusOK, bson.M{
			"bearer_token":  res.BearerToken,
			"refresh_token": res.RefreshToken,
		})
	}

	return
}
