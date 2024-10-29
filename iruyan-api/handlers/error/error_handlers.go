package errorhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler 構造体
type ErrorHandler struct{}

// ErrorResponse エラーレスポンスを統一する関数
func (e *ErrorHandler) ErrorResponse(c *gin.Context, statusCode int, message string) {
	if message == "" {
		switch statusCode {
		case http.StatusBadRequest:
			message = "Bad Request"
		case http.StatusUnauthorized:
			message = "Unauthorized"
		case http.StatusForbidden:
			message = "Forbidden"
		case http.StatusNotFound:
			message = "Not Found"
		case http.StatusInternalServerError:
			message = "Internal Server Error"
		}
	}
	c.JSON(statusCode, gin.H{
		"message": message,
		"status":  statusCode,
	})
}

// メソッド: BadRequest (400)
func (e *ErrorHandler) BadRequest(c *gin.Context, message string) {
	e.ErrorResponse(c, http.StatusBadRequest, message)
}

// メソッド: Unauthorized (401)
func (e *ErrorHandler) Unauthorized(c *gin.Context, message string) {
	e.ErrorResponse(c, http.StatusUnauthorized, message)
}

// メソッド: Forbidden (403)
func (e *ErrorHandler) Forbidden(c *gin.Context, message string) {
	e.ErrorResponse(c, http.StatusForbidden, message)
}

// メソッド: NotFoundError (404)
func (e *ErrorHandler) NotFoundError(c *gin.Context, message string) {
	e.ErrorResponse(c, http.StatusNotFound, message)
}

// メソッド: InternalServerError (500)
func (e *ErrorHandler) InternalServerError(c *gin.Context, message string) {
	e.ErrorResponse(c, http.StatusInternalServerError, message)
}
