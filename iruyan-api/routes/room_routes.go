package routes

import (
	"iruyan-api/handlers/room"

	"github.com/gin-gonic/gin"
)

func RegisterRoomRoutes(router *gin.Engine) {
	// Room作成
	router.POST("/rooms", room.CreateRoomHandler)

	// Room表示
	router.GET("/room/:room_id", room.GetRoomHandler)

	// Room入室
	router.POST("/room/:room_id/enter", room.EnterRoomHandler)

	// Room退室
	router.DELETE("/room/:room_id/leave", room.LeaveRoomHandler)

	// 着席
	router.PATCH("/room/:room_id/:seat_id", room.TakeSeatHandler)

	// 離席
	router.PATCH("/room/:room_id/:seat_id", room.LeaveSeatHandler)
}
