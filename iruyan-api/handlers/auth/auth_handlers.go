package auth

import (
	"net/http"

	errorhandler "iruyan-api/handlers/error" // errorhandlerとしてインポート
	"iruyan-api/infrastructure"
	"iruyan-api/models"

	"github.com/gin-gonic/gin"
)

// ログイン画面表示
func LoginPageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login page accessed successfully",
	})
}

// ログイン処理
func LoginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// ユーザーをデータベースから取得
	var user models.User
	result := infrastructure.DB.Where("username = ?", username).First(&user)

	// ユーザーが見つからない場合のエラーハンドリング
	if result.Error != nil {
		errorHandler := errorhandler.ErrorHandler{}
		errorHandler.Unauthorized(c, "authentication failed: invalid username")
		return
	}

	// パスワードの比較をモデルのメソッドで行う
	if !user.CheckPassword(password) {
		errorHandler := errorhandler.ErrorHandler{}
		errorHandler.Unauthorized(c, "authentication failed: invalid password")
		return
	}

	// ログイン成功時のレスポンス
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"username": user.Username,
			"name":     user.Name,
			"task":     user.Task,
			"email":    user.Email,
		},
	})
}

// 新規登録画面表示
func RegisterPageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Register page accessed successfully",
	})
}

// 新規登録処理
func RegisterHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	name := c.PostForm("name")
	task := c.PostForm("task")
	email := c.PostForm("email")

	errorHandler := errorhandler.ErrorHandler{}

	// GORMを使ってデータベースからusernameとemailの重複を確認しつつユーザーインスタンスを生成
	user, err := models.NewUser(infrastructure.DB, name, username, password, task, email)
	if err != nil {
		errorHandler.BadRequest(c, err.Error())
		return
	}

	// ユーザーをデータベースに保存
	result := infrastructure.DB.Create(user)
	if result.Error != nil {
		errorHandler.InternalServerError(c, "Failed to save user to database: "+result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful",
		"user": gin.H{
			"username": user.Username,
			"name":     user.Name,
			"task":     user.Task,
			"email":    user.Email,
		},
	})
}
