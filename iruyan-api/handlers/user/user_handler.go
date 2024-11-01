package user

import (
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/responses"
	"iruyan-api/handlers/worktime"

	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	if err = user.FindByID(infrastructure.DB, userID); err != nil {
		if err.Error() == "user not found" {
			// ユーザがデータベースに存在しない場合
			c.JSON(http.StatusNotFound, responses.ErrorResponse{
				Message: "User not found",
			})
			return
		} else {
			// その他のサーバ側のエラーの場合
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
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

	// ユーザが存在するか確認
	var user models.User
	if err = user.FindByID(infrastructure.DB, userID); err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, responses.ErrorResponse{
				Message: "User not found",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Message: "Database error",
			})
			return
		}
	}

	// 1週間のログを取得
	workLogs, err := worktime.GetLogForLastWeek(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to retrieve work logs",
		})
		return
	}

	// 日毎の作業時間を集計
	dailyWorkHours := make(map[string]time.Duration) // 日付をキー、作業時間を値とするマップ
	for _, log := range workLogs {
		dateStr := log.EntryTime.Format("2006-01-02") // EntryTimeを利用
		dailyWorkHours[dateStr] += log.Duration // Durationを時間に変換して加算
	}

	// レスポンスのデータを整形
	dailyLogs := []responses.DailyWorkLogResponse{}
	for date, hours := range dailyWorkHours {
		dailyLogs = append(dailyLogs, responses.DailyWorkLogResponse{
			Date:  date,
			Hours: hours,
		})
	}

	workInfo := responses.WorkLogForLastWeekResponse{
		UserID:   userID,
		DailyLogs: dailyLogs, // スライスを格納
	}

	// レスポンス成功時メッセージ
	c.JSON(http.StatusOK, workInfo)
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

// タスク内容の更新
func TaskHandler(c *gin.Context) {
	userIDParam := c.Param("user_id")
	task := c.PostForm("task")

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
	if err = user.FindByID(infrastructure.DB, userID); err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, responses.ErrorResponse{
				Message: "User not found",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Message: "Database error",
			})
			return
		}
	}

	// WorkTimeテーブルから、入室中のレコード（LeavingTimeが設定されていない）を取得
	var workTime models.WorkTime
	if err = infrastructure.DB.Where("user_id = ? AND leaving_time IS NULL", userID).First(&workTime).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Message: "User is not currently in a room",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Database error",
		})
		return
	}

	// タスク内容を更新
	workTime.Task = task
	if err = infrastructure.DB.Save(&workTime).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to update task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
	})
}
