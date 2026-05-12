package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


// Secret key used for JWT verification
// MUST match the key used during token generation
var SECRET = []byte("secret-key")


// AuthMiddleware protects routes
func AuthMiddleware() gin.HandlerFunc {

	// Return middleware function
	return func(c *gin.Context) {

		// Read Authorization header
		//
		// Example:
		// Authorization: Bearer eyJhbGciOi...
		authHeader := c.GetHeader("Authorization")

		// If header missing
		if authHeader == "" {

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header missing",
			})

			// Stop request immediately
			c.Abort()
			return
		}

		// Split header into parts
		//
		// Example:
		// ["Bearer", "TOKEN"]
		parts := strings.Split(authHeader, " ")

		// Validate header format
		if len(parts) != 2 || parts[0] != "Bearer" {

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token format",
			})

			c.Abort()
			return
		}

		// Extract actual JWT token
		tokenString := parts[1]

		// Parse and verify token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			// Return secret key for signature verification
			return SECRET, nil
		})

		// Invalid token
		if err != nil || !token.Valid {

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})

			c.Abort()
			return
		}

		// Extract JWT payload claims
		claims, ok := token.Claims.(jwt.MapClaims)

		if ok {

			// Store user email in request context
			// Future routes can access this
			c.Set("email", claims["email"])
		}

		// Continue to actual route
		c.Next()
	}
}