// Package worktime provides HTTP handlers for recording and retrieving work time logs.
package worktime

import (
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/responses"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// EntryHandler godoc
// @Summary Record entry time
// @Description Creates a work time record when a user enters a room
// @Tags worktime
// @Produce json
// @Param iruyanId formData string true "Iruyan ID" default(johndoe)
// @Param roomId formData string true "Room UUID"
// @Param task formData string false "Task description"
// @Success 200 {object} responses.WorkTimeResponseSwagger
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /worktime/entry [post]
func EntryHandler(c *gin.Context) {
	iruyanID := c.PostForm("iruyanId")
	roomIDStr := c.PostForm("roomId")
	task := c.PostForm("task")

	log.Printf("📥 Received Params - iruyanID: %s, roomIDStr: %s, task: %s", iruyanID, roomIDStr, task)

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

	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid room ID format"})
		return
	}

	var existing models.WorkTime
	inRoom, err := existing.IsUserAlreadyInRoom(infrastructure.DB, user.ID, roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: err.Error()})
		return
	}
	if inRoom {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "user is already in the room"})
		return
	}

	now := time.Now()

	workTime := models.WorkTime{
		RoomID:      roomID,
		UserID:      user.ID,
		Task:        task,
		EntryTime:   now,
		LeavingTime: time.Time{},
		Duration:    0,
	}

	if err := infrastructure.DB.Create(&workTime).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: err.Error()})
		return
	}

	resp := responses.WorkTimeResponse{
		RoomID:      workTime.RoomID,
		Task:        workTime.Task,
		EntryTime:   workTime.EntryTime,
		LeavingTime: time.Time{},
		Duration:    0,
	}

	c.JSON(http.StatusOK, resp)
}

// GetLatestEntryHandler godoc
// @Summary Get latest active entry
// @Description Retrieves the latest work time record for a user that has not ended (leaving_time is null)
// @Tags worktime
// @Produce json
// @Param iruyanId query string true "Iruyan ID" default(johndoe)
// @Param roomId query string true "Room UUID"
// @Success 200 {object} responses.WorkTimeResponseSwagger
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Router /worktime/latest [get]
func GetLatestEntryHandler(c *gin.Context) {
	iruyanID := c.Query("iruyanId")
	roomIDStr := c.Query("roomId")

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

	roomID, err := uuid.Parse(roomIDStr)
	if iruyanID == "" || err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid parameters"})
		return
	}

	var workTime models.WorkTime
	if err := infrastructure.DB.
		Where("user_id = ? AND room_id = ? AND leaving_time IS NULL", user.ID, roomID).
		Order("entry_time desc").
		First(&workTime).Error; err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "no active session found"})
		return
	}

	resp := responses.WorkTimeResponse{
		RoomID:      workTime.RoomID,
		Task:        workTime.Task,
		EntryTime:   workTime.EntryTime,
		LeavingTime: workTime.LeavingTime,
		Duration:    workTime.Duration,
	}

	c.JSON(http.StatusOK, resp)
}

// GetRecentLogsHandler godoc
// @Summary Get recent work logs
// @Description Retrieves the most recent N work time records for a user
// @Tags worktime
// @Produce json
// @Param iruyanId query string true "User ID"
// @Param limit query int false "Number of records to return" default(5)
// @Success 200 {array} responses.WorkTimeResponseSwagger
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /worktime/recent [get]
func GetRecentLogsHandler(c *gin.Context) {
	iruyanID := c.Query("iruyanId")
	limitStr := c.DefaultQuery("limit", "5")

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

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid limit"})
		return
	}

	var workTimes []models.WorkTime
	if err := infrastructure.DB.
		Where("user_id = ?", user.ID).
		Order("entry_time desc").
		Limit(limit).
		Find(&workTimes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: err.Error()})
		return
	}

	var result []responses.WorkTimeResponse
	for _, wt := range workTimes {
		result = append(result, responses.WorkTimeResponse{
			RoomID:      wt.RoomID,
			Task:        wt.Task,
			EntryTime:   wt.EntryTime,
			LeavingTime: wt.LeavingTime,
			Duration:    wt.Duration,
		})
	}

	c.JSON(http.StatusOK, result)
}

// GetWeeklyLogsHandler godoc
// @Summary Get weekly work logs
// @Description Retrieves all work time records from the past 7 days for a user
// @Tags worktime
// @Produce json
// @Param iruyanId query string true "User ID"
// @Success 200 {array} responses.WorkTimeResponseSwagger
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /worktime/weekly [get]
func GetWeeklyLogsHandler(c *gin.Context) {
	iruyanID := c.Query("iruyanId")

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

	oneWeekAgo := time.Now().AddDate(0, 0, -6)

	var workTimes []models.WorkTime
	if err := infrastructure.DB.
		Where("user_id = ? AND entry_time >= ?", user.ID, oneWeekAgo).
		Order("entry_time desc").
		Find(&workTimes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: err.Error()})
		return
	}

	var result []responses.WorkTimeResponse
	for _, wt := range workTimes {
		result = append(result, responses.WorkTimeResponse{
			RoomID:      wt.RoomID,
			Task:        wt.Task,
			EntryTime:   wt.EntryTime,
			LeavingTime: wt.LeavingTime,
			Duration:    wt.Duration,
		})
	}

	c.JSON(http.StatusOK, result)
}
