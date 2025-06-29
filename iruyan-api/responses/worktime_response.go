package responses

import (
	"iruyan-api/models"
	"time"

	"github.com/google/uuid"
)

// WorkTimeResponse represents the response for a work time record.
// swagger:model
type WorkTimeResponse struct {
	IruyanID    string        `json:"iruyanId"`
	RoomID      uuid.UUID     `json:"roomId"`
	Task        string        `json:"task"`
	EntryTime   time.Time     `json:"entryTime"`             // RFC3339
	LeavingTime time.Time     `json:"leavingTime,omitempty"` // RFC3339 or nil
	Duration    time.Duration `json:"duration"`              // formatted like "1h23m"
}

type WorkTimeResponseSwagger struct {
	IruyanID    string    `json:"iruyanId"`
	RoomID      uuid.UUID `json:"roomId"`
	Task        string    `json:"task"`
	EntryTime   time.Time `json:"entryTime"`             // RFC3339
	LeavingTime time.Time `json:"leavingTime,omitempty"` // RFC3339 or nil
	Duration    string    `json:"duration"`              // formatted like "1h23m"
}

type WorkTimeListResponse struct {
	WorkTimes []models.WorkTime `json:"work_times"`
}

type WorkTimeListResponseSwagger struct {
	WorkTimes []WorkTimeResponseSwagger `json:"work_times"`
}
