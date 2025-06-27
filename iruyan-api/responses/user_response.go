// Package responses provides response types for API endpoints related to user activity logs.
package responses

import (
	"time"
)

// WorkTimeLog holds a single record of entry and exit times with duration.
type WorkTimeLog struct {
	EntryTime   time.Time     `json:"entryTime"`
	LeavingTime time.Time     `json:"leavingTime"`
	Duration    time.Duration `json:"duration"`
}

// GetRecentLogResponse is the response returned when a user exits a room.
type GetRecentLogResponse struct {
	Message     string        `json:"message"`
	UserID      string        `json:"userId"`
	WorkTimeLog []WorkTimeLog `json:"workTimeLog"`
}

// DailyWorkLogResponse represents the work log summary for a single day.
type DailyWorkLogResponse struct {
	Date  string        `json:"date"`
	Hours time.Duration `json:"hours"`
}

// WorkLogForLastWeekResponse contains the daily work logs for the past week for a specific user.
type WorkLogForLastWeekResponse struct {
	UserID    uint                   `json:"userId"`
	DailyLogs []DailyWorkLogResponse `json:"dailyLogs"`
}
