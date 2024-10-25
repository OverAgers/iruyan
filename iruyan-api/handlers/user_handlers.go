package handlers

import (
	"fmt"
	"net/http"

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
	// フォームデータから値を取得
	username := c.PostForm("username")
	password := c.PostForm("password")
	name := c.PostForm("name")
	icon := c.PostForm("icon")
	task := c.PostForm("task")
	status := c.PostForm("status")
	email := c.PostForm("email")

	// エラーハンドラのインスタンスを作成
	errorHandler := ErrorHandler{}

	// User構造体のインスタンスを作成（この時点でパスワードハッシュ化とバリデーションも行われる）
	user, err := models.NewUser(name, username, password, icon, task, status, email)
	if err != nil {
		// パスワードが4桁でない場合や他のバリデーションに失敗した場合、エラーレスポンスを返す
		errorHandler.BadRequest(c, err.Error())
		return
	}

	// すべてのデータをレスポンスとして表示（パスワードはハッシュ化されているので表示しない）
	response := fmt.Sprintf(
		"Register User: \nUsername: %s\nName: %s\nIcon: %s\nTask: %s\nStatus: %s\nEmail: %s\n",
		user.Username,
		user.Name,
		user.Icon,
		user.Task,
		user.Status,
		user.Email,
	)

	// レスポンスを送信
	c.JSON(http.StatusOK, gin.H{
		"message":  "Registration successful",
		"response": response,
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
