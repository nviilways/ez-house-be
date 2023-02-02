package repository

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/cloud"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HouseRepository interface {
	GetHouseByID(uint) (*entity.House, error)
	GetCityList() ([]*entity.City, error)
	GetHouseListByVacancy(*dto.FilterHouse, *dto.Pagination) ([]*entity.House, int, error)
	GetHouseByHost(uint) ([]*entity.House, error)
	AddHouse(*entity.House, []*multipart.FileHeader) (*entity.House, error)
	UpdateHouse(uint, *entity.House) (*entity.House, error)
	DeleteHouse(uint, uint) (*entity.House, error)
	AddPhotoHouse(*entity.Photo) (*entity.Photo, error)
	DeletePhotoHouse(*entity.Photo) (*entity.Photo, error)
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

	err := h.db.Preload("Photo").Preload("City").Where("id = ?", id).First(&house).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.ErrRecordNotFound
		}
		return nil, err
	}

	return house, nil
}

func (h *houseRepositoryImpl) GetCityList() ([]*entity.City, error) {
	var city []*entity.City

	err := h.db.Find(&city).Error
	if err != nil {
		return nil, err
	}

	return city, nil
}

func (h *houseRepositoryImpl) GetHouseListByVacancy(filter *dto.FilterHouse, pagination *dto.Pagination) ([]*entity.House, int,error) {
	var house []*entity.House
	var count int64
	var columnKey = map[string]string {
		"name": "house_tab.name",
		"price": "house_tab.price",
		"city": "city_tab.name",
	}

	query := h.db.Model(&entity.House{}).Joins("LEFT JOIN city_tab ON city_tab.id = house_tab.city_id").Joins("LEFT JOIN reservation_tab ON reservation_tab.house_id = house_tab.id AND reservation_tab.check_in_date <= ? AND reservation_tab.check_out_date > ?", filter.CheckOutDate, filter.CheckInDate).Where("reservation_tab.id IS NULL").Where("house_tab.name ILIKE ?", fmt.Sprintf("%%%s%%", filter.SearchName)).Where("city_tab.name ILIKE ?", fmt.Sprintf("%%%s%%", filter.SearchCity)).Where("house_tab.max_guest >= ?", filter.SearchGuest).Order(fmt.Sprintf("%s %s", columnKey[filter.SortColumn], filter.SortBy)).Limit(pagination.Limit).Offset((pagination.Page - 1) * pagination.Limit ).Preload("City").Preload("Photo").Find(&house)
	if query.Error != nil {
		return nil, 0, query.Error
	}

	countQuery := h.db.Model(&entity.House{}).Joins("LEFT JOIN city_tab ON city_tab.id = house_tab.city_id").Joins("LEFT JOIN reservation_tab ON reservation_tab.house_id = house_tab.id AND reservation_tab.check_in_date <= ? AND reservation_tab.check_out_date > ?", filter.CheckOutDate, filter.CheckInDate).Where("reservation_tab.id IS NULL").Where("house_tab.name ILIKE ?", fmt.Sprintf("%%%s%%", filter.SearchName)).Where("city_tab.name ILIKE ?", fmt.Sprintf("%%%s%%", filter.SearchCity)).Where("house_tab.max_guest >= ?", filter.SearchGuest).Count(&count)
	if countQuery.Error != nil {
		return nil, 0, countQuery.Error
	}

	retCount := int(count)

	return house, retCount, nil
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
	isValid := h.ValidateHouseOwner(house.ID, id)
	if !isValid {
		return nil, errs.ErrNotHouseOwner
	}
	
	err := h.db.Updates(&house).Error
	if err != nil {
		return nil, err
	}

	return house, nil
}

func (h *houseRepositoryImpl) ValidateHouseOwner(id uint, user_id uint) bool {
	affected := h.db.Where("id = ? AND user_id = ?", id, user_id).First(&entity.House{}).RowsAffected
	return affected == 1
}

func (h *houseRepositoryImpl) ValidateHouseNoReservation(id uint) bool {
	affected := h.db.Where("house_id = ? AND check_out_date >= ?", id, time.Now()).First(&entity.Reservation{}).RowsAffected
	return affected == 1
}

func (h *houseRepositoryImpl) DeleteHouse(id uint, user_id uint) (*entity.House, error) {
	isValid := h.ValidateHouseOwner(id, user_id)
	if !isValid {
		return nil, errs.ErrNotHouseOwner
	}

	isReserved := h.ValidateHouseNoReservation(id)
	if isReserved {
		return nil, errs.ErrHouseStillReserved
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
