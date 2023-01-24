package entity

import "gorm.io/gorm"

type Role struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	Name string `json:"name"`
}

func (Role) TableName() string {
	return "role_tab"
}