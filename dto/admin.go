package dto

import "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"

type AdminRegister struct {
	Email string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required"`
}

func (a *AdminRegister) ToAdmin() *entity.Admin {
	return &entity.Admin{
		Email: a.Email,
		Password: a.Password,
	}
}