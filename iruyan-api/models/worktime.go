package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

// IsSeatTakenInRoom - IDを元にルームの座席が使用されているか検索するメソッド
func (w *WorkTime) IsSeatTakenInRoom(db *gorm.DB, roomID uuid.UUID, seatNumber int) error {
	if err := db.Where("room_id = ? AND seat_number = ?", roomID, seatNumber).First(w).Error; err != nil {
		// 座席が見つからなかったとき（空いているとき）
		return nil
	}
	// 座席が見つかったとき
	return errors.New("Seat is already taken")
}

// IsUserAlreadyInRoom - 指定したユーザーが指定したルームに入室中か確認する
func (w *WorkTime) IsUserAlreadyInRoom(db *gorm.DB, userID uint, roomID uuid.UUID) (bool, error) {
	err := db.
		Where("user_id = ? AND room_id = ? AND leaving_time IS NULL", userID, roomID).
		First(w).Error

	if err == nil {
		return true, nil // 入室中
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil // 入室していない
	}
	return false, err // その他のエラー
}
