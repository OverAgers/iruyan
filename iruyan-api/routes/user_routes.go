package routes

import (
	"iruyan-api/handlers/user"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	// ユーザー画面
	router.GET("/user/:iruyanId", user.PageHandler)

	// ユーザーの削除
	router.DELETE("/user/:iruyanId/delete", user.DeleteHandler)

	// 1週間の作業日取得
	router.GET("/user/:iruyanId/work_info", user.WorkInfoHandler)

	// 一緒に居た時間
	router.GET("/user/:iruyanId/together", user.TogetherTimeHandler)

	// 集中ランキング
	router.GET("/user/:iruyanId/ranking", user.RankingHandler)

	// 直近5回分の作業時間を取得
	router.GET("/user/:iruyanId/recent_log", user.GetRecentLog)

	// タスクの更新
	router.POST("/user/:iruyanId/task", user.TaskHandler)

	// ユーザーの全件取得
	router.GET("/user/fetch/all", user.GetAllUsersHandler)

	// 特定ユーザーの取得
	router.GET("/user/fetch/:iruyanId", user.GetUserByIruyanIDHandler)
}
