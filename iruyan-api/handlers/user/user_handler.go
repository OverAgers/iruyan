package user

import (
	"net/http"
	"iruyan-api/handlers/worktime"
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/responses"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

// ユーザー画面表示
func UserPageHandler(c *gin.Context) {
	userIDParam := c.Param("user_id")
	// user_idをuint型に変換してuserID変数に保存
	userIDUint64, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid user id format",
		})
		return 
	}
	userID := uint(userIDUint64)

	var user models.User
	// ユーザIDがDBに存在するか確認する
	// ユーザが見つからない場合はエラーを返す
	if err = user.FindByID(infrastructure.DB, userID); err != nil{
		if err.Error() == "user not found" {
			// ユーザがデータベースに存在しない場合
			c.JSON(http.StatusNotFound, responses.ErrorResponse {
				Message: "User not found",
			})
			return 
		} else {
			// その他のサーバ側のエラーの場合
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse {
				Message: "Database error",
			})
			return  
		}
	}

	// ユーザが存在する場合は200を返す
	c.JSON(http.StatusOK, gin.H{
		"message": "User page accessed successfully",
		"user_id": userID,
	})
}

// ユーザーの削除
func UserDeleteHandler(c *gin.Context) {
	userID := c.Param("user_id")
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
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

// 直近5回分の作業時間を取得
func GetRecentLog(c *gin.Context) {
	userIDParam := c.Param("user_id")
	// user_idをuint型に変換してuserID変数に保存
	userIDUint64, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid user id format",
		})
		return 
	}
	userID := uint(userIDUint64)

	// ユーザーが存在するか確認
	var user models.User
	if err := infrastructure.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "User not found",
		})
		return
	}
	
	// WorkTimeテーブルからユーザの直近5回の入室記録を取得
	workTimes, err := worktime.GetLatestLogs(userID, 5)	// 5件取得
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Message: "No worktime records found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Message: "Failed to find worktime record",
			})
		}
		return
	}

	// WorkTimeLog型に変換
	workTimeLogs := make([]responses.WorkTimeLog, len(workTimes))
	for i, workTime := range workTimes {
		workTimeLogs[i] = responses.WorkTimeLog{
			EntryTime:   workTime.EntryTime,
			LeavingTime: workTime.LeavingTime, // *time.Time であることを仮定
			Duration:    workTime.Duration,
		}
	}

	// 成功時のレスポンスを返す
	c.JSON(http.StatusOK, responses.GetRecentLogResponse{
		Message:      "Get recent log successfully",
		UserID:       userIDParam,
		WorkTimeLog:  workTimeLogs,
	})
}