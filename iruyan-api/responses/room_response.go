package responses

import (
	"time"

	"github.com/google/uuid"
)

// RoomCreateResponse Room作成成功時のレスポンス
type RoomCreateResponse struct {
	Message  string    `json:"message"`
	RoomID   uuid.UUID `json:"room_id"`
	RoomName string    `json:"room_name"`
}

// RoomDetailResponse Room詳細表示のレスポンス
type RoomDetailResponse struct {
	Message string `json:"message"`
	RoomID  string `json:"room_id"`
}

// RoomActionResponse Room入室・退室の成功レスポンス
type RoomActionResponse struct {
	Message string `json:"message"`
	RoomID  string `json:"room_id"`
}

// EnterRoomResponse - Room入室時のレスポンス
type EnterRoomResponse struct {
	Message   string    `json:"message"`
	RoomID    string    `json:"room_id"`
	RoomName  string    `json:"room_name"`
	UserID    string    `json:"user_id"`
	EntryTime time.Time `json:"entry_time"`
}

// LeaveRoomResponse - Room退室時のレスポンス
type LeaveRoomResponse struct {
	Message     string        `json:"message"`
	RoomID      string        `json:"room_id"`
	RoomName    string        `json:"room_name"`
	UserID      string        `json:"user_id"`
	EntryTime   time.Time     `json:"entry_time"`
	LeavingTime time.Time     `json:"leaving_time"`
	Duration    time.Duration `json:"duration"` // 滞在時間
}
