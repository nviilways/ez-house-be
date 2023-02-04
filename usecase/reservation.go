package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/repository"
)

type ReservationUsecase interface {
	GetReservationListByUserId(uint, *dto.Pagination) ([]*entity.Reservation, int, error)
	GetReservationById(uint) (*entity.Reservation, error)
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

func (r *reservationUsecaseImpl) GetReservationListByUserId(id uint, pagination *dto.Pagination) ([]*entity.Reservation, int, error) {
	return r.reservationRepository.GetReservationListByUserId(id, pagination)
}

func (r *reservationUsecaseImpl) GetReservationById(id uint) (*entity.Reservation, error) {
	return r.reservationRepository.GetReservationById(id)
}
