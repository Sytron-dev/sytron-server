package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"sytron-server/types"
	"sytron-server/util"
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

		claims, err := util.ValidateToken(clientToken)
		if err != "" {
			c.Status(fiber.ErrUnauthorized.Code)
			return c.JSON(bson.M{"error": err, "claims": claims})
		}

		c.Set("email", claims.Email)
		c.Set("uid", claims.ID)
		c.Set("role", claims.Role)

		return c.Next()
	}
}
