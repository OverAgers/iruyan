package routes

import (
	"iruyan-api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	// ログイン画面表示
	router.GET("/login", handlers.LoginPageHandler)

	// ログイン処理
	router.POST("/login", handlers.LoginHandler)

	// 新規登録画面表示
	router.GET("/register", handlers.RegisterPageHandler)

	// 新規登録処理
	router.POST("/register", handlers.RegisterHandler)

	// ユーザー画面
	router.GET("/user/:user_id", handlers.UserPageHandler)

	// 1週間の作業日取得
	router.GET("/user/:user_id/work_info", handlers.WorkInfoHandler)

	// 一緒に居た時間
	router.GET("/user/:user_id/together", handlers.TogetherTimeHandler)

	// 集中ランキング
	router.GET("/user/:user_id/ranking", handlers.RankingHandler)
}
