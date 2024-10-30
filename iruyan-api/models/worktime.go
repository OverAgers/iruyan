package models

import (
	"time"

	"github.com/google/uuid"
)

type WorkTime struct {
	ID          int           `gorm:"primaryKey;autoIncrement"`
	UserID      uint          `gorm:"type:uuid;not null"`
	RoomID      uuid.UUID     `gorm:"type:uuid;not null"`
	EntryTime   time.Time     `gorm:"not null"`
	LeavingTime time.Time     `gorm:"default:null"` // 空の値をデフォルトに
	SeatNumber  int           `gorm:"default:0"`    // デフォルト値を0に設定
	Duration    time.Duration `gorm:"default:0"`    // デフォルト値を0に設定
	CreatedAt   time.Time
	User        User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Room        Room `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE"`
}
