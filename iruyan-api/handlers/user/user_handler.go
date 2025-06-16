package user

import (
	"iruyan-api/handlers/worktime"
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/responses"
	"time"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PageHandler displays the user page.
func PageHandler(c *gin.Context) {
	userIDParam := c.Param("userId")
	userIDUint64, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid user id format",
		})
		return
	}
	userID := uint(userIDUint64)

	var user models.User
	if err = user.FindByID(infrastructure.DB, userID); err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, responses.ErrorResponse{
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Database error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User page accessed successfully",
		"userId":  userID,
	})
}

// DeleteHandler deletes a user.
func DeleteHandler(c *gin.Context) {
	userID := c.Param("userId")
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
		"userId":  userID,
	})
}

// WorkInfoHandler retrieves work logs for the past week.
func WorkInfoHandler(c *gin.Context) {
	userIDParam := c.Param("userId")
	userIDUint64, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid user id format",
		})
		return
	}
	userID := uint(userIDUint64)

	var user models.User
	if err = user.FindByID(infrastructure.DB, userID); err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, responses.ErrorResponse{
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Database error",
		})
		return
	}

	workLogs, err := worktime.GetLogForLastWeek(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to retrieve work logs",
		})
		return
	}

	dailyWorkHours := make(map[string]time.Duration)
	for _, log := range workLogs {
		dateStr := log.EntryTime.Format("2006-01-02")
		dailyWorkHours[dateStr] += log.Duration
	}

	dailyLogs := []responses.DailyWorkLogResponse{}
	for date, hours := range dailyWorkHours {
		dailyLogs = append(dailyLogs, responses.DailyWorkLogResponse{
			Date:  date,
			Hours: hours,
		})
	}

	workInfo := responses.WorkLogForLastWeekResponse{
		UserID:    userID,
		DailyLogs: dailyLogs,
	}

	c.JSON(http.StatusOK, workInfo)
}

// TogetherTimeHandler returns the time spent together with others.
func TogetherTimeHandler(c *gin.Context) {
	userID := c.Param("userId")
	c.JSON(http.StatusOK, gin.H{
		"message": "Together time accessed successfully",
		"userId":  userID,
	})
}

// RankingHandler returns the user's focus ranking.
func RankingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Ranking page accessed successfully",
	})
}

// TaskHandler updates the user's current task.
func TaskHandler(c *gin.Context) {
	userIDParam := c.Param("userId")
	task := c.PostForm("task")

	userIDUint64, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid user id format",
		})
		return
	}
	userID := uint(userIDUint64)

	var user models.User
	if err = user.FindByID(infrastructure.DB, userID); err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, responses.ErrorResponse{
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Database error",
		})
		return
	}

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

// GetRecentLog returns the user's most recent 5 work logs.
func GetRecentLog(c *gin.Context) {
	userIDParam := c.Param("userId")
	userIDUint64, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid user id format",
		})
		return
	}
	userID := uint(userIDUint64)

	var user models.User
	if err := infrastructure.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "User not found",
		})
		return
	}

	workTimes, err := worktime.GetLatestLogs(userID, 5)
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

	workTimeLogs := make([]responses.WorkTimeLog, len(workTimes))
	for i, workTime := range workTimes {
		workTimeLogs[i] = responses.WorkTimeLog{
			EntryTime:   workTime.EntryTime,
			LeavingTime: workTime.LeavingTime,
			Duration:    workTime.Duration,
		}
	}

	c.JSON(http.StatusOK, responses.GetRecentLogResponse{
		Message:     "Get recent log successfully",
		UserID:      userIDParam,
		WorkTimeLog: workTimeLogs,
	})
}
