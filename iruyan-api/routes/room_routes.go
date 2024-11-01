package routes

import (
	"iruyan-api/handlers/room"

	"github.com/gin-gonic/gin"
)

func RegisterRoomRoutes(router *gin.Engine) {
	// Roomの一覧取得
	router.GET("/rooms", room.GetRoomsHandler)

	// Room作成
	router.POST("/rooms", room.CreateRoomHandler)

	// Room表示
	router.GET("/rooms/:room_id", room.GetRoomHandler)

	// Room入室
	router.POST("/rooms/:room_id/enter", room.EnterRoomHandler)

	// Room退室
	router.POST("/rooms/:room_id/leave", room.LeaveRoomHandler)

	// 着席
	router.PATCH("/rooms/:room_id/:seat_number/take", room.TakeSeatHandler)

	// 離席
	router.PATCH("/rooms/:room_id/:seat_number/leave", room.LeaveSeatHandler)
}
