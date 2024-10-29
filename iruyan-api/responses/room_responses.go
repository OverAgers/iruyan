package responses

import "github.com/google/uuid"

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
