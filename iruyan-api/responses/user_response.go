// user_response.go

package responses

import (
	"time"
)

// 1日分の作業ログレスポンス
type DailyWorkLogResponse struct {
	Date  string `json:"date"`
	Hours time.Duration   `json:"hours"`
}

// 1週間分の作業ログレスポンス
type WorkLogForLastWeekResponse struct {
	UserID      uint                   `json:"user_id"`
	DailyLogs []DailyWorkLogResponse `json:"daily_logs"`
}