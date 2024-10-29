package models

import (
	"github.com/google/uuid"
)

type Seat struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	RoomID     uuid.UUID `gorm:"type:uuid;not null"`
	SeatNumber int       `gorm:"not null"`
	Room       Room      `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE"`
}
