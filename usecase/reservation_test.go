package usecase_test

import (
	"errors"
	"testing"
	"time"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	mocks "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/mocks/repository"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/usecase"
	"github.com/stretchr/testify/assert"
)

func TestAddReservation(t *testing.T) {
	t.Run("should return created reservation when successful", func(t *testing.T) {
		reserve := &entity.Reservation{
			CheckInDate: time.Now(),
			CheckOutDate: time.Now(),
			HouseId: 2,
		}
		mockRepository := new(mocks.ReservationRepository)
		usecase := usecase.NewReservationUsecase(&usecase.ReservationUConfig{
			ReservationRepository: mockRepository,
		})
		mockRepository.On("AddReservation", reserve).Return(reserve, nil)

		result, _ := usecase.AddReservation(reserve)
		
		assert.Equal(t, reserve, result)
	})

	t.Run("should return error when failed", func(t *testing.T) {
		reserve := &entity.Reservation{
			CheckInDate: time.Now(),
			CheckOutDate: time.Now(),
			HouseId: 2,
		}
		mockRepository := new(mocks.ReservationRepository)
		usecase := usecase.NewReservationUsecase(&usecase.ReservationUConfig{
			ReservationRepository: mockRepository,
		})
		mockRepository.On("AddReservation", reserve).Return(nil, errors.New("error"))
		expected := errors.New("error")

		_, err := usecase.AddReservation(reserve)
		
		assert.Error(t, expected, err)
	})
}

func TestGetReservationListByUserId(t *testing.T) {
	t.Run("should return reservation list of the user", func(t *testing.T) {
		userId := uint(1)
		pagination := &dto.Pagination{
			Page: 1,
			Limit: 10,
		}
		mockRepository := new(mocks.ReservationRepository)
		usecase := usecase.NewReservationUsecase(&usecase.ReservationUConfig{
			ReservationRepository: mockRepository,
		})
		var expected []*entity.Reservation
		mockRepository.On("GetReservationListByUserId", userId, pagination).Return(expected, 0, nil)

		result, _, _ := usecase.GetReservationListByUserId(userId, pagination)

		assert.Equal(t, expected, result)
	})

	t.Run("should return error when failed", func(t *testing.T) {
		userId := uint(1)
		pagination := &dto.Pagination{
			Page: 1,
			Limit: 10,
		}
		mockRepository := new(mocks.ReservationRepository)
		usecase := usecase.NewReservationUsecase(&usecase.ReservationUConfig{
			ReservationRepository: mockRepository,
		})
		mockRepository.On("GetReservationListByUserId", userId, pagination).Return(nil, 0, errors.New("error"))
		expected := errors.New("error")

		_, _, err := usecase.GetReservationListByUserId(userId, pagination)

		assert.Error(t, expected, err)
	})
}

func TestGetReservationById(t *testing.T) {
	t.Run("should return selected reservation when success", func(t *testing.T) {
		id := uint(1)
		mockRepository := new(mocks.ReservationRepository)
		usecase := usecase.NewReservationUsecase(&usecase.ReservationUConfig{
			ReservationRepository: mockRepository,
		})
		var expected *entity.Reservation
		mockRepository.On("GetReservationById", id).Return(expected, nil)

		result, _ := usecase.GetReservationById(id)

		assert.Equal(t, expected, result)
	})

	t.Run("should return error when failed", func(t *testing.T) {
		id := uint(1)
		mockRepository := new(mocks.ReservationRepository)
		usecase := usecase.NewReservationUsecase(&usecase.ReservationUConfig{
			ReservationRepository: mockRepository,
		})
		mockRepository.On("GetReservationById", id).Return(nil, errors.New("error"))
		expected := errors.New("error")

		_, err := usecase.GetReservationById(id)

		assert.Error(t, expected, err)	
	})
}