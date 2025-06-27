// Package middleware provides Gin middleware functions for error recovery handling.
package middleware

import (
	"fmt"
	errorhandler "iruyan-api/handlers/error"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware - パニック発生時のリカバリとエラーレスポンス
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				// パニック内容を文字列に変換
				errMsg := fmt.Sprintf("Internal Server Error: %v", rec)
				// スタックトレースの取得
				stackTrace := string(debug.Stack())

				// ハンドラーでエラーレスポンスを作成
				handler := errorhandler.ErrorHandler{}
				handler.InternalServerError(c, fmt.Sprintf("%s\n%s", errMsg, stackTrace))

				// リクエストの処理を中断
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Internal Server Error",
					"error":   errMsg,
					"stack":   stackTrace, // スタックトレースを含める
				})
			}
		}()
		c.Next()
	}
}
