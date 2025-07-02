package responses

import (
	"iruyan-api/models"
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
	Message   string    `json:"message"`
	RoomID    string    `json:"roomId"`
	RoomName  string    `json:"roomName"`
	IruyanID  string    `json:"iruyanId"`
	EntryTime time.Time `json:"entryTime"`
	Task      string    `json:"task"`
}

// LeaveRoomResponse - Room退室時のレスポンス
type LeaveRoomResponse struct {
	Message     string        `json:"message"`
	RoomID      string        `json:"roomId"`
	RoomName    string        `json:"roomName"`
	IruyanID    string        `json:"iruyanId"`
	EntryTime   time.Time     `json:"entryTime"`
	LeavingTime time.Time     `json:"leavingTime"`
	Duration    time.Duration `json:"duration"` // 滞在時間
}

type LeaveRoomResponseSwagger struct {
	Message     string `json:"message"`
	RoomID      string `json:"roomId"`
	RoomName    string `json:"roomName"`
	IruyanID    string `json:"iruyanId"`
	EntryTime   string `json:"entryTime"`   // ISO8601表記想定
	LeavingTime string `json:"leavingTime"` // ISO8601表記想定
	Duration    string `json:"duration"`    // e.g., "1h2m3s"
}

type SeatStatusResponse struct {
	SeatNumber int    `json:"seat_number"`
	IruyanID   string `json:"iruyan_id"`
	UserName   string `json:"user_name"`
	Email      string `json:"email"`
	AvatarUrl  string `json:"avatar_url,omitempty"`
	Task       string `json:"task,omitempty"`
	Note       string `json:"note,omitempty"`
	Status     string `json:"status"`
	StartTime  int64  `json:"start_time"`
}

func ConvertToRoomDetail(room *models.Room) RoomDetail {
	seats := make([]SeatDetail, len(room.Seats))
	for i, seat := range room.Seats {
		seats[i] = SeatDetail{
			SeatID:     seat.ID.String(),
			RoomID:     seat.RoomID.String(),
			SeatNumber: seat.SeatNumber,
		}
	}
	return RoomDetail{
		RoomID:   room.ID.String(),
		RoomName: room.Name,
		Seats:    seats,
	}
}
