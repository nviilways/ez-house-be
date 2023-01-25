package repository

import (
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"gorm.io/gorm"
)

type PickupRepository interface {
	GetPickupById(uint) (*entity.Pickup, error)
	GetPickupList() ([]*entity.Pickup, error)
	GetPickupPrice(*entity.Reservation) (*dto.PickupPrice, error)
	RequestPickup(*entity.Pickup) (*entity.Pickup, error)
}

type pickupRepositoryImpl struct {
	db *gorm.DB
}

type PickupRConfig struct {
	DB *gorm.DB
}

func NewPickupRepository(cfg *PickupRConfig) PickupRepository {
	return &pickupRepositoryImpl{
		db: cfg.DB,
	}
}

func (p *pickupRepositoryImpl) GetPickupById(id uint) (*entity.Pickup, error) {
	var pickup *entity.Pickup

	err := p.db.Where("id = ?", id).First(&pickup).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.ErrRecordNotFound
		}
		return nil, err
	}

	return pickup, nil

}

func (p *pickupRepositoryImpl) GetPickupList() ([]*entity.Pickup, error) {
	var pickup []*entity.Pickup

	err := p.db.Find(&pickup).Error 
	if err != nil {
		return nil, err
	}

	return pickup, nil
}

func (r *pickupRepositoryImpl) GetPickupPrice(res *entity.Reservation) (*dto.PickupPrice, error) {
	err := r.db.Where("id = ?", res.ID).Preload("House").Preload("User").First(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.ErrRecordNotFound
		}
		return nil, err
	}

	var price *dto.PickupPrice
	price.Price = pickupCostRate
	if res.User.CityID != res.House.CityID {
		price.Price *= 3
	}

	return price, nil
}

func (r *pickupRepositoryImpl) RequestPickup(pick *entity.Pickup) (*entity.Pickup, error) {
	pick.PickupStatusID = 1

	errExist := r.db.Where("user_id = ? AND reservation_id = ?", pick.UserID, pick.ReservationID).First(&entity.Pickup{}).Error
	if errExist != nil {
		return nil, errs.ErrDoublePickup
	}

	err := r.db.Create(&pick).Error
	if err != nil {
		return nil, err
	}

	return pick, nil
}
