package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/repository"
)

type PickupUsecase interface {
	GetPickupById(uint) (*entity.Pickup, error)
	GetPickupList() ([]*entity.Pickup, error)
	GetPickupPrice(*entity.Reservation) (*dto.PickupPrice, error)
	RequestPickup(*entity.Pickup) (*entity.Pickup, error)
	UpdateStatus(uint) (*entity.Pickup, error)
}

type pickupUsecaseImpl struct {
	pickupRepository repository.PickupRepository
}

type PickupUConfig struct {
	PickupRepository repository.PickupRepository
}

func NewPickupUsecase(cfg *PickupUConfig) PickupUsecase {
	return &pickupUsecaseImpl{
		pickupRepository: cfg.PickupRepository,
	}
}

func (p *pickupUsecaseImpl) GetPickupById(id uint) (*entity.Pickup, error) {
	return p.pickupRepository.GetPickupById(id)
}

func (p *pickupUsecaseImpl) GetPickupList() ([]*entity.Pickup, error) {
	return p.pickupRepository.GetPickupList()
}

func (p *pickupUsecaseImpl) GetPickupPrice(res *entity.Reservation) (*dto.PickupPrice, error) {
	return p.pickupRepository.GetPickupPrice(res)
}

func (p *pickupUsecaseImpl) RequestPickup(pickup *entity.Pickup) (*entity.Pickup, error) {
	return p.pickupRepository.RequestPickup(pickup)
}

func (p *pickupUsecaseImpl) UpdateStatus(id uint) (*entity.Pickup, error) {
	return p.pickupRepository.UpdateStatus(id)
}
