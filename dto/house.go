package dto

import (
	"mime/multipart"
	"time"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
)

type FilterHouse struct {
	CheckInDate  time.Time
	CheckOutDate time.Time
	SortColumn string
	SortBy string
	SearchName string
	SearchCity string
	SearchGuest int
}

type NewHouse struct {
	Name string `form:"name" binding:"required"`
	Price int `form:"price" binding:"required,min=1"`
	Description string `form:"description" binding:"required"`
	Address string `form:"address" binding:"required"`
	CityID uint `form:"city_id" binding:"required"`
	MaxGuest int `form:"max_guest" binding:"required,min=1"`
	Photos []*multipart.FileHeader `form:"photo" binding:"required"`
}

type UpdateHouse struct {
	Name string `form:"name"`
	Price int `form:"price"`
	Description string `form:"description"`
	Address string `form:"address"`
	CityID uint `form:"city_id"`
	MaxGuest int `form:"max_guest"`
}

func (n *NewHouse) ToHouse() *entity.House {
	return &entity.House{
		Name: n.Name,
		Price: n.Price,
		Description: n.Description,
		Address: n.Address,
		CityID: n.CityID,
		MaxGuest: n.MaxGuest,
	}
}

func (u *UpdateHouse) ToHouse() *entity.House {
	return &entity.House{
		Name: u.Name,
		Price: u.Price,
		Description: u.Description,
		Address: u.Address,
		CityID: u.CityID,
		MaxGuest: u.MaxGuest,
	}
}