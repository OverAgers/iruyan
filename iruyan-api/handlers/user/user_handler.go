// Package user provides HTTP handlers related to user operations such as viewing
// profiles, retrieving work logs, updating tasks, and ranking functionality.
package user

import (
	"iruyan-api/handlers/worktime"
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/responses"
	"strings"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PageHandler godoc
// @Summary Show user page
// @Description Display user page for a specific user ID
// @Tags user
// @Param iruyanId path string true "Iruyan ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /user/{iruyanId} [get]
func PageHandler(c *gin.Context) {
	iruyanID := c.Param("iruyanId")

	var user models.User
	err := user.FindByIruyanID(infrastructure.DB, iruyanID)
	if err != nil {
		if strings.HasPrefix(err.Error(), "user not found with iruyan_id:") {
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
		"message":  "User page accessed successfully",
		"iruyanId": iruyanID,
	})
}

// DeleteHandler godoc
// @Summary Delete a user
// @Description Delete user by ID
// @Tags user
// @Param iruyanId path string true "UIruyan ID"
// @Success 200 {object} map[string]interface{}
// @Router /user/{iruyanId}/delete [delete]
func DeleteHandler(c *gin.Context) {
	iruyanID := c.Param("iruyanId")

	var user models.User
	if err := user.DeleteByIruyanID(infrastructure.DB, iruyanID); err != nil {
		if strings.HasPrefix(err.Error(), "user not found with iruyan_id:") {
			c.JSON(http.StatusNotFound, responses.ErrorResponse{
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to delete user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "User deleted successfully",
		"iruyanId": iruyanID,
	})
}

// WorkInfoHandler godoc
// @Summary Get user's work log for the past week
// @Description Retrieve user's work logs grouped by day (Duration as ISO 8601 format string)
// @Tags user
// @Param iruyanId path string true "Iruyan ID"
// @Success 200 {object} responses.WorkLogForLastWeekResponseSwagger
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /user/{iruyanId}/work_info [get]
func WorkInfoHandler(c *gin.Context) {
	iruyanID := c.Param("iruyanId")

	var user models.User
	err := user.FindByIruyanID(infrastructure.DB, iruyanID)
	if err != nil {
		if strings.HasPrefix(err.Error(), "user not found with iruyan_id:") {
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

	workLogs, err := worktime.GetLogForLastWeek(user.ID)
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
		IruyanID:  iruyanID,
		DailyLogs: dailyLogs,
	}

	c.JSON(http.StatusOK, workInfo)
}

// TogetherTimeHandler godoc
// @Summary Get time spent together
// @Description Returns mock together time
// @Tags user
// @Param iruyanId path string true "Iruyan ID"
// @Success 200 {object} map[string]interface{}
// @Router /user/{iruyanId}/together [get]
func TogetherTimeHandler(c *gin.Context) {
	iruyanID := c.Param("iruyanId")

	var user models.User
	err := user.FindByIruyanID(infrastructure.DB, iruyanID)
	if err != nil {
		if strings.HasPrefix(err.Error(), "user not found with iruyan_id:") {
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
		"message":  "Together time accessed successfully",
		"iruyanId": iruyanID,
	})
}

// RankingHandler godoc
// @Summary Get user ranking
// @Description Returns focus ranking of the user
// @Tags user
// @Param iruyanId path string true "Iruyan ID"
// @Success 200 {object} map[string]interface{}
// @Router /user/{iruyanId}/ranking [get]
func RankingHandler(c *gin.Context) {
	iruyanID := c.Param("iruyanId")

	var user models.User
	err := user.FindByIruyanID(infrastructure.DB, iruyanID)
	if err != nil {
		if strings.HasPrefix(err.Error(), "user not found with iruyan_id:") {
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
		"message": "Ranking page accessed successfully",
	})
}

// TaskHandler godoc
// @Summary Update current task
// @Description Updates the user's current task for ongoing work session
// @Tags user
// @Param iruyanId path int true "User ID"
// @Param task formData string true "Task Description"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /user/{iruyanId}/task [post]
func TaskHandler(c *gin.Context) {
	iruyanID := c.Param("iruyanId")
	task := c.PostForm("task")

	var user models.User
	err := user.FindByIruyanID(infrastructure.DB, iruyanID)
	if err != nil {
		if strings.HasPrefix(err.Error(), "user not found with iruyan_id:") {
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
	if err = infrastructure.DB.Where("user_id = ? AND leaving_time IS NULL", user.ID).First(&workTime).Error; err != nil {
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

// GetRecentLog godoc
// @Summary Get user's recent work logs
// @Description Returns the latest 5 work logs for the user (Duration as ISO 8601 format string)
// @Tags user
// @Param iruyanId path string true "Iruyan ID"
// @Success 200 {object} responses.GetRecentLogResponseSwagger
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /user/{iruyanId}/recent_log [get]
func GetRecentLog(c *gin.Context) {
	iruyanID := c.Param("iruyanId")

	var user models.User
	err := user.FindByIruyanID(infrastructure.DB, iruyanID)
	if err != nil {
		if strings.HasPrefix(err.Error(), "user not found with iruyan_id:") {
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

	workTimes, err := worktime.GetLatestLogs(user.ID, 5)
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
		IruyanID:    iruyanID,
		WorkTimeLog: workTimeLogs,
	})
}

// GetAllUsersHandler godoc
// @Summary Get all users
// @Description Retrieves a list of all registered users
// @Tags user
// @Success 200 {array} responses.User
// @Failure 500 {object} responses.ErrorResponse
// @Router /user/fetch/all [get]
func GetAllUsersHandler(c *gin.Context) {
	users, err := models.GetAllUsers(infrastructure.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to retrieve users",
		})
		return
	}

	responseUsers := make([]responses.User, 0, len(users))
	for _, user := range users {
		responseUsers = append(responseUsers, responses.User{
			Email:    user.Email,
			IruyanID: user.IruyanID,
			Name:     user.Name,
		})
	}

	c.JSON(http.StatusOK, responseUsers)
}

// GetUserByIruyanIDHandler godoc
// @Summary Get user by Iruyan ID
// @Description Retrieves a user based on the provided Iruyan ID
// @Tags user
// @Param iruyanId path string true "Iruyan ID"
// @Success 200 {object} models.User
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /user/fetch/{iruyanId} [get]
func GetUserByIruyanIDHandler(c *gin.Context) {
	iruyanId := c.Param("iruyanId")
	var user models.User
	if err := user.FindByIruyanID(infrastructure.DB, iruyanId); err != nil {
		if err.Error() == "user not found with iruyan_id: "+iruyanId {
			c.JSON(http.StatusNotFound, responses.ErrorResponse{
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to retrieve user",
		})
		return
	}

	responseUser := responses.User{
		Email:    user.Email,
		IruyanID: user.IruyanID,
		Name:     user.Name,
	}

	c.JSON(http.StatusOK, responseUser)
}
