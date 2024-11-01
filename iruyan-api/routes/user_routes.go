package routes

import (
	"iruyan-api/handlers/user"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	// ユーザー画面
	router.GET("/user/:user_id", user.UserPageHandler)

	// ユーザーの削除
	router.POST("/user/:user_id/delete", user.UserDeleteHandler)

	// 1週間の作業日取得
	router.GET("/user/:user_id/work_info", user.WorkInfoHandler)

	// 一緒に居た時間
	router.GET("/user/:user_id/together", user.TogetherTimeHandler)

	// 集中ランキング
	router.GET("/user/:user_id/ranking", user.RankingHandler)

	// タスクの更新
	router.POST("/user/:user_id/task", user.TaskHandler)
}
