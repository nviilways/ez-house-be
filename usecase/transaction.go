package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/repository"
)

type TransactionUsecase interface {
	GetTransactionByWalletId(uint) ([]*entity.Transaction, error)
	TopUp(int, *entity.Transaction) (*entity.Transaction, error)
}

type transactionUsecaseImpl struct {
	transactionRepository repository.TransactionRepository
}

type TransactionUConfig struct {
	TransactionRepository repository.TransactionRepository
}

func NewTransactionUsecase(cfg *TransactionUConfig) TransactionUsecase {
	return &transactionUsecaseImpl {
		transactionRepository: cfg.TransactionRepository,
	}
}

func (t *transactionUsecaseImpl) GetTransactionByWalletId(id uint) ([]*entity.Transaction, error) {
	return t.transactionRepository.GetTransactionByWalletId(id)
}

func (t *transactionUsecaseImpl) TopUp(target int,tr *entity.Transaction) (*entity.Transaction, error) {
	return t.transactionRepository.TopUp(target, tr)
}