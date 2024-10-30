package room

import (
	"fmt"
	"iruyan-api/handlers/worktime"
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/responses"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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
	userIDStr := c.PostForm("user_id") // ユーザーIDを文字列で取得

	// デバッグ用のログでリクエストパラメータを確認
	log.Printf("Received user_id: %s", userIDStr)

	userIDUint, err := strconv.ParseUint(userIDStr, 10, 32) // uintに変換
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid user ID format",
		})
		return
	}
	userID := uint(userIDUint) // `uint`型にキャスト

	var room models.Room

	// ルームが存在するか確認
	if err := room.FindByID(infrastructure.DB, roomID); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "Room not found",
		})
		return
	}

	// WorkTimeに入室情報を記録（handlers/worktimeに委譲）
	workTime, err := worktime.RecordEntry(userID, roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to record entry to the room",
		})
		return
	}

	// 成功時のレスポンスとしてWorkTime情報を返す
	c.JSON(http.StatusOK, responses.EnterRoomResponse{
		Message:   "Room entry recorded successfully",
		RoomID:    workTime.RoomID.String(),
		RoomName:  room.Name,                          // 部屋の名前を返す
		UserID:    fmt.Sprintf("%d", workTime.UserID), // UserIDを文字列に変換
		EntryTime: workTime.EntryTime,
	})
}

// LeaveRoomHandler Room退室
// @Summary Leave a room
// @Description Allows a user to leave a specific room and records the leaving time
// @Tags room
// @Produce json
// @Param room_id path string true "Room ID"
// @Param user_id formData string true "User ID"
// @Success 200 {object} responses.RoomActionResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /rooms/{room_id}/leave [post]
func LeaveRoomHandler(c *gin.Context) {
	roomID := c.Param("room_id")
	userIDStr := c.PostForm("user_id") // ユーザーIDを文字列で取得

	// デバッグ用のログでリクエストパラメータを確認
	log.Printf("Received user_id: %s", userIDStr)

	userIDUint, err := strconv.ParseUint(userIDStr, 10, 32) // uintに変換
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid user ID format",
		})
		return
	}
	userID := uint(userIDUint) // `uint`型にキャスト

	var room models.Room
	// ルームが存在するか確認
	if err := room.FindByID(infrastructure.DB, roomID); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "Room not found",
		})
		return
	}

	// WorkTimeテーブルから最新の入室記録を取得
	workTime, err := worktime.GetLatestEntry(userID, roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to find entry record",
		})
		return
	}

	// LeavingTimeとDurationを設定
	leavingTime := time.Now()
	workTime.LeavingTime = leavingTime
	workTime.Duration = leavingTime.Sub(workTime.EntryTime)

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
