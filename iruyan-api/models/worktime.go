package models

import (
	"time"

	"github.com/google/uuid"
)

type WorkTime struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	RoomID      uuid.UUID `gorm:"type:uuid;not null"`
	EntryTime   time.Time
	LeavingTime time.Time
	SeatNumber  int
	Duration    time.Duration
	CreatedAt   time.Time
	User        User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Room        Room `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE"`
}
