package models

import (
	"time"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

// IsSeatTakenInRoom - IDを元にルームの座席が使用されているか検索するメソッド
func (w *WorkTime) IsSeatTakenInRoom(db *gorm.DB, roomID uuid.UUID, seatNumber int) error {
	if err := db.Where("room_id = ? AND seat_number = ?", roomID, seatNumber).First(w).Error; err != nil {
		// 座席が見つからなかったとき（空いているとき）
		return nil
	}
	// 座席が見つかったとき
	return errors.New("Seat is already taken")
}