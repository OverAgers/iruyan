package models

import (
	"iruyan-api/pkg/errdefs"

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
		ID:         uuid.New(),
		RoomID:     roomID,
		SeatNumber: seatNumber,
	}, nil
}

// IsSeatExistsInRoom - 部屋に指定された座席番号が存在するか確認するメソッド
func (s *Seat) IsSeatExistsInRoom(db *gorm.DB, roomID uuid.UUID, seatNumber int) error {
	if err := db.Where("room_id = ? AND seat_number = ?", roomID, seatNumber).First(s).Error; err != nil {
		// 座席が存在しないとき
		return errdefs.ErrSeatNotFound
	}
	// 座席が存在するとき
	return nil
}
