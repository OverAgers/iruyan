package seat

import (
	"errors"
	"net/http"
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/responses"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ルームIDに基づくシート作成
func SeatCreateHandler(c *gin.Context) {
	roomIDParam := c.Param("roomId")

	// room_idをUUID型に変換してroomID変数に保存
	roomID, err := uuid.Parse(roomIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid roomID format",
		})
		return
	}

	// CreateSeat関数を呼び出しシートを作成
	seat, err := CreateSeat(infrastructure.DB, roomID)
	if err != nil {
		if err.Error() == "room not found" {
			c.JSON(http.StatusNotFound, responses.ErrorResponse{
				Message: err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Message: "Failed to create seat",
			})
		}
		return
	}

	c.JSON(http.StatusOK, responses.SeatCreateResponse{
		Message: "seat create successful",
		Seat: responses.SeatInfo{
			RoomID:     seat.RoomID,
			SeatNumber: seat.SeatNumber,
		},
	})
}


func CreateSeat(tx *gorm.DB, roomID uuid.UUID) (*models.Seat, error) {
	// room_idが有効であるか（Room DBに存在するか）確認
	var room models.Room
	if err := room.FindByID(tx, roomID.String()); err != nil {
		return nil, errors.New("room not found")
	}

	// シートオブジェクトを作成する
	seat, err := models.NewSeat(tx, roomID)
	if err != nil {
		return nil, err
	}

	// DBにシートを作成する
	result := tx.Create(seat)
	if result.Error != nil {
		return nil, result.Error
	}

	return seat, nil
}