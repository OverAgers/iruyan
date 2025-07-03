// Package auth provides HTTP handlers for user authentication,
// including login, registration, and logout functionality.
package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"iruyan-api/pkg/errdefs"
	"iruyan-api/utils"

	"iruyan-api/responses"
	usecase "iruyan-api/usecases/auth"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	AuthUsecase usecase.AuthUsecase
}

// [DI] AuthUsecase
func NewAuthHandler(authUsecase usecase.AuthUsecase) *authHandler {
	return &authHandler{
		AuthUsecase: authUsecase,
	}
}

// LoginPageHandler ログイン画面表示
// @Summary Show login page
// @Description Displays the login page with a message
// @Tags auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} responses.ErrorResponse
// @Router /login [get]
func (h *authHandler) LoginPageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, responses.ErrorResponse{
		Message: "Login page accessed successfully",
	})
}

// LoginHandler ログイン処理
// @Summary User login
// @Description Authenticates the user based on iruyanID and password
// @Tags auth
// @Security BearerAuth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param iruyanId formData string true "Iruyan ID" default(johndoe)
// @Param password formData string true "Password" default(pass1234)
// @Success 200 {object} responses.LoginSuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /login [post]
func (h *authHandler) LoginHandler(c *gin.Context) {
	iruyanID := c.PostForm("iruyanId")
	password := c.PostForm("password")

	if iruyanID == "" || password == "" {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "iruyanId and password are required",
		})
		return
	}

	user, err := h.AuthUsecase.LoginUser(iruyanID, password)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Message: "user not found"})
		case errors.Is(err, errdefs.ErrInvalidPassword):
			c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Message: "invalid password"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
		return
	}

	// JWT発行
	token, err := utils.GenerateJWT(user.IruyanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, responses.LoginSuccessResponse{
		Message: "Login successful",
		User: responses.UserInfo{
			IruyanID: user.IruyanID,
			Name:     user.Name,
			Email:    user.Email,
		},
		Token: token,
	})
}

// RegisterPageHandler 新規登録画面表示
// @Summary Show registration page
// @Description Displays the registration page with a message
// @Tags auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} responses.ErrorResponse
// @Router /register [get]
func (h *authHandler) RegisterPageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, responses.ErrorResponse{
		Message: "Register page accessed successfully",
	})
}

// RegisterHandler 新規登録処理
// @Summary User registration
// @Description Registers a new user with the provided details
// @Tags auth
// @Security BearerAuth
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
func (h *authHandler) RegisterHandler(c *gin.Context) {
	iruyanID := c.PostForm("iruyanId")
	password := c.PostForm("password")
	name := c.PostForm("userName")
	email := c.PostForm("email")

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
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Missing required parameter(s): " + strings.Join(missing, ", "),
		})

		return
	}

	user, err := h.AuthUsecase.RegisterUser(name, iruyanID, password, email)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrDuplicateIruyanID):
			c.JSON(http.StatusConflict, responses.ErrorResponse{Message: "iruyan_id is already taken"})
		case errors.Is(err, errdefs.ErrDuplicateEmail):
			c.JSON(http.StatusConflict, responses.ErrorResponse{Message: "email is already registered"})
		case errors.Is(err, errdefs.ErrInvalidPassword):
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid password"})
		case errors.Is(err, errdefs.ErrPasswordTooShort):
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "password must be at least 6 characters long"})
		case errors.Is(err, errdefs.ErrPasswordMissingChars):
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "password must contain at least one letter and one number"})
		case errors.Is(err, errdefs.ErrInvalidEmail):
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid email format"})
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "user not found"})
		default:
			log.Printf("[RegisterUser] unexpected error: %+v", err)
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
		return
	}

	// JWT発行
	token, err := utils.GenerateJWT(user.IruyanID)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrJWTSecretNotSet):
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Message: "JWT secret not set",
			})
		case errors.Is(err, errdefs.ErrJWTInvalidToken):
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Message: "invalid JWT token",
			})
		case errors.Is(err, errdefs.ErrJWTSignToken):
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Message: "failed to sign token",
			})
		case errors.Is(err, errdefs.ErrJWTParseFailure):
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Message: "failed to parse JWT token",
			})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
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
		Token: token,
	})
}

// LogoutHandler ログアウト処理
// @Summary User logout
// @Description Logs out the user based on provided IruyanID
// @Tags auth
// @Security BearerAuth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param iruyanId formData string true "Iruyan ID" default(johndoe)
// @Success 200 {object} responses.LogoutSuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Router /logout [post]
func (h *authHandler) LogoutHandler(c *gin.Context) {
	iruyanID := c.PostForm("iruyanId")

	if iruyanID == "" {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "iruyanId is required",
		})
		return
	}

	user, err := h.AuthUsecase.LogoutUser(iruyanID)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Message: "user not found"})
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
