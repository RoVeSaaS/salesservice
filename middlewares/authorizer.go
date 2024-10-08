package middleware

import (
	"net/http"
	"strings"

	"github.com/rovesaas/salesservice/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthenticationMiddleware checks if the user has a valid JWT token
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication token"})
			c.Abort()
			return
		}

		// The token should be prefixed with "Bearer "
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token in Split"})
			c.Abort()
			return
		}

		tokenString = tokenParts[1]

		token, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token", "message": err})
			c.Abort()
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		c.Set("tenant_id", claims["org_id"])
		c.Set("role", claims["role"])
		c.Next()
	}
}
