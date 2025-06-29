package room

import (
	"fmt"

	"iruyan-api/handlers/seat"
	"iruyan-api/infrastructure"
	"iruyan-api/models"
)

func CreateRoomWithSeats(name string, seatCount int) error {
	db := infrastructure.DB
	tx := db.Begin()

	room := models.Room{Name: name}

	if err := room.ValidateName(); err != nil {
		return fmt.Errorf("room name invalid: %w", err)
	}

	if err := tx.Create(&room).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create room: %w", err)
	}

	for i := 0; i < seatCount; i++ {
		if _, err := seat.CreateSeat(tx, room.ID); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create seat: %w", err)
		}
	}

	return tx.Commit().Error
}
