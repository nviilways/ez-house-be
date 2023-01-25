package usecase

import (
	"mime/multipart"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/repository"
)

type HouseUsecase interface {
	GetHouseByID(uint) (*entity.House, error)
	GetHouseList() ([]*entity.House, error)
	GetHouseListByVacancy(*dto.FilterHouse) ([]*entity.House, error)
	GetHouseByHost(uint) ([]*entity.House, error)
	AddHouse(*entity.House, []*multipart.FileHeader) (*entity.House, error)
	UpdateHouse(uint, *entity.House) (*entity.House, error)
	DeleteHouse(uint, uint) (*entity.House, error)
	AddPhotoHouse(*entity.Photo, []*multipart.FileHeader) (*entity.Photo, error)
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

func (h *houseUsecaseImpl) GetHouseList() ([]*entity.House, error) {
	return h.houseRepository.GetHouseList()
}

func (h *houseUsecaseImpl) GetHouseListByVacancy(filter *dto.FilterHouse) ([]*entity.House, error) {
	return h.houseRepository.GetHouseListByVacancy(filter)
}

func (h *houseUsecaseImpl) GetHouseByHost(id uint) ([]*entity.House, error) {
	return h.houseRepository.GetHouseByHost(id)
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

func (h *houseUsecaseImpl) AddPhotoHouse(ph *entity.Photo, photos []*multipart.FileHeader) (*entity.Photo, error) {
	return h.houseRepository.AddPhotoHouse(ph)
}
