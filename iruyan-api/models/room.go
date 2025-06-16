// Package models contains the data models used throughout the application.
package models

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Room represents a room entity with seats and associated work times.
type Room struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Name      string     `gorm:"size:255;not null"`
	Seats     []Seat     `gorm:"foreignKey:RoomID"`
	WorkTimes []WorkTime `gorm:"foreignKey:RoomID"`
}

// roomNamePattern is a regular expression for validating room names.
var roomNamePattern = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]+$`)

// ValidateName validates the room name using a predefined pattern.
func (r *Room) ValidateName() error {
	if r.Name == "" {
		return errors.New("room name is required")
	}
	if !roomNamePattern.MatchString(r.Name) {
		return errors.New("room name must contain only lowercase letters and numbers, with no special characters")
	}
	return nil
}

// FindByID finds a room by its ID using the provided database instance.
func (r *Room) FindByID(db *gorm.DB, roomID string) error {
	if err := db.Where("id = ?", roomID).First(r).Error; err != nil {
		return errors.New("room not found")
	}
	return nil
}

// BeforeCreate sets a new UUID for the room before saving to the database.
func (r *Room) BeforeCreate(_ *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}
