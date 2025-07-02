// Package worktime provides HTTP handlers for recording and retrieving work time logs.
package worktime

import (
	"errors"
	"iruyan-api/pkg/errdefs"
	usecase "iruyan-api/usecases/worktime"

	"iruyan-api/responses"
	"net/http"
	"strconv"

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

	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid room ID format"})
		return
	}

	var worktimeUsecase usecase.WorkTimeUsecase
	workTime, err := worktimeUsecase.EnterRoom(iruyanID, roomID, task)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "user not found"})
		case errors.Is(err, errdefs.ErrAlreadyInRoom):
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "user is already in the room"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
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

	roomID, err := uuid.Parse(roomIDStr)
	if iruyanID == "" || err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid parameters"})
		return
	}

	var worktimeUsecase usecase.WorkTimeUsecase
	workTime, err := worktimeUsecase.GetLatestEntry(iruyanID, roomID)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "user not found"})
		case errors.Is(err, errdefs.ErrNoActiveSession):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "no active session found"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
		return
	}

	resp := responses.WorkTimeResponse{
		IruyanID:    workTime.User.IruyanID,
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
// @Param iruyanId query string true "Iruyan ID" default(johndoe)
// @Param limit query int false "Number of records to return" default(5)
// @Success 200 {array} responses.WorkTimeResponseSwagger
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /worktime/recent [get]
func GetRecentLogsHandler(c *gin.Context) {
	iruyanID := c.Query("iruyanId")
	limitStr := c.DefaultQuery("limit", "5")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid limit"})
		return
	}

	var worktimeUsecase usecase.WorkTimeUsecase
	logs, err := worktimeUsecase.GetRecentLogs(iruyanID, limit)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "user not found"})
		case errors.Is(err, errdefs.ErrNotInRoom):
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "no worktime records found"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
		return
	}

	c.JSON(http.StatusOK, logs)
}

// GetWeeklyLogsHandler godoc
// @Summary Get weekly work logs
// @Description Retrieves all work time records from the past 7 days for a user
// @Tags worktime
// @Produce json
// @Param iruyanId query string true "Iruyan ID" default(johndoe)
// @Success 200 {array} responses.WorkTimeResponseSwagger
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /worktime/weekly [get]
func GetWeeklyLogsHandler(c *gin.Context) {
	iruyanID := c.Query("iruyanId")
	var worktimeUsecase usecase.WorkTimeUsecase

	logs, err := worktimeUsecase.GetWeeklyLogs(iruyanID)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "user not found"})
		case errors.Is(err, errdefs.ErrNotInRoom): // 仮に週ログが見つからないときの独自定義
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "no worktime records found"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
		return
	}

	c.JSON(http.StatusOK, logs)
}

// GetAllWorkTimeHandler godoc
// @Summary Get all WorkTime records
// @Description Retrieves all work time records from the database
// @Tags worktime
// @Produce json
// @Success 200 {object} responses.WorkTimeListResponseSwagger
// @Failure 500 {object} responses.ErrorResponse
// @Router /worktime/fetch/all [get]
func GetAllWorkTimeHandler(c *gin.Context) {
	var worktimeUsecase usecase.WorkTimeUsecase
	workTimes, err := worktimeUsecase.GetAllWorkTimes()
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "user not found"})
		case errors.Is(err, errdefs.ErrNotInRoom): // 仮に週ログが見つからないときの独自定義
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "no worktime records found"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
		return
	}

	c.JSON(http.StatusOK, responses.WorkTimeListResponse{
		WorkTimes: workTimes,
	})
}

// GetWorkTimeByIruyanIDHandler godoc
// @Summary Get WorkTime records for a specific user
// @Description Retrieves all work time records for a user identified by iruyanID
// @Tags worktime
// @Produce json
// @Param iruyanId path string true "Iruyan ID" default(johndoe)
// @Success 200 {object} responses.WorkTimeListResponseSwagger
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /worktime/fetch/{iruyanId} [get]
func GetWorkTimeByIruyanIDHandler(c *gin.Context) {
	iruyanID := c.Param("iruyanId")

	var worktimeUsecase usecase.WorkTimeUsecase
	workTimes, err := worktimeUsecase.GetWorkTimeByIruyanID(iruyanID)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "user not found"})
		case errors.Is(err, errdefs.ErrNotInRoom):
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "no worktime records found"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
		return
	}

	c.JSON(http.StatusOK, responses.WorkTimeListResponse{
		WorkTimes: workTimes,
	})
}
