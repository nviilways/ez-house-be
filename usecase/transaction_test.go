package usecase_test

import (
	"errors"
	"testing"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	mocks "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/mocks/repository"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/usecase"
	"github.com/stretchr/testify/assert"
)

func TestTopUp(t *testing.T) {
	t.Run("should return nil when top up successful", func(t *testing.T) {
		userId := 1
		tx := &entity.Transaction{
			ID: 1,
			WalletID: 1,
			Balance: 50000,
		}
		mockRepository := new(mocks.TransactionRepository)
		usecase := usecase.NewTransactionUsecase(&usecase.TransactionUConfig{
			TransactionRepository: mockRepository,
		})
		mockRepository.On("TopUp", userId, tx).Return(tx, nil)

		result, _ := usecase.TopUp(userId, tx)

		assert.Equal(t, tx, result)		
	})

	t.Run("should return error when failed to top up", func(t *testing.T) {
		userId := 1
		tx := &entity.Transaction{
			ID: 1,
			WalletID: 1,
			Balance: 50000,
		}
		mockRepository := new(mocks.TransactionRepository)
		usecase := usecase.NewTransactionUsecase(&usecase.TransactionUConfig{
			TransactionRepository: mockRepository,
		})
		mockRepository.On("TopUp", userId, tx).Return(tx, errors.New("error"))
		expectedErr := errors.New("error")

		_, err := usecase.TopUp(userId, tx)

		assert.Error(t, expectedErr, err)
	})
}
