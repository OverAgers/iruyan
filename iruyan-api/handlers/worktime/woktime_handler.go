// handlers/worktime/worktime_handler.go
package worktime

import (
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"time"

	"github.com/google/uuid"
)

// RecordEntry - WorkTimeテーブルに入室情報を記録
func RecordEntry(userID uint, roomID string) (*models.WorkTime, error) {
	// 入室時刻を現在時刻として取得
	entryTime := time.Now()

	// WorkTimeレコードの作成
	workTime := &models.WorkTime{
		UserID:    userID,
		RoomID:    uuid.MustParse(roomID),
		EntryTime: entryTime,
	}

	// WorkTimeをデータベースに保存
	if err := infrastructure.DB.Create(workTime).Error; err != nil {
		return nil, err // エラーが発生した場合はnilとエラーを返す
	}

	return workTime, nil // 成功時にworkTimeレコードとnil（エラーなし）を返す
}
