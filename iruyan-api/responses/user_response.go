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
