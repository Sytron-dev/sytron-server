package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"sytron-server/helpers"
)

// Authz validates token and authorizes users
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": fmt.Sprintf("No Authorization header provided")},
			)
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("uid", claims.ID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
