package models

import (
	"github.com/google/uuid"
)

type Seat struct {
	ID         uuid.UUID `gorm:"type:uuid;;primaryKey"`
	RoomID     uuid.UUID `gorm:"type:uuid;not null"`
	SeatNumber int       `gorm:"not null"`
	Room       Room      `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE"`
}
