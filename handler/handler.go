package handler

import "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/usecase"

type Handler struct {
	userUsecase usecase.UserUsecase
	transactionUsecase usecase.TransactionUsecase
}

type Config struct {
	UserUsecase usecase.UserUsecase
	TransactionUsecase usecase.TransactionUsecase
}

func New(cfg *Config) *Handler {
	return &Handler{
		userUsecase: cfg.UserUsecase,
		transactionUsecase: cfg.TransactionUsecase,
	}
}