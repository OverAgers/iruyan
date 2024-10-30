package seat

import (
	"net/http"
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/responses"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ルームIDに基づくシート作成
func SeatCreateHandler(c *gin.Context) {
	roomIDParam := c.Param("room_id")
	
	// room_idをUUID型に変換してroomID変数に保存
	// roomID, err := uuid.Parse(roomIDParam)
	roomID, err := uuid.Parse(roomIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid roomID format",
		})
		return
	}

	// room_idが有効であるか（Room DBに存在するか）確認
	var room models.Room
	if err := room.FindByID(infrastructure.DB, roomIDParam); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "Room not found",
		})
		return
	}

	// シートオブジェクトを作成する
	seat, err := models.NewSeat(infrastructure.DB, roomID)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// DBにシートを作成する
	result := infrastructure.DB.Create(seat)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to save seat to database",
		})
		return
	}

	c.JSON(http.StatusOK, responses.SeatCreateResponse {
		Message: "seat create successful",
		Seat: responses.SeatInfo{
			RoomID:     seat.RoomID,
			SeatNumber: seat.SeatNumber,
		},
	})
}