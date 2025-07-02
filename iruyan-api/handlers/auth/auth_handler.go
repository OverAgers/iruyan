// Package auth provides HTTP handlers for user authentication,
// including login, registration, and logout functionality.
package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	errorhandler "iruyan-api/handlers/error"
	"iruyan-api/pkg/errdefs"

	"iruyan-api/responses"
	usecase "iruyan-api/usecases/auth"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthUsecase usecase.AuthUsecase
}

// LoginPageHandler ログイン画面表示
// @Summary Show login page
// @Description Displays the login page with a message
// @Tags auth
// @Produce json
// @Success 200 {object} responses.ErrorResponse
// @Router /login [get]
func LoginPageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, responses.ErrorResponse{
		Message: "Login page accessed successfully",
	})
}

// LoginHandler ログイン処理
// @Summary User login
// @Description Authenticates the user based on iruyanID and password
// @Tags auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param iruyanId formData string true "Iruyan ID" default(johndoe)
// @Param password formData string true "Password" default(pass1234)
// @Success 200 {object} responses.LoginSuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	iruyanID := c.PostForm("iruyanId")
	password := c.PostForm("password")

	if iruyanID == "" || password == "" {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "iruyanId and password are required",
		})
		return
	}

	var authUsecase usecase.AuthUsecase
	user, err := authUsecase.LoginUser(iruyanID, password)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Message: "user not found"})
		case errors.Is(err, errdefs.ErrInvalidPassword):
			c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Message: "invalid password"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
	}

	c.JSON(http.StatusOK, responses.LoginSuccessResponse{
		Message: "Login successful",
		User: responses.UserInfo{
			IruyanID: user.IruyanID,
			Name:     user.Name,
			Email:    user.Email,
		},
	})
}

// RegisterPageHandler 新規登録画面表示
// @Summary Show registration page
// @Description Displays the registration page with a message
// @Tags auth
// @Produce json
// @Success 200 {object} responses.ErrorResponse
// @Router /register [get]
func RegisterPageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, responses.ErrorResponse{
		Message: "Register page accessed successfully",
	})
}

// RegisterHandler 新規登録処理
// @Summary User registration
// @Description Registers a new user with the provided details
// @Tags auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param iruyanId formData string true "Iruyan ID" default(johndoe)
// @Param password formData string true "Password" default(pass1234)
// @Param userName formData string true "Name" default(John Doe)
// @Param email formData string true "Email" default(johndoe@example.com)
// @Success 200 {object} responses.RegisterSuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /register [post]
func RegisterHandler(c *gin.Context) {
	iruyanID := c.PostForm("iruyanId")
	password := c.PostForm("password")
	name := c.PostForm("userName")
	email := c.PostForm("email")

	errorHandler := errorhandler.ErrorHandler{}

	// --- パラメータのバリデーション ---
	missing := []string{}
	if iruyanID == "" {
		missing = append(missing, "iruyanId")
	}
	if password == "" {
		missing = append(missing, "password")
	}
	if name == "" {
		missing = append(missing, "userName")
	}
	if email == "" {
		missing = append(missing, "email")
	}
	if len(missing) > 0 {
		errorHandler.BadRequest(c, "Missing required parameter(s): "+strings.Join(missing, ", "))
		return
	}

	var authUsecase usecase.AuthUsecase
	user, err := authUsecase.RegisterUser(name, iruyanID, password, email)
	if err != nil {
		msg := err.Error()
		switch {
		case errors.Is(err, errdefs.ErrDuplicateIruyanID),
			errors.Is(err, errdefs.ErrDuplicateEmail):
			errorHandler.Conflict(c, msg)

		case errors.Is(err, errdefs.ErrInvalidPassword),
			errors.Is(err, errdefs.ErrInvalidEmail):
			errorHandler.BadRequest(c, msg)

		case errors.Is(err, errdefs.ErrUserNotFound):
			errorHandler.NotFoundError(c, msg)

		default:
			// ログだけ詳細、レスポンスは汎用文言
			log.Printf("[RegisterUser] unexpected error: %+v", err)
			errorHandler.InternalServerError(c, msg)
		}

	}

	c.JSON(http.StatusOK, responses.RegisterSuccessResponse{
		Message: "Registration successful",
		User: responses.UserInfo{
			IruyanID: user.IruyanID,
			Name:     user.Name,
			Email:    user.Email,
		},
	})
}

// LogoutHandler ログアウト処理
// @Summary User logout
// @Description Logs out the user based on provided IruyanID
// @Tags auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param iruyanId formData string true "Iruyan ID" default(johndoe)
// @Success 200 {object} responses.LogoutSuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Router /logout [post]
func LogoutHandler(c *gin.Context) {
	iruyanID := c.PostForm("iruyanId")

	var authUsecase usecase.AuthUsecase
	user, err := authUsecase.LogoutUser(iruyanID)
	if err != nil {
		switch err.Error() {
		case fmt.Sprintf("user not found with iruyan_id: %s", iruyanID):
			c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Message: "authentication failed: user not found"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "internal error: " + err.Error()})
		}
		return
	}

	// ログアウト成功レスポンス
	c.JSON(http.StatusOK, responses.LogoutSuccessResponse{
		Message: "Logout successful",
		User: responses.UserInfo{
			IruyanID: user.IruyanID,
			Name:     user.Name,
			Email:    user.Email,
		},
	})
}
