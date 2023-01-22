package dto

import (
	"time"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
)

type FilterHouse struct {
	CheckInDate  time.Time `json:"check_in"`
	CheckOutDate time.Time `json:"check_out"`
}

type NewHouse struct {
	Name string `json:"name" binding:"required"`
	Price int `json:"price" binding:"required"`
	Description string `json:"description"`
	CityID int `json:"city_id" binding:"required"`
	MaxGuest int `json:"max_guest"`
}

type UpdateHouse struct {
	Name string `json:"name"`
	Price int `json:"price"`
	Description string `json:"description"`
	CityID int `json:"city_id"`
	MaxGuest int `json:"max_guest"`
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