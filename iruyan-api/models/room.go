// Package models contains the data models used throughout the application.
package models

import (
	"iruyan-api/pkg/errdefs"
	"regexp"

	"github.com/google/uuid"
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
		return errdefs.ErrRoomNameRequired
	}
	if !roomNamePattern.MatchString(r.Name) {
		return errdefs.ErrRoomNameInvalidFormat
	}
	return nil
}
