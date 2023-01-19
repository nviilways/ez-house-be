package dto

import "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"

type UserLogin struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name"` 
	Address string `json:"address"`
	CityID uint `json:"city_id"`
}

func (u *UserLogin) ToUser() *entity.User {
	return &entity.User{
		Email: u.Email,
		Password: u.Password,
	}
}

func (u *UserRegister) ToUser() *entity.User {
	return &entity.User{
		Email: u.Email,
		Password: u.Password,
		FullName: u.FullName,
		Address: u.Address,
		CityID: u.CityID,
	}
}