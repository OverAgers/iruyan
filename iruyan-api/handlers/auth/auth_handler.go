// Package auth provides HTTP handlers for user authentication,
// including login, registration, and logout functionality.
package auth

import (
	"errors"
	"net/http"
	"strings"

	errorhandler "iruyan-api/handlers/error"
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/repository"
	"iruyan-api/responses"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

	// --- [1] バリデーションチェック（空欄チェック） ---
	if iruyanID == "" || password == "" {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "iruyanId and password are required",
		})
		return
	}

	// --- [2] ユーザー検索 ---
	var user models.User
	result := infrastructure.DB.Where("iruyan_id = ?", iruyanID).First(&user)

	if result.Error != nil {
		// ユーザーが見つからなかった場合
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
				Message: "authentication failed: user not found",
			})
			return
		}

		// データベース関連のその他のエラー（500）
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "internal error: failed to retrieve user",
		})
		return
	}

	// --- [3] パスワードチェック ---
	if !user.CheckPassword(password) {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Message: "authentication failed: invalid password",
		})
		return
	}

	// --- [4] 認証成功時 ---
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
	if iruyanID == "" || password == "" || name == "" || email == "" {
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

		errorHandler.BadRequest(c, "Missing required parameter(s): "+strings.Join(missing, ", "))
		return
	}

	user, err := repository.CreateUser(infrastructure.DB, name, iruyanID, password, email)
	if err != nil {
		switch err {
		case repository.ErrDuplicateIruyanID, repository.ErrDuplicateEmail:
			errorHandler.Conflict(c, err.Error())
		default:
			errorHandler.BadRequest(c, err.Error())
		}
		return
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

	// エラーハンドラをインスタンス化
	errorHandler := errorhandler.ErrorHandler{}

	// ユーザーが存在するか確認
	var user models.User
	if err := infrastructure.DB.Where("iruyan_id = ?", iruyanID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			errorHandler.NotFoundError(c, "User not found")
		} else {
			errorHandler.InternalServerError(c, "Database error")
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
