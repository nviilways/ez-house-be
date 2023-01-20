package entity

import "gorm.io/gorm"

type Admin struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	Email string `json:"email"`
	Password string `json:"-"`
	RoleID uint `json:"role_id"`
}

func (Admin) TableName() string {
	return "admin_tab"
}