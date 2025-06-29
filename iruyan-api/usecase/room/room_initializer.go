package room

import (
	"errors"
	"fmt"

	"iruyan-api/handlers/seat"
	"iruyan-api/infrastructure"
	"iruyan-api/models"

	"gorm.io/gorm"
)

func CreateRoomWithSeats(name string, seatCount int) error {
	db := infrastructure.DB

	// 既に同名のルームが存在するか確認
	var existingRoom models.Room
	if err := db.Where("name = ?", name).First(&existingRoom).Error; err == nil {
		// ルームが存在する場合はスキップ
		fmt.Printf("Room '%s' already exists. Skipping creation.\n", name)
		return nil
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// 検索中に別のエラーが発生した場合は終了
		return fmt.Errorf("failed to check existing room: %w", err)
	}

	// トランザクション開始
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
