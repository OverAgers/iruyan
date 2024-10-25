package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler構造体
type ErrorHandler struct{}

// BadRequest: 400エラーを返す
func (e *ErrorHandler) BadRequest(c *gin.Context, message string) {
	if message == "" {
		message = "Bad Request"
	}
	data := gin.H{
		"message": message,
		"status":  400,
	}
	c.JSON(http.StatusBadRequest, data)
}

// Unauthorized: 401エラーを返す
func (e *ErrorHandler) Unauthorized(c *gin.Context, message string) {
	if message == "" {
		message = "Unauthorized"
	}
	data := gin.H{
		"message": message,
		"status":  401,
	}
	c.JSON(http.StatusUnauthorized, data)
}

// Forbidden: 403エラーを返す
func (e *ErrorHandler) Forbidden(c *gin.Context, message string) {
	if message == "" {
		message = "Forbidden"
	}
	data := gin.H{
		"message": message,
		"status":  403,
	}
	c.JSON(http.StatusForbidden, data)
}

// NotFoundError: 404エラーを返す
func (e *ErrorHandler) NotFoundError(c *gin.Context, message string) {
	if message == "" {
		message = "Not Found"
	}
	data := gin.H{
		"message": message,
		"status":  404,
	}
	c.JSON(http.StatusNotFound, data)
}

// InternalServerError: 500エラーを返す
func (e *ErrorHandler) InternalServerError(c *gin.Context, message string) {
	if message == "" {
		message = "Internal Server Error"
	}
	data := gin.H{
		"message": message,
		"status":  500,
	}
	c.JSON(http.StatusInternalServerError, data)
}

// MessageError: 汎用的なエラーメッセージを返す
func (e *ErrorHandler) MessageError() map[string]interface{} {
	return map[string]interface{}{
		"code":    "---",
		"message": "Error",
	}
}
