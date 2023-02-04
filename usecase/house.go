package usecase

import (
	"mime/multipart"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/repository"
)

type HouseUsecase interface {
	GetHouseByID(uint) (*entity.House, error)
	GetCityList() ([]*entity.City, error)
	GetHouseListByVacancy(*dto.FilterHouse, *dto.Pagination) ([]*entity.House, int, error)
	GetHouseByHost(uint, *dto.FilterHouse, *dto.Pagination) ([]*entity.House, int, error)
	AddHouse(*entity.House, []*multipart.FileHeader) (*entity.House, error)
	UpdateHouse(uint, *entity.House) (*entity.House, error)
	DeleteHouse(uint, uint) (*entity.House, error)
}

type houseUsecaseImpl struct {
	houseRepository repository.HouseRepository
}

type HouseUConfig struct {
	HouseRepository repository.HouseRepository
}

func NewHouseUsecase(cfg *HouseUConfig) HouseUsecase {
	return &houseUsecaseImpl{
		houseRepository: cfg.HouseRepository,
	}
}

func (h *houseUsecaseImpl) GetHouseByID(id uint) (*entity.House, error) {
	return h.houseRepository.GetHouseByID(id)
}

func (h *houseUsecaseImpl) GetCityList() ([]*entity.City, error) {
	return h.houseRepository.GetCityList()
}

func (h *houseUsecaseImpl) GetHouseListByVacancy(filter *dto.FilterHouse, pagination *dto.Pagination) ([]*entity.House, int, error) {
	return h.houseRepository.GetHouseListByVacancy(filter, pagination)
}

func (h *houseUsecaseImpl) GetHouseByHost(id uint, filter *dto.FilterHouse, pagination *dto.Pagination) ([]*entity.House, int, error) {
	return h.houseRepository.GetHouseByHost(id, filter, pagination)
}

func (h *houseUsecaseImpl) AddHouse(house *entity.House, photos []*multipart.FileHeader) (*entity.House, error) {
	return h.houseRepository.AddHouse(house, photos)
}

func (h *houseUsecaseImpl) UpdateHouse(id uint, house *entity.House) (*entity.House, error) {
	return h.houseRepository.UpdateHouse(id, house)
}

func (h *houseUsecaseImpl) DeleteHouse(id uint, user_id uint) (*entity.House, error) {
	return h.houseRepository.DeleteHouse(id, user_id)
}
