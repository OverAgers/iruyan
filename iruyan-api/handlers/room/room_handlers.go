package room

import (
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Room作成
func CreateRoomHandler(c *gin.Context) {
	roomName := c.PostForm("name")

	// 新しいRoomインスタンスを作成
	room := models.Room{Name: roomName}

	// ルーム名のバリデーション
	if err := room.ValidateName(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// データベースにRoomを保存
	if err := infrastructure.DB.Create(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create room: " + err.Error(),
		})
		return
	}

	// 作成成功レスポンス
	c.JSON(http.StatusOK, gin.H{
		"message":   "Room created successfully",
		"room_id":   room.ID,
		"room_name": room.Name,
	})
}

// Room表示
func GetRoomHandler(c *gin.Context) {
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

	// Roomのモデルを定義
	var room models.Room

	// データベースからroom_idでRoomを検索
	if err := infrastructure.DB.Where("id = ?", roomID).First(&room).Error; err != nil {
		// Roomが見つからない場合、404エラーレスポンスを返す
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Room not found",
			"room_id": roomID,
		})
		return
	}

	// Roomが見つかった場合の成功レスポンス
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
