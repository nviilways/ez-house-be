package handler

import "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/usecase"

type Handler struct {
	userUsecase usecase.UserUsecase
	transactionUsecase usecase.TransactionUsecase
	adminUsecase usecase.AdminUsecase
	houseUsecase usecase.HouseUsecase
	reservationUsecase usecase.ReservationUsecase
}

type Config struct {
	UserUsecase usecase.UserUsecase
	TransactionUsecase usecase.TransactionUsecase
	AdminUsecase usecase.AdminUsecase
	HouseUsecase usecase.HouseUsecase
	ReservationUsecase usecase.ReservationUsecase
}

func New(cfg *Config) *Handler {
	return &Handler{
		userUsecase: cfg.UserUsecase,
		transactionUsecase: cfg.TransactionUsecase,
		adminUsecase: cfg.AdminUsecase,
		houseUsecase: cfg.HouseUsecase,
		reservationUsecase: cfg.ReservationUsecase,
	}
}