// Package worktime provides functions for handling work time records such as
// recording entry times, retrieving recent logs, and calculating logs for the past week.
package worktime

import (
	"fmt"
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"log"
	"time"

	"github.com/google/uuid"
)

// RecordEntry - WorkTimeテーブルに入室情報を記録
func RecordEntry(userID uint, roomIDStr string, task string) (*models.WorkTime, error) {
	log.Printf("[INFO] 入室処理開始 (user_id=%d, room_id=%s)", userID, roomIDStr)

	// 文字列のUUIDをパース
	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		log.Printf("[ERROR] 無効なroomID形式: %s", roomIDStr)
		return nil, fmt.Errorf("invalid room ID format")
	}

	// 入室中かどうかを確認
	var existing models.WorkTime
	inRoom, err := existing.IsUserAlreadyInRoom(infrastructure.DB, userID, roomID)
	if err != nil {
		log.Printf("[ERROR] 入室確認に失敗 (user_id=%d, room_id=%s): %v", userID, roomID, err)
		return nil, err
	}
	if inRoom {
		log.Printf("[WARN] すでに入室中 (user_id=%d, room_id=%s)", userID, roomID)
		return nil, fmt.Errorf("user is already in the room")
	}

	// 入室レコードを作成
	entryTime := time.Now()
	workTime := &models.WorkTime{
		UserID:    userID,
		RoomID:    roomID,
		Task:      task,
		EntryTime: entryTime,
	}

	if err := infrastructure.DB.Create(workTime).Error; err != nil {
		log.Printf("[ERROR] WorkTime作成失敗 (user_id=%d, room_id=%s): %v", userID, roomID, err)
		return nil, err
	}

	log.Printf("[INFO] 入室記録完了 (user_id=%d, room_id=%s, entry_time=%s)", userID, roomID, entryTime.Format(time.RFC3339))
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
