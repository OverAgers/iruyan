package models

import (
	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string     `gorm:"size:255;not null"`
	Seats     []Seat     `gorm:"foreignKey:RoomID"`
	WorkTimes []WorkTime `gorm:"foreignKey:RoomID"`
}
