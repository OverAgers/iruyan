package presenters

import (
	"iruyan-api/models"
	"iruyan-api/responses"
)

func ToRoomDetail(room *models.Room) responses.RoomDetail {
	seats := make([]responses.SeatDetail, len(room.Seats))
	for i, seat := range room.Seats {
		seats[i] = responses.SeatDetail{
			SeatID:     seat.ID.String(),
			RoomID:     seat.RoomID.String(),
			SeatNumber: seat.SeatNumber,
		}
	}
	return responses.RoomDetail{
		RoomID:   room.ID.String(),
		RoomName: room.Name,
		Seats:    seats,
	}
}
