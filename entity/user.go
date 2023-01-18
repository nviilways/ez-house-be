package entity

import "gorm.io/gorm"

type User struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	Email string `json:"email"`
	Password string `json:"-"`
	FullName string `json:"full_name"`
	Address string `json:"address"`
	CityID uint `json:"city_id"`
	City City `json:"city"`
	RoleID uint `json:"role_id"`
	Role Role `json:"role"`
	Wallet Wallet `json:"wallet"`
}

func (User) TableName() string {
	return "users_tab"
}