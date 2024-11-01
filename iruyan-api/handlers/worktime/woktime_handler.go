// handlers/worktime/worktime_handler.go
package worktime

import (
	"fmt"
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RecordEntry - WorkTimeテーブルに入室情報を記録
func RecordEntry(userID uint, roomID string) (*models.WorkTime, error) {
	// ユーザーが存在するか確認
	var user models.User
	if err := infrastructure.DB.First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// すでに入室しているか確認 (LeavingTimeがゼロのレコードをチェック)
	var activeEntry models.WorkTime
	if err := infrastructure.DB.
		Where("user_id = ? AND room_id = ? AND leaving_time IS NULL", userID, roomID).
		First(&activeEntry).Error; err == nil {
		return nil, fmt.Errorf("user is already in the room")
	} else if err != gorm.ErrRecordNotFound {
		return nil, err // その他のエラーが発生した場合は返す
	}

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
		return nil, err
	}

	return workTime, nil
}

// GetLatestEntry - 最新の入室記録を取得
func GetLatestEntry(userID uint, roomID string) (*models.WorkTime, error) {
	var workTime models.WorkTime
	if err := infrastructure.DB.
		Where("user_id = ? AND room_id = ? AND leaving_time IS NULL", userID, roomID).
		Order("entry_time desc").
		First(&workTime).Error; err != nil {
		return nil, err
	}
	return &workTime, nil
}

// GetLatestLogs - 最近N回の入室記録を取得
func GetLatestLogs(userID uint, limitN int) ([]models.WorkTime, error) {
	var workTimes []models.WorkTime
	if err := infrastructure.DB.
		Where("user_id = ?", userID).
		Order("entry_time desc").
		Limit(limitN).
		Find(&workTimes).Error; err != nil {
		return nil, err
	}
	return workTimes, nil
}

// GetLogForLastWeek - 1週間の作業記録を取得するメソッド
func GetLogForLastWeek(userID uint) ([]models.WorkTime, error) {
	var workTimes []models.WorkTime
	// 1週間前の日時を計算
	oneWeekAgo := time.Now().AddDate(0, 0, -6)
	if err := infrastructure.DB.
		Where("user_id = ? AND entry_time >= ?", userID, oneWeekAgo).
		Order("entry_time desc").
		Find(&workTimes).Error; err != nil {
		return nil, err
	}
	return workTimes, nil
}
