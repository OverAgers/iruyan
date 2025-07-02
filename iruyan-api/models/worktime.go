package models

import (
	"time"

	"github.com/google/uuid"
)

type WorkTime struct {
	ID          int           `gorm:"primaryKey;autoIncrement"`
	UserID      uint          `gorm:"not null"`
	RoomID      uuid.UUID     `gorm:"type:uuid;not null"`
	Task        string        `gorm:"size:255;default:''"`
	EntryTime   time.Time     `gorm:"not null"`
	LeavingTime time.Time     `gorm:"default:null"`
	SeatNumber  int           `gorm:"default:0"`
	Duration    time.Duration `gorm:"default:0"`
	CreatedAt   time.Time
	User        User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Room        Room `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE"`
}
