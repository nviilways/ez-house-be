package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/repository"
)

type ReservationUsecase interface {
	AddReservation(*entity.Reservation) (*entity.Reservation, error)
}

type reservationUsecaseImpl struct {
	reservationRepository repository.ReservationRepository
}

type ReservationUConfig struct {
	ReservationRepository repository.ReservationRepository
}

func NewReservationUsecase(cfg *ReservationUConfig) ReservationUsecase {
	return &reservationUsecaseImpl{
		reservationRepository: cfg.ReservationRepository,
	}
}

func (r *reservationUsecaseImpl) AddReservation(res *entity.Reservation) (*entity.Reservation, error) {
	return r.reservationRepository.AddReservation(res)
}
