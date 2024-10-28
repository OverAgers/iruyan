package handlers

import (
	"net/http"

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
	// 仮の処理
	username := c.PostForm("username")
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    username,
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
	status := c.PostForm("status")
	email := c.PostForm("email")

	errorHandler := ErrorHandler{}

	// GORMを使ってデータベースからusernameの重複を確認しつつユーザーインスタンスを生成
	user, err := models.NewUser(infrastructure.DB, name, username, password, task, status, email)
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
			"status":   user.Status,
			"email":    user.Email,
		},
	})
}

// ユーザー画面表示
func UserPageHandler(c *gin.Context) {
	userID := c.Param("user_id")
	c.JSON(http.StatusOK, gin.H{
		"message": "User page accessed successfully",
		"user_id": userID,
	})
}

// 1週間の作業日取得
func WorkInfoHandler(c *gin.Context) {
	userID := c.Param("user_id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Work info accessed successfully",
		"user_id": userID,
	})
}

// 一緒に居た時間
func TogetherTimeHandler(c *gin.Context) {
	userID := c.Param("user_id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Together time accessed successfully",
		"user_id": userID,
	})
}

// 集中ランキング取得
func RankingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Ranking page accessed successfully",
	})
}
