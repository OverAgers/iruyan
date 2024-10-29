package models

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Room構造体
type Room struct {
	ID        uuid.UUID  `gorm:"type:uuid;;primaryKey"`
	Name      string     `gorm:"size:255;not null"`
	Seats     []Seat     `gorm:"foreignKey:RoomID"`
	WorkTimes []WorkTime `gorm:"foreignKey:RoomID"`
}

// 正規表現パターン：大文字・小文字の英数字および基本記号のみ
var roomNamePattern = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]+$`)

// ValidateName: Room名のバリデーション
func (r *Room) ValidateName() error {
	if r.Name == "" {
		return errors.New("room name is required")
	}
	if !roomNamePattern.MatchString(r.Name) {
		return errors.New("room name must contain only lowercase letters and numbers, with no special characters")
	}
	return nil
}

// BeforeCreate: GORMのフックで、新しいIDを自動的に生成
func (r *Room) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}
