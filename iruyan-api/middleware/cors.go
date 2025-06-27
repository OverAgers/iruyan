// Package middleware provides Gin middleware functions for CORS handling.
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware sets CORS headers for allowing requests from frontend.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000") // 特定のオリジンを許可
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
