package room

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Room作成
func CreateRoomHandler(c *gin.Context) {
	// Room作成のロジックをここに追加
	c.JSON(http.StatusOK, gin.H{
		"message": "Room created successfully",
	})
}

// Room表示
func ShowRoomHandler(c *gin.Context) {
	roomID := c.Param("room_id")
	// Room情報取得のロジックをここに追加
	c.JSON(http.StatusOK, gin.H{
		"message": "Room details retrieved successfully",
		"room_id": roomID,
	})
}

// Room入室
func EnterRoomHandler(c *gin.Context) {
	roomID := c.Param("room_id")
	// 入室処理のロジックをここに追加
	c.JSON(http.StatusOK, gin.H{
		"message": "Entered the room successfully",
		"room_id": roomID,
	})
}

// Room退室
func LeaveRoomHandler(c *gin.Context) {
	roomID := c.Param("room_id")
	// 退室処理のロジックをここに追加
	c.JSON(http.StatusOK, gin.H{
		"message": "Left the room successfully",
		"room_id": roomID,
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
