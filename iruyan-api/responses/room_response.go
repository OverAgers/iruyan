package responses

import (
	"time"

	"github.com/google/uuid"
)

// RoomListResponse ルーム一覧取得時のレスポンス
type RoomListResponse struct {
	Message string       `json:"message"`
	Rooms   []RoomDetail `json:"rooms"`
}

// RoomDetail ルーム情報の詳細
type RoomDetail struct {
	RoomID   string       `json:"roomId"`
	RoomName string       `json:"roomName"`
	Seats    []SeatDetail `json:"seats"`
}

// SeatDetail シート情報の詳細
type SeatDetail struct {
	SeatID     string `json:"seatId"`
	RoomID     string `json:"roomId"`
	SeatNumber int    `json:"seatNumber"`
}

// RoomCreateResponse Room作成成功時のレスポンス
type RoomCreateResponse struct {
	Message  string    `json:"message"`
	RoomID   uuid.UUID `json:"roomId"`
	RoomName string    `json:"roomName"`
}

// RoomResponse 単一のルーム詳細を含むレスポンス
type RoomResponse struct {
	Message string     `json:"message"`
	RoomID  string     `json:"roomId"`
	Room    RoomDetail `json:"room"`
}

// RoomDetailResponse Room詳細表示のレスポンス
type RoomDetailResponse struct {
	Message string `json:"message"`
	RoomID  string `json:"roomId"`
}

// RoomActionResponse Room入室・退室の成功レスポンス
type RoomActionResponse struct {
	Message string `json:"message"`
	RoomID  string `json:"roomId"`
}

// EnterRoomResponse - Room入室時のレスポンス
type EnterRoomResponse struct {
	Message		string		`json:"message"`
	RoomID		string		`json:"roomId"`
	RoomName	string		`json:"roomName"`
	UserID		string		`json:"userId"`
	EntryTime	time.Time	`json:"entryTime"`
	Task		string		`json:"task"`
}

// LeaveRoomResponse - Room退室時のレスポンス
type LeaveRoomResponse struct {
	Message     string        `json:"message"`
	RoomID      string        `json:"roomId"`
	RoomName    string        `json:"roomName"`
	UserID      string        `json:"userId"`
	EntryTime   time.Time     `json:"entryTime"`
	LeavingTime time.Time     `json:"leavingTime"`
	Duration    time.Duration `json:"duration"` // 滞在時間
}
