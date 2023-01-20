package repository

import (
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AdminRepository interface {
	SignIn(*entity.Admin) (*entity.Admin, error)
	SignUp(*entity.Admin) (*entity.Admin, error)
}

type adminRepositoryImpl struct {
	db *gorm.DB
}

type AdminRConfig struct {
	DB *gorm.DB
}

func NewAdminRepository(cfg *AdminRConfig) AdminRepository {
	return &adminRepositoryImpl {
		db: cfg.DB,
	}
}

func (a *adminRepositoryImpl) SignIn(admin *entity.Admin) (*entity.Admin, error) {
	err := a.db.Where("email = ?", admin.Email).Find(&admin).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (a *adminRepositoryImpl) SignUp(admin *entity.Admin) (*entity.Admin, error) {
	admin.RoleID = 1
	err := a.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&admin)

	if err.Error != nil {
		return nil, err.Error
	}

	if err.RowsAffected == 0 {
		return nil, errs.ErrDuplicateEntry
	}

	return admin, nil
}