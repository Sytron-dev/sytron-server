package auth_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"sytron-server/constants"
	helpers "sytron-server/helpers"
	"sytron-server/models"
	"sytron-server/resolvers"
)

func GetGenericAuthCredentials(
	role string,
) (CreateAuthCredentials gin.HandlerFunc, Login gin.HandlerFunc) {
	var resolver resolvers.CollectionResolver[models.AuthCredential]

	switch role {
	case constants.USER_ROLE_MERCHANT:
		resolver = resolvers.MerchantAuthCredentialsResolver
	case constants.USER_ROLE_BACKOFFICER:
		resolver = resolvers.BackOfficerAuthCredentialsResolver
	case constants.USER_ROLE_CONSUMER:
		fallthrough
	default:
		resolver = resolvers.UserAuthCredentialsResolver
	}

	// creates user credentials
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
		filter := bson.M{"credential": credentials.Email}
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
			Credential:     credentials.Email,
			CredentialType: "email",
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

	// returns user credentials
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
		filter := bson.M{
			"credential": credentials.Email,
		}

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
		if !helpers.VerifyPassword(credentials.Password, res.Password) {
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
			res.Credential,
			constants.USER_ROLE_MERCHANT,
		)

		// update DB
		if res, err = resolver.UpdateOne(res.ID.String(), *res); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Error updating credentials. You may have to try again or contact support",
				Error:   err,
			})
			return
		}

		ctx.JSON(http.StatusOK, res)
	}

	return
}
