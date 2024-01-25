package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	"sytron-server/constants"
)

func InitJWTAuth() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key:    []byte(constants.SECRET_KEY),
			JWTAlg: jwtware.HS256,
		},
	})
}
