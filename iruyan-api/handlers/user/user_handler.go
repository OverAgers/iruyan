// Package user provides HTTP handlers related to user operations such as viewing
// profiles, retrieving work logs, updating tasks, and ranking functionality.
package user

import (
	"errors"
	"iruyan-api/pkg/errdefs"
	"iruyan-api/responses"
	userusecase "iruyan-api/usecases/user"
	worktimeusecase "iruyan-api/usecases/worktime"

	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	UserUsecase     userusecase.UserUsecase
	WorkTimeUsecase worktimeusecase.WorkTimeUsecase
}

func NewUserHandler(userUsecase userusecase.UserUsecase, workTimeUsecase worktimeusecase.WorkTimeUsecase) *userHandler {
	return &userHandler{
		UserUsecase:     userUsecase,
		WorkTimeUsecase: workTimeUsecase,
	}
}

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
func (h *userHandler) PageHandler(c *gin.Context) {
	iruyanID := c.Param("iruyanId")

	err := h.UserUsecase.CheckUserExists(iruyanID)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "user not found"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
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
func (h *userHandler) DeleteHandler(c *gin.Context) {
	iruyanID := c.Param("iruyanId")

	err := h.UserUsecase.DeleteUser(iruyanID)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "user not found"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "failed to delete user"})
		}
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
func (h *userHandler) WorkInfoHandler(c *gin.Context) {
	iruyanID := c.Param("iruyanId")

	err := h.UserUsecase.CheckUserExists(iruyanID)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "user not found"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
		return
	}

	var worktimeUsecase worktimeusecase.WorkTimeUsecase
	workInfo, err := worktimeUsecase.GetWorkLogForLastWeek(iruyanID)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrNotSeated):
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "the user is not currently seated"})
		case errors.Is(err, errdefs.ErrSeatMismatch):
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
		return
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
func (h *userHandler) TogetherTimeHandler(c *gin.Context) {
	iruyanID := c.Param("iruyanId")

	err := h.UserUsecase.CheckUserExists(iruyanID)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "user not found"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
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
func (h *userHandler) RankingHandler(c *gin.Context) {
	iruyanID := c.Param("iruyanId")

	err := h.UserUsecase.CheckUserExists(iruyanID)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "user not found"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
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
func (h *userHandler) TaskHandler(c *gin.Context) {
	iruyanID := c.Param("iruyanId")
	task := c.PostForm("task")

	if err := h.WorkTimeUsecase.UpdateTask(iruyanID, task); err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "user not found"})
		case errors.Is(err, errdefs.ErrNotInRoom):
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "user is not currently in a room"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
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
func (h *userHandler) GetRecentLogHandler(c *gin.Context) {
	iruyanID := c.Param("iruyanId")

	logs, err := h.WorkTimeUsecase.GetRecentLogs(iruyanID, 5)
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

	c.JSON(http.StatusOK, responses.GetRecentLogResponse{
		Message:     "Get recent log successfully",
		IruyanID:    iruyanID,
		WorkTimeLog: logs,
	})
}

// GetAllUsersHandler godoc
// @Summary Get all users
// @Description Retrieves a list of all registered users
// @Tags user
// @Success 200 {array} responses.User
// @Failure 500 {object} responses.ErrorResponse
// @Router /user/fetch/all [get]
func (h *userHandler) GetAllUsersHandler(c *gin.Context) {
	users, err := h.UserUsecase.GetAllUsers()
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
func (h *userHandler) GetUserByIruyanIDHandler(c *gin.Context) {
	iruyanId := c.Param("iruyanId")

	user, err := h.UserUsecase.GetUserByIruyanID(iruyanId)
	if err != nil {
		if errors.Is(err, errdefs.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, responses.ErrorResponse{
				Message: "User not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Message: "Failed to retrieve user",
			})
		}
		return
	}

	responseUser := responses.User{
		Email:    user.Email,
		IruyanID: user.IruyanID,
		Name:     user.Name,
	}

	c.JSON(http.StatusOK, responseUser)
}
