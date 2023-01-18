package dto

import "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"

type UserLogin struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"-" binding:"required"`
}

func (u *UserLogin) ToUser() *entity.User {
	return &entity.User{
		Email: u.Email,
		Password: u.Password,
	}
}