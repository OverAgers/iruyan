package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Seat struct {
	ID         uuid.UUID `gorm:"type:uuid;;primaryKey"`
	RoomID     uuid.UUID `gorm:"type:uuid;not null"`
	SeatNumber int       `gorm:"not null"`
	Room       Room      `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE"`
}

// NewSeat: Seat構造体のコンストラクタ関数
func NewSeat(db *gorm.DB, roomID uuid.UUID) (*Seat, error) {
	var maxSeatNumber *int

	// 指定されたRoomIDの最大SeatNumberを取得する
	err := db.Model(&Seat{}).Where("room_id = ?", roomID).Select("max(seat_number)").Scan(&maxSeatNumber).Error
	if err != nil {
		return nil, err
	}

	// maxSeatNumberがNULLの場合は1を割り振る
	seatNumber := 1
	if maxSeatNumber != nil {
		seatNumber = *maxSeatNumber + 1
	}

	// 新しいSeatインスタンスを生成する
	return &Seat{
		ID			: uuid.New(),
		RoomID		: roomID,
		SeatNumber	: seatNumber,
	}, nil	
}