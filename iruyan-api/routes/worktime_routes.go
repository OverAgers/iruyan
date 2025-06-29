package routes

import (
	"iruyan-api/handlers/worktime"

	"github.com/gin-gonic/gin"
)

func RegisterWorktimeRoutes(router *gin.Engine) {
	worktimeGroup := router.Group("/worktime")
	{
		worktimeGroup.GET("/fetch/all", worktime.GetAllWorkTimeHandler)

		// POST: 入室処理（エントリ登録）
		worktimeGroup.POST("/entry", worktime.EntryHandler)

		// GET: 最新の入室記録（まだ退室していない場合）
		worktimeGroup.GET("/latest", worktime.GetLatestEntryHandler)

		// GET: 最近の作業ログ（N件）
		worktimeGroup.GET("/recent", worktime.GetRecentLogsHandler)

		// GET: 直近1週間の作業ログ
		worktimeGroup.GET("/weekly", worktime.GetWeeklyLogsHandler)
	}
}
