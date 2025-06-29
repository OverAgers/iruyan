// Package responses provides response types for API endpoints related to user activity logs.
package responses

import (
	"time"
)

type User struct {
	Email    string `json:"email"`
	IruyanID string `json:"iruyanId"`
	Name     string `json:"name"`
}

type WorkTimeLog struct {
	EntryTime   time.Time     `json:"entryTime"`
	LeavingTime time.Time     `json:"leavingTime"`
	Duration    time.Duration `json:"duration"`
}

type GetRecentLogResponse struct {
	Message     string        `json:"message"`
	IruyanID    string        `json:"iruyanId"`
	WorkTimeLog []WorkTimeLog `json:"workTimeLog"`
}

type DailyWorkLogResponse struct {
	Date  string        `json:"date"`
	Hours time.Duration `json:"hours"`
}

type WorkLogForLastWeekResponse struct {
	IruyanID  string                 `json:"iruyanId"`
	DailyLogs []DailyWorkLogResponse `json:"dailyLogs"`
}

// --- Swagger用構造体 ---

type WorkTimeLogSwagger struct {
	EntryTime   time.Time `json:"entryTime"`
	LeavingTime time.Time `json:"leavingTime"`
	Duration    string    `json:"duration"`
}

type GetRecentLogResponseSwagger struct {
	Message     string               `json:"message"`
	UserID      string               `json:"userId"`
	WorkTimeLog []WorkTimeLogSwagger `json:"workTimeLog"`
}

type DailyWorkLogResponseSwagger struct {
	Date  string `json:"date"`
	Hours string `json:"hours"`
}

type WorkLogForLastWeekResponseSwagger struct {
	UserID    uint                          `json:"userId"`
	DailyLogs []DailyWorkLogResponseSwagger `json:"dailyLogs"`
}
