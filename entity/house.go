package entity

import "gorm.io/gorm"

type House struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	Name string `json:"name"`
	UserID uint `json:"user_id"`
	User User `json:"user"`
	Price int `json:"price"`
	Description string `json:"description"`
	CityID int `json:"city_id"`
	City City `json:"city"`
	MaxGuest int `json:"max_guest"`
	Photo []Photo `json:"house_photos"`
}

func (House) TableName() string {
	return "houses_tab"
}