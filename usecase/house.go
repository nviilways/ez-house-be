package usecase

import (
	"time"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/repository"
)

type HouseUsecase interface {
	GetHouseByID(uint) (*entity.House, error)
	GetHouseList() ([]*entity.House, error)
	GetHouseListByVacancy(time.Time, time.Time) ([]*entity.House, error)
	GetHouseByHost(uint) ([]*entity.House, error)
	AddHouse(*entity.House) (*entity.House, error)
	UpdateHouse(uint, *entity.House) (*entity.House, error)
	DeleteHouse(uint, *entity.House) (*entity.House, error)
}

type houseUsecaseImpl struct {
	houseRepository repository.HouseRepository
}

type HouseUConfig struct {
	HouseRepository repository.HouseRepository
}

func NewHouseRepository(cfg *HouseUConfig) HouseUsecase {
	return &houseUsecaseImpl {
		houseRepository: cfg.HouseRepository,
	}
}

func (h *houseUsecaseImpl) GetHouseByID(id uint) (*entity.House, error) {
	return h.houseRepository.GetHouseByID(id)
}

func (h *houseUsecaseImpl) GetHouseList() ([]*entity.House, error) {
	return h.houseRepository.GetHouseList()
}

func (h *houseUsecaseImpl) GetHouseListByVacancy(in time.Time, out time.Time) ([]*entity.House, error) {
	return h.houseRepository.GetHouseListByVacancy(in, out)
}

func (h *houseUsecaseImpl) GetHouseByHost(id uint) ([]*entity.House, error) {
	return h.houseRepository.GetHouseByHost(id)
}

func (h *houseUsecaseImpl) AddHouse(house *entity.House) (*entity.House, error) {
	return h.houseRepository.AddHouse(house)
}

func (h *houseUsecaseImpl) UpdateHouse(id uint, house *entity.House) (*entity.House, error) {
	return h.houseRepository.UpdateHouse(id, house)
}

func (h *houseUsecaseImpl) DeleteHouse(id uint, house *entity.House) (*entity.House, error) {
	return h.houseRepository.DeleteHouse(id, house)
}
