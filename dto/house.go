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
	Price int `form:"price" binding:"required"`
	Description string `form:"description"`
	CityID uint `form:"city_id" binding:"required"`
	MaxGuest int `form:"max_guest"`
	Photos []*multipart.FileHeader `form:"photo"`
}

type UpdateHouse struct {
	Name string `form:"name"`
	Price int `form:"price"`
	Description string `form:"description"`
	CityID uint `form:"city_id"`
	MaxGuest int `form:"max_guest"`
}

func (n *NewHouse) ToHouse() *entity.House {
	return &entity.House{
		Name: n.Name,
		Price: n.Price,
		Description: n.Description,
		CityID: n.CityID,
		MaxGuest: n.MaxGuest,
	}
}

func (u *UpdateHouse) ToHouse() *entity.House {
	return &entity.House{
		Name: u.Name,
		Price: u.Price,
		Description: u.Description,
		CityID: u.CityID,
		MaxGuest: u.MaxGuest,
	}
}