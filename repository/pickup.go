package repository

import (
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"gorm.io/gorm"
)

const updateStatus = 1
const pickupCostRate = 100000

type PickupRepository interface {
	GetPickupById(uint) (*entity.Pickup, error)
	GetPickupList() ([]*entity.Pickup, error)
	GetPickupPrice(*entity.Reservation) (*dto.PickupPrice, error)
	RequestPickup(*entity.Pickup) (*entity.Pickup, error)
	UpdateStatus(uint) (*entity.Pickup, error)
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
	var user *entity.User
	var reservation *entity.Reservation

	errUser := r.db.Where("id = ?", pick.UserID).Preload("Wallet").First(&user).Error
	if errUser != nil {
		return nil, errUser
	}

	errReservation := r.db.Where("id = ?", pick.ReservationID).Preload("House").First(&reservation).Error
	if errReservation != nil {
		return nil, errReservation
	}

	errExist := r.db.Where("user_id = ? AND reservation_id = ?", pick.UserID, pick.ReservationID).First(&entity.Pickup{}).RowsAffected
	if errExist != 0 {
		return nil, errs.ErrDoublePickup
	}

	price := pickupCostRate
	if(reservation.House.CityID != user.CityID) {
		price *= 3
	}

	if(user.Wallet.Balance < price) {
		return nil, errs.ErrInsufficientBalance
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&pick).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.Wallet{}).Where("id = ?", user.Wallet.ID).Update("balance", gorm.Expr("balance - ?", price)).Error; err != nil {
			return err
		}

		transaction := &entity.Transaction{
			WalletID: user.Wallet.ID,
			TransactionTypeID: 2,
			Balance: price,
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.Reservation{}).Where("id = ?", pick.ReservationID).Update("total_price", gorm.Expr("total_price + ?", price)).Error; err != nil {
			return err
		}
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}

	return pick, nil
}

func (r *pickupRepositoryImpl) UpdateStatus(id uint) (*entity.Pickup, error) {
	var pickup *entity.Pickup

	err := r.db.Where("id = ?", id).First(&pickup).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.ErrRecordNotFound
		}
		return nil, err
	}

	if(pickup.PickupStatusID == 5) {
		return nil, errs.ErrCompletedPickup
	}

	err = r.db.Model(&pickup).Update("status_pickup_id", gorm.Expr("status_pickup_id + ?", updateStatus)).Error
	if err != nil {
		return nil, err
	}

	return pickup, nil
}