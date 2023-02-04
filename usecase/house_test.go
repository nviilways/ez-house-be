package usecase_test

import (
	"errors"
	"mime/multipart"
	"testing"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	mocks "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/mocks/repository"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetHouseByID(t *testing.T) {
	t.Run("should return selected house", func(t *testing.T) {
		id := uint(1)
		house := &entity.House{
			ID: 1,
		}
		mockRepository := new(mocks.HouseRepository)
		usecase := usecase.NewHouseUsecase(&usecase.HouseUConfig{
			HouseRepository: mockRepository,
		})
		mockRepository.On("GetHouseByID", id).Return(house, nil)

		result, _ := usecase.GetHouseByID(id)

		assert.Equal(t, house, result)
	})

	t.Run("should return error when failed", func(t *testing.T) {
		id := uint(1)
		mockRepository := new(mocks.HouseRepository)
		usecase := usecase.NewHouseUsecase(&usecase.HouseUConfig{
			HouseRepository: mockRepository,
		})
		mockRepository.On("GetHouseByID", id).Return(nil, errors.New("error"))
		expected := errors.New("error")

		_, err := usecase.GetHouseByID(id)

		assert.Error(t, expected, err)
	})
}

func TestGetCityList(t *testing.T) {
	t.Run("should return list of city", func(t *testing.T) {
		var city []*entity.City
		mockRepository := new(mocks.HouseRepository)
		usecase := usecase.NewHouseUsecase(&usecase.HouseUConfig{
			HouseRepository: mockRepository,
		})
		mockRepository.On("GetCityList").Return(city, nil)

		result, _ := usecase.GetCityList()

		assert.Equal(t, city, result)
	})

	t.Run("should return error when failed", func(t *testing.T) {
		mockRepository := new(mocks.HouseRepository)
		usecase := usecase.NewHouseUsecase(&usecase.HouseUConfig{
			HouseRepository: mockRepository,
		})
		mockRepository.On("GetCityList").Return(nil, errors.New("error"))
		expected := errors.New("error")

		_, err := usecase.GetCityList()

		assert.Error(t, expected, err)
	})
}

func TestGetHouseListByVacancy(t *testing.T) {
	t.Run("should return vacant house list", func(t *testing.T) {
		var houses []*entity.House
		var filter *dto.FilterHouse
		var pagination *dto.Pagination
		mockRepository := new(mocks.HouseRepository)
		usecase := usecase.NewHouseUsecase(&usecase.HouseUConfig{
			HouseRepository: mockRepository,
		})
		mockRepository.On("GetHouseListByVacancy", filter, pagination).Return(houses, 0, nil)

		result, _, _ := usecase.GetHouseListByVacancy(filter, pagination)

		assert.Equal(t, houses, result)
	})
	
	t.Run("should return error when failed", func(t *testing.T) {
		var filter *dto.FilterHouse
		var pagination *dto.Pagination
		mockRepository := new(mocks.HouseRepository)
		usecase := usecase.NewHouseUsecase(&usecase.HouseUConfig{
			HouseRepository: mockRepository,
		})
		mockRepository.On("GetHouseListByVacancy", filter, pagination).Return(nil, 0, errors.New("error"))
		expected := errors.New("error")

		_, _, err := usecase.GetHouseListByVacancy(filter, pagination)
	
		assert.Error(t, expected, err)
	})
}

func TestGetHouseByHost(t *testing.T) {
	t.Run("should return vacant house list", func(t *testing.T) {
		var houses []*entity.House
		userId := uint(1)
		var filter *dto.FilterHouse
		var pagination *dto.Pagination
		mockRepository := new(mocks.HouseRepository)
		usecase := usecase.NewHouseUsecase(&usecase.HouseUConfig{
			HouseRepository: mockRepository,
		})
		mockRepository.On("GetHouseByHost", userId, filter, pagination).Return(houses, 0, nil)

		result, _, _ := usecase.GetHouseByHost(userId ,filter, pagination)

		assert.Equal(t, houses, result)
	})
	
	t.Run("should return error when failed", func(t *testing.T) {
		userId := uint(1)
		var filter *dto.FilterHouse
		var pagination *dto.Pagination
		mockRepository := new(mocks.HouseRepository)
		usecase := usecase.NewHouseUsecase(&usecase.HouseUConfig{
			HouseRepository: mockRepository,
		})
		mockRepository.On("GetHouseByHost", userId, filter, pagination).Return(nil, 0, errors.New("error"))
		expected := errors.New("error")

		_, _, err := usecase.GetHouseByHost(userId, filter, pagination)
	
		assert.Error(t, expected, err)
	})
}

func TestAddHouse(t *testing.T) {
	t.Run("should return created house when successful", func(t *testing.T) {
		var house *entity.House
		var photos []*multipart.FileHeader
		mockRepository := new(mocks.HouseRepository)
		usecase := usecase.NewHouseUsecase(&usecase.HouseUConfig{
			HouseRepository: mockRepository,
		})
		mockRepository.On("AddHouse", house, photos).Return(house, nil)

		result, _ := usecase.AddHouse(house, photos)

		assert.Equal(t, house, result)
	})

	t.Run("should return error when failed to create house", func(t *testing.T) {
		var house *entity.House
		var photos []*multipart.FileHeader
		mockRepository := new(mocks.HouseRepository)
		usecase := usecase.NewHouseUsecase(&usecase.HouseUConfig{
			HouseRepository: mockRepository,
		})
		mockRepository.On("AddHouse", house, photos).Return(nil, errors.New("error"))
		expected := errors.New("error")

		_, err := usecase.AddHouse(house, photos)

		assert.Error(t, expected, err)
	})
}

func TestUpdateHouse(t *testing.T) {
	t.Run("should return created house when successful", func(t *testing.T) {
		var house *entity.House
		id := uint(1)
		mockRepository := new(mocks.HouseRepository)
		usecase := usecase.NewHouseUsecase(&usecase.HouseUConfig{
			HouseRepository: mockRepository,
		})
		mockRepository.On("UpdateHouse", id, house).Return(house, nil)

		result, _ := usecase.UpdateHouse(id, house)

		assert.Equal(t, house, result)
	})

	t.Run("should return error when failed to create house", func(t *testing.T) {
		var house *entity.House
		id := uint(1)
		mockRepository := new(mocks.HouseRepository)
		usecase := usecase.NewHouseUsecase(&usecase.HouseUConfig{
			HouseRepository: mockRepository,
		})
		mockRepository.On("UpdateHouse", id, house).Return(nil, errors.New("error"))
		expected := errors.New("error")

		_, err := usecase.UpdateHouse(id, house)

		assert.Error(t, expected, err)
	})
}

func TestDeleteHouse(t *testing.T){
	t.Run("should return deleted house when successful", func(t *testing.T) {
		userId := uint(1)
		houseId := uint(1)
		var house *entity.House
		mockRepository := new(mocks.HouseRepository)
		usecase := usecase.NewHouseUsecase(&usecase.HouseUConfig{
			HouseRepository: mockRepository,
		})
		mockRepository.On("DeleteHouse", houseId, userId).Return(house, nil)

		result, _ := usecase.DeleteHouse(houseId, userId)

		assert.Equal(t, house, result)
	})

	t.Run("should return error when failed to delete house", func(t *testing.T) {
		userId := uint(1)
		houseId := uint(1)
		mockRepository := new(mocks.HouseRepository)
		usecase := usecase.NewHouseUsecase(&usecase.HouseUConfig{
			HouseRepository: mockRepository,
		})
		mockRepository.On("DeleteHouse", houseId, userId).Return(nil, errors.New("error"))
		expected := errors.New("error")

		_, err := usecase.DeleteHouse(houseId, userId)

		assert.Error(t, expected, err)
	})
}