package middleware

import (
	"sytron-server/helpers"
	"sytron-server/types"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Authz validates token and authorizes users
func Authentication() types.HandlerFunc {
	return func(c *fiber.Ctx) error {
		clientToken := c.Get("Token")
		if clientToken == "" {
			c.Status(fiber.ErrBadRequest.Code)
			return c.JSON(
				bson.M{"error": "No Authorization header provided"},
			)
		}

		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.Status(fiber.ErrUnauthorized.Code)
			return c.JSON( bson.M{"error": err, "claims": claims})
		}

		c.Set("email", claims.Email)
		c.Set("uid", claims.ID)
		c.Set("role", claims.Role)

		return c.Next()
	}
}
