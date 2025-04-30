package infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secretKey string) gin.HandlerFunc{
	return func(c *gin.Context){
		// Check Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer"){
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := ValidateToken(token)
		if err != nil{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}

		c.Set("userID", claims["userID"])
		c.Set("role", claims["role"])
		c.Next()
	}
}