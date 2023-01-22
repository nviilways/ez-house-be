package repository

import (
	"time"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HouseRepository interface {
	GetHouseByID(uint) (*entity.House, error)
	GetHouseList() ([]*entity.House, error)
	GetHouseListByVacancy(time.Time, time.Time) ([]*entity.House, error)
	GetHouseByHost(uint) ([]*entity.House, error)
	AddHouse(*entity.House) (*entity.House, error)
	UpdateHouse(uint, *entity.House) (*entity.House, error)
	DeleteHouse(uint) (*entity.House, error)
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

	err := h.db.Where("id = ?", id).First(&house).Error
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

func (h *houseRepositoryImpl) AddHouse(house *entity.House) (*entity.House, error) {
	err := h.db.Create(&house).Error
	if err != nil {
		return nil, err
	}

	return house, nil
}

func (h *houseRepositoryImpl) UpdateHouse(id uint, house *entity.House) (*entity.House, error) {
	err := h.db.Where("id = ?", id).Updates(&house).Error
	if err != nil {
		return nil, err
	}

	return house, nil
}

func (h *houseRepositoryImpl) DeleteHouse(id uint) (*entity.House, error) {
	var house *entity.House

	err := h.db.Where("id = ?", id).Clauses(clause.Returning{}).Delete(&house).Error
	if err != nil {
		return nil, err
	}

	return house, nil
}
