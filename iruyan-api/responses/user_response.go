// user_response.go
package responses

import (
	"time"
)

// WorkTimeLog - 入室記録の情報を保持する型（DBの1行に該当）
type WorkTimeLog struct {
	EntryTime   time.Time		`json:"entry_time"`
	LeavingTime time.Time      	`json:"leaving_time"`		  
	Duration    time.Duration  	`json:"duration"`
}


// GetRecentWorktimeLogResponse - Room退室時のレスポンス
type GetRecentLogResponse struct {
	Message		string	`json:"message"`
	UserID		string	`json:"user_id"`
	WorkTimeLog	[]WorkTimeLog	`json:"worktime_log"`	
}


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