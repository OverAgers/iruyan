package routes

import (
	"iruyan-api/handlers/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	// ログイン画面表示
	router.GET("/login", auth.LoginPageHandler)

	// ログイン処理
	router.POST("/login", auth.LoginHandler)

	// 新規登録画面表示
	router.GET("/register", auth.RegisterPageHandler)

	// 新規登録処理
	router.POST("/register", auth.RegisterHandler)

	// ログアウト
	router.POST("/logout", auth.LogoutHandler)
}
