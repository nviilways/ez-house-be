package entity

import "gorm.io/gorm"

type Photo struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	HouseID uint `json:"house_id"`
	PublicID string `json:"public_id"`
	PhotoUrl string `json:"photo_url"`
}

func (Photo) TableName() string {
	return "houses_photos_tab"
}