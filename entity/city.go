package entity

import "gorm.io/gorm"

type City struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	Name string `json:"name"`
}

func (City) TableName() string {
	return "cities_tab"
}