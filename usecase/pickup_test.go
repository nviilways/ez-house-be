package usecase_test

import (
	"errors"
	"testing"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	mocks "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/mocks/repository"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetPickupById(t *testing.T) {
	t.Run("should return selected pickup", func(t *testing.T) {
		id := uint(1)
		mockRepository := new(mocks.PickupRepository)
		usecase := usecase.NewPickupUsecase(&usecase.PickupUConfig{
			PickupRepository: mockRepository,
		})
		var expected *entity.Pickup
		mockRepository.On("GetPickupById", id).Return(expected, nil)

		result, _ := usecase.GetPickupById(id)

		assert.Equal(t, expected, result)
	})

	t.Run("should return error when failed", func(t *testing.T) {
		id := uint(1)
		mockRepository := new(mocks.PickupRepository)
		usecase := usecase.NewPickupUsecase(&usecase.PickupUConfig{
			PickupRepository: mockRepository,
		})
		mockRepository.On("GetPickupById", id).Return(nil, errors.New("error"))
		expected := errors.New("error")

		_, err := usecase.GetPickupById(id)

		assert.Error(t, expected, err)
	})
}

func TestGetPickupList(t *testing.T) {
	t.Run("should return pickup list", func(t *testing.T) {
		var pickup []*entity.Pickup
		mockRepository := new(mocks.PickupRepository)
		usecase := usecase.NewPickupUsecase(&usecase.PickupUConfig{
			PickupRepository: mockRepository,
		})
		mockRepository.On("GetPickupList").Return(pickup, nil)

		result, _ := usecase.GetPickupList()

		assert.Equal(t, pickup, result)
	})

	t.Run("should return error when failed", func(t *testing.T) {
		mockRepository := new(mocks.PickupRepository)
		usecase := usecase.NewPickupUsecase(&usecase.PickupUConfig{
			PickupRepository: mockRepository,
		})
		mockRepository.On("GetPickupList").Return(nil, errors.New("error"))
		expected := errors.New("error")

		_, err := usecase.GetPickupList()

		assert.Equal(t, expected, err)
	})
}

func TestRequestPickup(t *testing.T) {
	t.Run("should return requested pickup data", func(t *testing.T) {
		pickup := &entity.Pickup{
			UserID: 1,
			ReservationID: 1,
			PickupStatusID: 1,
		}
		mockRepository := new(mocks.PickupRepository)
		usecase := usecase.NewPickupUsecase(&usecase.PickupUConfig{
			PickupRepository: mockRepository,
		})
		mockRepository.On("RequestPickup", pickup).Return(pickup, nil)

		result, _ := usecase.RequestPickup(pickup)

		assert.Equal(t, pickup, result)
	})

	t.Run("should return error when failed to create request", func(t *testing.T) {
		pickup := &entity.Pickup{
			UserID: 1,
			ReservationID: 1,
			PickupStatusID: 1,
		}
		mockRepository := new(mocks.PickupRepository)
		usecase := usecase.NewPickupUsecase(&usecase.PickupUConfig{
			PickupRepository: mockRepository,
		})
		mockRepository.On("RequestPickup", pickup).Return(nil, errors.New("error"))
		expected := errors.New("error")

		_, err := usecase.RequestPickup(pickup)

		assert.Error(t, expected, err)
	})
}

func TestUpdateStatus(t *testing.T) {
	t.Run("should return updated status pickup data", func(t *testing.T) {
		pickup := &entity.Pickup{
			UserID: 1,
			ReservationID: 1,
			PickupStatusID: 1,
		}
		mockRepository := new(mocks.PickupRepository)
		usecase := usecase.NewPickupUsecase(&usecase.PickupUConfig{
			PickupRepository: mockRepository,
		})
		mockRepository.On("UpdateStatus", pickup.ID).Return(pickup, nil)
		
		result, _ := usecase.UpdateStatus(pickup.ID)

		assert.Equal(t, pickup, result)
	})

	t.Run("should return error when failed to update status", func(t *testing.T) {
		pickup := &entity.Pickup{
			UserID: 1,
			ReservationID: 1,
			PickupStatusID: 1,
		}
		mockRepository := new(mocks.PickupRepository)
		usecase := usecase.NewPickupUsecase(&usecase.PickupUConfig{
			PickupRepository: mockRepository,
		})
		mockRepository.On("UpdateStatus", pickup.ID).Return(nil, errors.New("error"))
		expected := errors.New("error")
		
		_, err := usecase.UpdateStatus(pickup.ID)

		assert.Error(t, expected, err)
	})
}