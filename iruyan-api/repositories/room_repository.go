package repositories

import (
	"errors"
	"fmt"
	"iruyan-api/models"
	"iruyan-api/pkg/errdefs"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoomRepositoryInterface interface {
	FindByID(roomID uuid.UUID) (*models.Room, error)
	CreateWithTx(tx *gorm.DB, room *models.Room) error
	GetRoomsWithSeats() ([]*models.Room, error)
	FindByIDWithSeats(roomID uuid.UUID) (*models.Room, error)
	ExistsByName(name string) (bool, error)
}

type roomRepository struct {
	DB *gorm.DB
}

func (r *roomRepository) CreateWithTx(tx *gorm.DB, room *models.Room) error {
	if err := tx.Create(room).Error; err != nil {
		return fmt.Errorf("failed to create room: %w", err)
	}
	return nil
}

// FindByID finds a room by its ID using the provided database instance.
func (r *roomRepository) FindByID(roomID uuid.UUID) (*models.Room, error) {
	var room models.Room
	if err := r.DB.Where("id = ?", roomID).First(&room).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("room not found")
		}
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) GetRoomsWithSeats() ([]*models.Room, error) {
	var rooms []*models.Room
	if err := r.DB.Preload("Seats").Find(&rooms).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve rooms: %w", err)
	}
	return rooms, nil
}

func (r *roomRepository) FindByIDWithSeats(roomID uuid.UUID) (*models.Room, error) {
	var room models.Room
	if err := r.DB.Preload("Seats").Where("id = ?", roomID).First(&room).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errdefs.ErrRoomNotFound
		}
		return nil, fmt.Errorf("failed to retrieve room: %w", err)
	}
	return &room, nil
}

func (r *roomRepository) ExistsByName(name string) (bool, error) {
	var count int64
	if err := r.DB.Model(&models.Room{}).Where("name = ?", name).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
