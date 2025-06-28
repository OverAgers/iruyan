package routes

import (
	"iruyan-api/handlers/seat"

	"github.com/gin-gonic/gin"
)

func RegisterSeatRoutes(router *gin.Engine) {
	// シート作成
	router.POST("/rooms/:roomId/seats", seat.SeatCreateHandler)
}
