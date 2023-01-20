package dto

import "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"

type TopUp struct {
	Amount int `json:"amount" binding:"min=50000,max=10000000"`
}

func (t *TopUp) ToTransaction() *entity.Transaction {
	return &entity.Transaction{
		TransactionTypeID: 1,
		Balance: t.Amount,
	}
}