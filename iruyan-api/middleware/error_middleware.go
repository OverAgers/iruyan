package middleware

import (
	errorhandler "iruyan-api/handlers/error" // エラーハンドラーをインポート

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware パニックをキャッチして500エラーレスポンスを返すミドルウェア
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				handler := errorhandler.ErrorHandler{}
				handler.InternalServerError(c, "Internal Server Error")
				c.Abort()
			}
		}()
		c.Next()
	}
}
