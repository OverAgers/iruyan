package routes

import (
	handlers "iruyan-api/handlers/worktime"
	"iruyan-api/infrastructure"
	"iruyan-api/repositories"
	usecases "iruyan-api/usecases/worktime"

	"github.com/gin-gonic/gin"
)

func RegisterWorktimeRoutes(router *gin.Engine) {
	// --- 依存性の注入 ---
	db := infrastructure.DB // *gorm.DB のインスタンス
	workTimeRepo := repositories.NewWorkTimeRepository(db)
	userRepo := repositories.NewUserRepository(db)
	workTimeUsecase := usecases.NewWorkTimeUsecase(userRepo, workTimeRepo)
	workTimeHandler := handlers.NewWorkTimeHandler(workTimeUsecase)

	worktimeGroup := router.Group("/worktime")
	{
		worktimeGroup.GET("/fetch/all", workTimeHandler.GetAllWorkTimeHandler)

		worktimeGroup.GET("/fetch/:iruyanId", workTimeHandler.GetWorkTimeByIruyanIDHandler)

		// POST: 入室処理（エントリ登録）
		worktimeGroup.POST("/entry", workTimeHandler.EntryHandler)

		// GET: 最新の入室記録（まだ退室していない場合）
		worktimeGroup.GET("/latest", workTimeHandler.GetLatestEntryHandler)

		// GET: 最近の作業ログ（N件）
		worktimeGroup.GET("/recent", workTimeHandler.GetRecentLogsHandler)

		// GET: 直近1週間の作業ログ
		worktimeGroup.GET("/weekly", workTimeHandler.GetWeeklyLogsHandler)
	}
}
