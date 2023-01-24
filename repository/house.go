package repository

import (
	"context"
	"mime/multipart"
	"time"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/cloud"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HouseRepository interface {
	GetHouseByID(uint) (*entity.House, error)
	GetHouseList() ([]*entity.House, error)
	GetHouseListByVacancy(time.Time, time.Time) ([]*entity.House, error)
	GetHouseByHost(uint) ([]*entity.House, error)
	AddHouse(*entity.House, []*multipart.FileHeader) (*entity.House, error)
	UpdateHouse(uint, *entity.House) (*entity.House, error)
	DeleteHouse(uint, uint) (*entity.House, error)
	AddPhotoHouse(*entity.Photo) (*entity.Photo, error)
}

type houseRepositoryImpl struct {
	db *gorm.DB
}

type HouseRConfig struct {
	DB *gorm.DB
}

func NewHouseRepository(cfg *HouseRConfig) HouseRepository {
	return &houseRepositoryImpl{
		db: cfg.DB,
	}
}

func (h *houseRepositoryImpl) GetHouseByID(id uint) (*entity.House, error) {
	var house *entity.House

	err := h.db.Preload("Photo").Where("id = ?", id).First(&house).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.ErrRecordNotFound
		}
		return nil, err
	}

	return house, nil
}

func (h *houseRepositoryImpl) GetHouseList() ([]*entity.House, error) {
	var house []*entity.House

	err := h.db.Preload("User").Preload("City").Preload("Photo").Find(&house).Error
	if err != nil {
		return nil, err
	}

	return house, nil
}

func (h *houseRepositoryImpl) GetHouseListByVacancy(in time.Time, out time.Time) ([]*entity.House, error) {
	var house []*entity.House

	err := h.db.Model(&entity.House{}).Joins("reservations_tab", h.db.Not("check_in_date >= ? AND check_out_date <= ?", in, out)).Find(&house).Error
	if err != nil {
		return nil, err
	}

	return house, nil
}

func (h *houseRepositoryImpl) GetHouseByHost(id uint) ([]*entity.House, error) {
	var house []*entity.House

	err := h.db.Where("user_id = ?", id).Find(&house).Error
	if err != nil {
		return nil, err
	}

	return house, nil
}

func (h *houseRepositoryImpl) AddHouse(house *entity.House, photos []*multipart.FileHeader) (*entity.House, error) {
	var ph []*entity.Photo

	err := h.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&house).Error; err != nil {
			return err
		}

		for i := range photos {
			file, errOpen := photos[i].Open()
			if errOpen != nil {
				return errOpen
			}

			resp, errCloud := cloud.Cloud.Upload.Upload(context.Background(), file, uploader.UploadParams{
				Folder: "ez-house",
			})

			if errCloud != nil {
				return errCloud
			}

			ph = append(ph, &entity.Photo{
				HouseID:  house.ID,
				PhotoUrl: resp.SecureURL,
				PublicID: resp.PublicID,
			})
		}

		if err := tx.Create(&ph).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	house.Photo = ph

	return house, nil
}

func (h *houseRepositoryImpl) UpdateHouse(id uint, house *entity.House) (*entity.House, error) {
	err := h.db.Where("id = ?", id).Updates(&house).Error
	if err != nil {
		return nil, err
	}

	return house, nil
}

func (h *houseRepositoryImpl) ValidateHouseOwner(id uint, user_id uint) (bool) {
	affected := h.db.Where("id = ? AND user_id = ?", id, user_id).First(&entity.House{}).RowsAffected
	return affected == 1
}

func (h *houseRepositoryImpl) DeleteHouse(id uint, user_id uint) (*entity.House, error) {
	isValid := h.ValidateHouseOwner(id, user_id)
	if !isValid {
		return nil, errs.ErrRecordNotFound
	}

	var house *entity.House

	err := h.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("house_id = ?", id).Delete(&entity.Photo{}).Error; err != nil {
			return err
		}

		if err := tx.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&house).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return house, nil
}

func (h *houseRepositoryImpl) AddPhotoHouse(photo *entity.Photo) (*entity.Photo, error) {
	err := h.db.Create(&photo).Error
	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (h *houseRepositoryImpl) DeletePhotoHouse(photo *entity.Photo) (*entity.Photo, error) {
	err := h.db.Clauses(clause.Returning{}).Delete(&photo).Error
	if err != nil {
		return nil, err
	}

	return photo, nil
}