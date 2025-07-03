package routes

import (
	handlers "iruyan-api/handlers/user"
	"iruyan-api/infrastructure"
	"iruyan-api/repositories"
	userusecases "iruyan-api/usecases/user"
	worktimeusecases "iruyan-api/usecases/worktime"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	// --- 依存性の注入 ---
	db := infrastructure.DB // *gorm.DB のインスタンス
	userRepo := repositories.NewUserRepository(db)
	userUsecase := userusecases.NewUserUsecase(userRepo)
	workTimeRepo := repositories.NewWorkTimeRepository(db)
	workTimeUsecase := worktimeusecases.NewWorkTimeUsecase(userRepo, workTimeRepo)
	userHandler := handlers.NewUserHandler(userUsecase, workTimeUsecase)

	userGroup := router.Group("/user/:iruyanId")
	{
		userGroup.GET("", userHandler.PageHandler)
		userGroup.DELETE("/delete", userHandler.DeleteHandler)
		userGroup.GET("/work_info", userHandler.WorkInfoHandler)
		userGroup.GET("/together", userHandler.TogetherTimeHandler)
		userGroup.GET("/ranking", userHandler.RankingHandler)
		userGroup.GET("/recent_log", userHandler.GetRecentLogHandler)
		userGroup.POST("/task", userHandler.TaskHandler)
	}

	fetchGroup := router.Group("/user/fetch")
	{
		fetchGroup.GET("/all", userHandler.GetAllUsersHandler)
		fetchGroup.GET("/:iruyanId", userHandler.GetUserByIruyanIDHandler)
	}
}
