package routes

import (
	"iruyan-api/handlers/user"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user/:iruyanId")
	{
		userGroup.GET("", user.PageHandler)
		userGroup.DELETE("/delete", user.DeleteHandler)
		userGroup.GET("/work_info", user.WorkInfoHandler)
		userGroup.GET("/together", user.TogetherTimeHandler)
		userGroup.GET("/ranking", user.RankingHandler)
		userGroup.GET("/recent_log", user.GetRecentLog)
		userGroup.POST("/task", user.TaskHandler)
	}

	fetchGroup := router.Group("/user/fetch")
	{
		fetchGroup.GET("/all", user.GetAllUsersHandler)
		fetchGroup.GET("/:iruyanId", user.GetUserByIruyanIDHandler)
	}
}
