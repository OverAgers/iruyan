package usecases

import (
	"iruyan-api/repositories"
)

type SeatUsecase interface {
}

type seatUsecase struct {
	Repo repositories.SeatRepositoryInterface
}
