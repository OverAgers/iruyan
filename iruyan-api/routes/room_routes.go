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
	router.GET("/rooms/:roomId", room.GetRoomHandler)

	// Room入室
	router.POST("/rooms/:roomId/enter/:iruyanId", room.EnterRoomHandler)

	// Room退室
	router.POST("/rooms/:roomId/leave", room.LeaveRoomHandler)

	// 着席
	router.PUT("/rooms/:roomId/seat/:seatNumber/take", room.TakeSeatHandler)

	// 離席
	router.PUT("/rooms/:roomId/seat/:seatNumber/leave", room.LeaveSeatHandler)

	// 入室している部屋の着席状況一覧取得
	router.GET("/rooms/:roomId/seats/status", room.GetSeatedUsersInRoomHandler)
}
