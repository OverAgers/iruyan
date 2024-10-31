package room

import (
	"fmt"
	"iruyan-api/handlers/worktime"
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/responses"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateRoomHandler Room作成
// @Summary Create a new room
// @Description Creates a room with the provided name
// @Tags room
// @Accept x-www-form-urlencoded
// @Produce json
// @Param name formData string true "Room Name"
// @Success 200 {object} responses.RoomCreateResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /rooms [post]
func CreateRoomHandler(c *gin.Context) {
	roomName := c.PostForm("name")

	room := models.Room{Name: roomName}

	if err := room.ValidateName(); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: err.Error()})
		return
	}

	if err := infrastructure.DB.Create(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to create room: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.RoomCreateResponse{
		Message:  "Room created successfully",
		RoomID:   room.ID,
		RoomName: room.Name,
	})
}

// GetRoomHandler Room表示
// @Summary Get room details
// @Description Retrieves the details of a specific room
// @Tags room
// @Produce json
// @Param room_id path string true "Room ID"
// @Success 200 {object} responses.RoomDetailResponse
// @Failure 404 {object} responses.ErrorResponse
// @Router /rooms/{room_id} [get]
func GetRoomHandler(c *gin.Context) {
	roomID := c.Param("room_id")
	c.JSON(http.StatusOK, responses.RoomDetailResponse{
		Message: "Room details retrieved successfully",
		RoomID:  roomID,
	})
}

// EnterRoomHandler 部屋への入室処理
// @Summary Enter a room
// @Description Records the entry time of a user entering a specific room and stores the entry information in WorkTime.
// @Tags room
// @Accept json
// @Produce json
// @Param room_id path string true "Room ID"
// @Param user_id path string true "User ID"
// @Success 200 {object} responses.RoomActionResponse "Entered the room successfully"
// @Failure 404 {object} responses.ErrorResponse "Room not found"
// @Failure 500 {object} responses.ErrorResponse "Failed to record entry to the room"
// @Router /rooms/{room_id}/enter/{user_id} [post]
func EnterRoomHandler(c *gin.Context) {
	roomID := c.Param("room_id")
	userIDStr := c.PostForm("user_id")

	// ユーザーIDを文字列からuintに変換
	userIDUint, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid user ID format",
		})
		return
	}
	userID := uint(userIDUint)

	var room models.Room
	// ルームが存在するか確認
	if err := room.FindByID(infrastructure.DB, roomID); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "Room not found",
		})
		return
	}

	// WorkTimeに入室情報を記録
	workTime, err := worktime.RecordEntry(userID, roomID)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, responses.ErrorResponse{
				Message: "User not found",
			})
		} else if err.Error() == "user is already in the room" {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Message: "User is already in the room",
			})
		} else {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Message: "Failed to record entry to the room",
			})
		}
		return
	}

	// 成功時のレスポンスとしてWorkTime情報を返す
	c.JSON(http.StatusOK, responses.EnterRoomResponse{
		Message:   "Room entry recorded successfully",
		RoomID:    workTime.RoomID.String(),
		RoomName:  room.Name,
		UserID:    fmt.Sprintf("%d", workTime.UserID),
		EntryTime: workTime.EntryTime,
	})
}

// LeaveRoomHandler Room退室処理
// @Summary Leave a room
// @Description Allows a user to leave a specific room and records the leaving time
// @Tags room
// @Produce json
// @Param room_id path string true "Room ID"
// @Param user_id formData string true "User ID"
// @Param duration formData string true "Duration (e.g., 1h2m3s)"
// @Success 200 {object} responses.LeaveRoomResponse "Left the room successfully"
// @Failure 400 {object} responses.ErrorResponse "Invalid duration format or other validation errors"
// @Failure 404 {object} responses.ErrorResponse "Room or user not found"
// @Failure 500 {object} responses.ErrorResponse "Failed to record leaving time"
// @Router /rooms/{room_id}/leave [post]
func LeaveRoomHandler(c *gin.Context) {
	roomID := c.Param("room_id")
	userIDStr := c.PostForm("user_id")
	durationStr := c.PostForm("duration") // DurationをPostFormで取得

	// ユーザーIDをuintに変換
	userIDUint, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid user ID format",
		})
		return
	}
	userID := uint(userIDUint)

	// Durationをtime.Duration型に変換
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid duration format. Use format like '1h2m3s'.",
		})
		return
	}

	// ルームが存在するか確認
	var room models.Room
	if err := room.FindByID(infrastructure.DB, roomID); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "Room not found",
		})
		return
	}

	// ユーザーが存在するか確認
	var user models.User
	if err := infrastructure.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "User not found",
		})
		return
	}

	// WorkTimeテーブルから最新の入室記録を取得（LeavingTimeがNULLのレコードを取得）
	workTime, err := worktime.GetLatestEntry(userID, roomID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Message: "User is not currently in the room",
			})
		} else {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Message: "Failed to find entry record",
			})
		}
		return
	}

	// LeavingTimeとDurationを設定
	leavingTime := time.Now()
	workTime.LeavingTime = leavingTime
	workTime.Duration = duration // クライアントから受け取ったDurationを使用

	// Durationがエントリー時間から計算されるDurationより短ければエラーを返す
	if duration < leavingTime.Sub(workTime.EntryTime) {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid duration: Duration is shorter than time spent in room",
		})
		return
	}

	// 退室記録をデータベースに保存
	if err := infrastructure.DB.Save(workTime).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to record leaving time",
		})
		return
	}

	// 成功時のレスポンス
	c.JSON(http.StatusOK, responses.LeaveRoomResponse{
		Message:     "Left the room successfully",
		RoomID:      roomID,
		RoomName:    room.Name,
		UserID:      fmt.Sprintf("%d", workTime.UserID),
		EntryTime:   workTime.EntryTime,
		LeavingTime: leavingTime,
		Duration:    workTime.Duration,
	})
}

// 着席
func TakeSeatHandler(c *gin.Context) {
	roomID := c.Param("room_id")
	seatID := c.Param("seat_id")
	// 着席処理のロジックをここに追加
	c.JSON(http.StatusOK, gin.H{
		"message": "Seated successfully",
		"room_id": roomID,
		"seat_id": seatID,
	})
}

// 離席
func LeaveSeatHandler(c *gin.Context) {
	roomID := c.Param("room_id")
	seatID := c.Param("seat_id")
	// 離席処理のロジックをここに追加
	c.JSON(http.StatusOK, gin.H{
		"message": "Left the seat successfully",
		"room_id": roomID,
		"seat_id": seatID,
	})
}
