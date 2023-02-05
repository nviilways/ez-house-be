package entity

import "gorm.io/gorm"

type User struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	Email string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Address string `json:"address"`
	CityID uint `json:"city_id"`
	City *City `json:"city,omitempty"`
	RoleID uint `json:"role_id"`
	Role *Role `json:"role,omitempty"`
	Wallet *Wallet `json:"wallet,omitempty"`
}

func (User) TableName() string {
	return "user_tab"
}