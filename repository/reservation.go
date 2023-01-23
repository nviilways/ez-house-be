package repository

import (
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"gorm.io/gorm"
)

type ReservationRepository interface {
	AddReservation(*entity.Reservation) (*entity.Reservation, error)
}

type reservationRepositoryImpl struct {
	db *gorm.DB
}

type ReservationRConfig struct {
	DB *gorm.DB
}

func NewReservationRepository(cfg *ReservationRConfig) ReservationRepository {
	return &reservationRepositoryImpl{
		db: cfg.DB,
	}
}

func (r *reservationRepositoryImpl) AddReservation(res *entity.Reservation) (*entity.Reservation, error) {
	err := r.db.Not("check_in_date = ? AND check_out_date = ? AND house_id = ?", res.CheckInDate, res.CheckOutDate, res.HouseId).Create(&res).Error
	if err != nil {
		if err == gorm.ErrRegistered {
			return nil, errs.ErrDuplicateEntry
		}
		return nil, err
	}

	return res, nil
}