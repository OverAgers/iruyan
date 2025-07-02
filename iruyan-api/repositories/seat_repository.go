package repositories

import (
	"iruyan-api/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeatRepositoryInterface interface {
	CreateSeat(tx *gorm.DB, roomID uuid.UUID, seatNumber int) (*models.Seat, error)
	ExistsInRoom(roomID uuid.UUID, seatNumber int) (bool, error)
}

type seatRepository struct {
	DB *gorm.DB
}

func NewSeatRepository(db *gorm.DB) SeatRepositoryInterface {
	return &seatRepository{
		DB: db,
	}
}

// CreateSeat は指定された RoomID に紐づく Seat を1つ作成する
func (r *seatRepository) CreateSeat(tx *gorm.DB, roomID uuid.UUID, seatNumber int) (*models.Seat, error) {
	seat := models.Seat{
		ID:     uuid.New(),
		RoomID: roomID,
	}
	if err := tx.Create(&seat).Error; err != nil {
		return nil, err
	}
	return &seat, nil
}

func (r *seatRepository) ExistsInRoom(roomID uuid.UUID, seatNumber int) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Seat{}).
		Where("room_id = ? AND seat_number = ?", roomID, seatNumber).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
