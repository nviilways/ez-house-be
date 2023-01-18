package entity

import "gorm.io/gorm"

type City struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	Name string `json:"name"`
	UserID uint `json:"-"`	
}