package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkTime struct {
	ID          int           `gorm:"primaryKey;autoIncrement"`
	UserID      uint          `gorm:"not null"`
	RoomID      uuid.UUID     `gorm:"type:uuid;not null"`
	Task        string        `gorm:"size:255;default:''"`
	EntryTime   time.Time     `gorm:"not null"`
	LeavingTime time.Time     `gorm:"default:null"`
	SeatNumber  int           `gorm:"default:0"`
	Duration    time.Duration `gorm:"default:0"`
	CreatedAt   time.Time
	User        User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Room        Room `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE"`
}

// IsSeatTakenInRoom - IDを元にルームの座席が使用されているか検索するメソッド
func (w *WorkTime) IsSeatTakenInRoom(db *gorm.DB, roomID uuid.UUID, seatNumber int) error {
	if err := db.Where("room_id = ? AND seat_number = ?", roomID, seatNumber).First(w).Error; err != nil {
		// 座席が見つからなかったとき（空いているとき）
		return nil
	}
	// 座席が見つかったとき
	return errors.New("Seat is already taken")
}

// IsUserAlreadyInRoom - 指定したユーザーが指定したルームに入室中か確認する
func (w *WorkTime) IsUserAlreadyInRoom(db *gorm.DB, userID uint, roomID uuid.UUID) (bool, error) {
	err := db.
		Where("user_id = ? AND room_id = ? AND leaving_time IS NULL", userID, roomID).
		First(w).Error

	if err == nil {
		return true, nil // 入室中
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil // 入室していない
	}
	return false, err // その他のエラー
}

// CreateWorkTime 新しい入室レコードを作成して保存する
func RecordEntry(db *gorm.DB, iruyanID string, roomID uuid.UUID, task string) (*WorkTime, error) {
	// ユーザーを検索
	var user User
	if err := user.FindByIruyanID(db, iruyanID); err != nil {
		return nil, err
	}

	workTime := &WorkTime{
		UserID:      user.ID,
		RoomID:      roomID,
		Task:        task,
		EntryTime:   time.Now(),
		LeavingTime: time.Time{}, // null相当
		Duration:    0,
		SeatNumber:  0,
	}

	if err := db.Create(workTime).Error; err != nil {
		return nil, err
	}

	return workTime, nil
}

// GetLatestEntry retrieves the latest WorkTime entry for a given user and room where LeavingTime is NULL.
func GetLatestEntry(db *gorm.DB, iruyanID string, roomID uuid.UUID) (*WorkTime, error) {
	// ユーザーを検索
	var user User
	if err := user.FindByIruyanID(db, iruyanID); err != nil {
		return nil, err
	}

	var workTime WorkTime
	err := db.
		Where("user_id = ? AND room_id = ? AND leaving_time IS NULL", user.ID, roomID).
		Order("entry_time DESC").
		First(&workTime).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return &workTime, nil
}

// GetLatestEntryByUser retrieves the latest WorkTime entry for a given user across all rooms where LeavingTime is NULL.
func GetLatestEntryByUser(db *gorm.DB, iruyanID string) (*WorkTime, error) {
	// ユーザーを検索
	var user User
	if err := user.FindByIruyanID(db, iruyanID); err != nil {
		return nil, err
	}

	var workTime WorkTime
	err := db.
		Where("user_id = ? AND leaving_time IS NULL", user.ID).
		Order("entry_time DESC").
		First(&workTime).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return &workTime, nil
}

// GetLatestLogsByUser retrieves the latest 5 WorkTime logs for a given user, regardless of room.
func GetLatestEntriesByUser(db *gorm.DB, iruyanID string, limit int) ([]WorkTime, error) {
	// ユーザーを検索
	var user User
	if err := user.FindByIruyanID(db, iruyanID); err != nil {
		return nil, err
	}

	var logs []WorkTime
	if err := db.
		Where("user_id = ?", user.ID).
		Order("entry_time DESC").
		Limit(limit).
		Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

// GetLatestLogs returns the WorkTime records for the past 5 days for a given IruyanID.
func GetLogForLastWeek(db *gorm.DB, iruyanID string) ([]WorkTime, error) {
	// ユーザーを検索
	var user User
	if err := user.FindByIruyanID(db, iruyanID); err != nil {
		return nil, err
	}

	var workTimes []WorkTime

	fiveDaysAgo := time.Now().AddDate(0, 0, -5)

	err := db.
		Where("user_id = ? AND entry_time >= ?", user.ID, fiveDaysAgo).
		Order("entry_time DESC").
		Find(&workTimes).Error

	if err != nil {
		return nil, err
	}

	return workTimes, nil // データが0件でも空スライスで返る
}
