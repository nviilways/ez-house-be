package repository

import (
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"gorm.io/gorm"
)

type UserRepository interface {
	SignIn(*entity.User) (*entity.User, error)
	Register(*entity.User) (*entity.User, error)
	GetUserByID(*entity.User) (*entity.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

type UserRConfig struct {
	DB *gorm.DB
}

func NewUserRepository(cfg *UserRConfig) UserRepository {
	return &userRepositoryImpl {
		db: cfg.DB,
	}
}

func (u *userRepositoryImpl) SignIn(user *entity.User) (*entity.User, error) {
	err := u.db.Where("email = ?", user.Email).Preload("Wallet").Find(&user).Error

	if(err != nil) {
		return nil, err
	}

	return user, nil
}

func (u *userRepositoryImpl) Register(user *entity.User) (*entity.User, error) {
	var wallet entity.Wallet

	err := u.db.Transaction(func(tx *gorm.DB) error {
		if affected := tx.FirstOrCreate(&user).RowsAffected; affected == 0 {
			return errs.ErrDuplicateEntry
		}

		wallet.UserID = uint(user.ID)
		if err := tx.Create(&wallet).Error; err != nil {
			return err
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepositoryImpl) GetUserByID(user *entity.User) (*entity.User, error) {
	err := u.db.Where("id = ?", user.ID).Preload("Wallet").Find(&user).Error

	if(err != nil) {
		return nil, err
	}

	return user, nil
}