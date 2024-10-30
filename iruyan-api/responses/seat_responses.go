package responses

import "github.com/google/uuid"

// SeatCreateResponse Seat作成成功時のレスポンス
type SeatCreateResponse struct {
	Message  	string    	`json:"message"`
	Seat		SeatInfo	`json: "seat"`	 
}

// SeatInfo シート情報のレスポンス
type SeatInfo struct {
	RoomID		uuid.UUID	`json:"room_id"`
	SeatNumber	int 		`json:"seat_number"`
}
