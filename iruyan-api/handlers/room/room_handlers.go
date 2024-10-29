package room

import (
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/responses"
	"net/http"

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

// EnterRoomHandler Room入室
// @Summary Enter a room
// @Description Allows a user to enter a specific room
// @Tags room
// @Produce json
// @Param room_id path string true "Room ID"
// @Success 200 {object} responses.RoomActionResponse
// @Failure 404 {object} responses.ErrorResponse
// @Router /rooms/{room_id}/enter [post]
func EnterRoomHandler(c *gin.Context) {
	roomID := c.Param("room_id")
	var room models.Room

	if err := room.FindByID(infrastructure.DB, roomID); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "Room not found",
		})
		return
	}

	c.JSON(http.StatusOK, responses.RoomActionResponse{
		Message: "Entered the room successfully",
		RoomID:  roomID,
	})
}

// LeaveRoomHandler Room退室
// @Summary Leave a room
// @Description Allows a user to leave a specific room
// @Tags room
// @Produce json
// @Param room_id path string true "Room ID"
// @Success 200 {object} responses.RoomActionResponse
// @Failure 404 {object} responses.ErrorResponse
// @Router /rooms/{room_id}/leave [delete]
func LeaveRoomHandler(c *gin.Context) {
	roomID := c.Param("room_id")
	var room models.Room

	if err := room.FindByID(infrastructure.DB, roomID); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "Room not found",
		})
		return
	}

	c.JSON(http.StatusOK, responses.RoomActionResponse{
		Message: "Left the room successfully",
		RoomID:  roomID,
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
