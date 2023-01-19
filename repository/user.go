package repository

import (
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	SignIn(*entity.User) (*entity.User, error)
	SignUp(*entity.User) (*entity.User, error)
	GetUserByID(uint) (*entity.User, error)
	Update(*entity.User) (*entity.User, error)
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

func (u *userRepositoryImpl) SignUp(user *entity.User) (*entity.User, error) {
	var wallet entity.Wallet
	var games entity.Game

	err := u.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
		if err.Error != nil {
			return err.Error
		}

		if err.RowsAffected == 0 {
			return errs.ErrDuplicateEntry
		}

		if(user.RoleID != 1){
			wallet.UserID = uint(user.ID)
			if err := tx.Create(&wallet).Error; err != nil {
				return err
			}
			user.Wallet = &wallet

			games.UserID = uint(user.ID)
			if err := tx.Create(&games).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepositoryImpl) GetUserByID(user_id uint) (*entity.User, error) {
	var user *entity.User
	err := u.db.Where("id = ?", user_id).Preload("Wallet").Preload("Role").Preload("City").Find(&user).Error

	if(err != nil) {
		return nil, err
	}

	return user, nil
}

func (u *userRepositoryImpl) Update(user *entity.User) (*entity.User, error) {
	err := u.db.Where("id = ?" , user.ID).Clauses(clause.Returning{}).Updates(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}