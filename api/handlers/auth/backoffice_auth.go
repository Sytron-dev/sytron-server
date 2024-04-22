package auth

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"sytron-server/constants"
	"sytron-server/storage/queries"
	"sytron-server/types"
	"sytron-server/types/models"
	"sytron-server/util"
)

// creates backoffice auth credentials
// var CreateBackOfficeAuth, LoginBackOfficer = GetGenericAuthCredentials(
// 	constants.USER_ROLE_BACKOFFICER,
// 	resolvers.BackOfficerAuthCredentialsResolver,
// )

func CreateBackOfficeAuth() types.HandlerFunc {
	return func(ctx *fiber.Ctx) (err error) {
		const DEFAULT_PASSWORD = "super.secret.shhh!"

		// 1. Init struct values

		var body models.AuthBody

		err = ctx.BodyParser(&body)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed reading request body",
				Metadata: err.Error(),
			})
		}

		credentials := models.AuthCredential{
			Value: body.Email,
			Type:  constants.AUTH_TYPE_EMAIL_AND_PASSWORD,
		}

		// 2. Set the default password
		if pwd, err := util.HashPassword("super.secret.shhh!"); err != nil {
			ctx.Status(fiber.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message: "Failed generating password",

				Metadata: err.Error(),
			})
		} else {
			credentials.Password = pwd
		}

		// 3. Add credentials to the database
		_, err = queries.CreateBackofficerAuthCredentials(credentials)
		if err != nil {
			ctx.Status(fiber.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message: "Failed reading/writing to database",

				Metadata: err.Error(),
			})
		}

		ctx.Status(fiber.StatusOK)

		return
	}
}

func LoginBackOfficer() types.HandlerFunc {
	return func(ctx *fiber.Ctx) (err error) {
		// 1. Init struct values
		var body models.AuthBody
		err = ctx.BodyParser(&body)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed reading request body",
				Metadata: err.Error(),
			})
		}

		// 3. Get credentials

		if credentials, err := queries.GetBackofficerAuthCredential(models.AuthCredential{Value: body.Email}); err != nil {
			ctx.Status(fiber.StatusInternalServerError)
			return ctx.JSON(types.ErrorResponse{
				Message:  "Failed reading/writing to database",
				Metadata: err.Error(),
			})
		} else {

			// 2. Verify credentials

			if !util.VerifyPassword(credentials.Password, body.Password) {
				ctx.Status(fiber.StatusUnauthorized)
				return ctx.JSON(types.ErrorResponse{
					Message: "Invalid credentials",
				})
			}

			// 4. Generate tokens

			credentials.BearerToken, credentials.RefreshToken, err = util.GenerateAllTokens(
				credentials.ID.String(),
				credentials.Value,
				constants.USER_ROLE_BACKOFFICER,
			)
			if err != nil {
				ctx.Status(fiber.StatusInternalServerError)
				return ctx.JSON(types.ErrorResponse{
					Message:  "Failed generating tokens",
					Metadata: err.Error(),
				})
			}

			// return credentials
			return ctx.JSON(bson.M{
				"bearer_token": credentials.BearerToken,
			})
		}
	}
}
