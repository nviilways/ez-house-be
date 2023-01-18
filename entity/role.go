package entity

import "gorm.io/gorm"

type Role struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	Name string `json:"name"`
	UserID uint `json:"-"`
}

func (Role) TableName() string {
	return "roles_tab"
}