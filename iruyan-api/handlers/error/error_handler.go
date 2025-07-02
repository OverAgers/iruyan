// Package errorhandler provides centralized error response utilities
// for handling common HTTP errors such as 400, 401, 403, 404, and 500.
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
		case http.StatusConflict:
			message = "StatusConflict"
		case http.StatusInternalServerError:
			message = "Internal Server Error"
		}
	}
	c.JSON(statusCode, gin.H{
		"message": message,
		"status":  statusCode,
	})
}

// BadRequest handles HTTP 400 Bad Request errors.
func (e *ErrorHandler) BadRequest(c *gin.Context, message string) {
	e.ErrorResponse(c, http.StatusBadRequest, message)
}

// Unauthorized handles HTTP 401 Unauthorized errors.
func (e *ErrorHandler) Unauthorized(c *gin.Context, message string) {
	e.ErrorResponse(c, http.StatusUnauthorized, message)
}

// Forbidden handles HTTP 403 Forbidden errors.
func (e *ErrorHandler) Forbidden(c *gin.Context, message string) {
	e.ErrorResponse(c, http.StatusForbidden, message)
}

// NotFoundError handles HTTP 404 Not Found errors.
func (e *ErrorHandler) NotFoundError(c *gin.Context, message string) {
	e.ErrorResponse(c, http.StatusNotFound, message)
}

// Conflict handles HTTP 409 Conflict errors.
func (e *ErrorHandler) Conflict(c *gin.Context, message string) {
	e.ErrorResponse(c, http.StatusConflict, message)
}

// UnprocessableEntity handles HTTP 422 Unprocessable Entity.
func (e *ErrorHandler) UnprocessableEntity(c *gin.Context, message string) {
	e.ErrorResponse(c, http.StatusUnprocessableEntity, message)
}

// InternalServerError handles HTTP 500 Internal Server Error.
func (e *ErrorHandler) InternalServerError(c *gin.Context, message string) {
	e.ErrorResponse(c, http.StatusInternalServerError, message)
}
