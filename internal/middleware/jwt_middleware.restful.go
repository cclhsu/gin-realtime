package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Middleware to verify the JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Extract the token from the header
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		tokenString := tokenParts[1]

		jwtSecret := []byte(os.Getenv("JWT_SECRET"))

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		// Verify the token
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Check if the token is valid
		if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
			// Set the username in the request context
			c.Set("username", claims.ID)
			c.Set("UUID", claims.Sub)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
	}
}
