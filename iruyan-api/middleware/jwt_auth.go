package middleware

import (
	"net/http"
	"strings"

	"iruyan-api/pkg/errdefs"
	"iruyan-api/utils"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		_, claims, err := utils.ParseJWT(tokenString)
		switch err {
		case nil:
			// OK
		case errdefs.ErrJWTSecretNotSet:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Server misconfiguration: JWT secret not set"})
			return
		case errdefs.ErrJWTInvalidToken:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		case errdefs.ErrJWTParseFailure:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token parsing failed"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// JWTが有効な場合は claims を context にセット
		c.Set("user", claims)
		c.Next()
	}
}
