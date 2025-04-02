package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tird4d/user-api/utils"
)

func JWTAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeaser := c.GetHeader("Authorization")
		if authHeaser == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missed"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeaser, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		token := parts[1]

		claims, err := utils.ValidateJWT(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Next()

	}
}
