package routes

import (
	handlers "iruyan-api/handlers/room"
	"iruyan-api/infrastructure"
	"iruyan-api/repositories"
	usecases "iruyan-api/usecases/room"

	"github.com/gin-gonic/gin"
)

func RegisterRoomRoutes(router *gin.RouterGroup) {
	// --- 依存性の注入 ---
	db := infrastructure.DB

	userRepo := repositories.NewUserRepository(db)
	roomRepo := repositories.NewRoomRepository(db)
	seatRepo := repositories.NewSeatRepository(db)
	workTimeRepo := repositories.NewWorkTimeRepository(db)

	roomUsecase := usecases.NewRoomUsecase(userRepo, roomRepo, seatRepo, workTimeRepo, db)
	roomHandler := handlers.NewRoomHandler(roomUsecase)

	// Roomの一覧取得
	router.GET("/rooms", roomHandler.GetRoomsHandler)

	// Room作成
	router.POST("/rooms", roomHandler.CreateRoomHandler)

	// Room表示
	router.GET("/rooms/:roomId", roomHandler.GetRoomHandler)

	// Room入室
	router.POST("/rooms/:roomId/enter/:iruyanId", roomHandler.EnterRoomHandler)

	// Room退室
	router.POST("/rooms/:roomId/leave", roomHandler.LeaveRoomHandler)

	// 着席
	router.PUT("/rooms/:roomId/seats/:seatNumber/take", roomHandler.TakeSeatHandler)

	// 離席
	router.PUT("/rooms/:roomId/seats/:seatNumber/leave", roomHandler.LeaveSeatHandler)

	// 入室している部屋の着席状況一覧取得
	router.GET("/rooms/:roomId/seats/status", roomHandler.GetSeatedUsersInRoomHandler)
}
