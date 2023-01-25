package dto

import (
	"time"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
)

type NewReservation struct {
	HouseId      uint      `json:"house_id" binding:"required"`
	CheckInDate  string `json:"check_in_date" binding:"required"`
	CheckOutDate string `json:"check_out_date" binding:"required"`
	TotalPrice   int       `json:"total_price"`
}

type PickupPrice struct {
	Price int `json:"price"`
}

func (n NewReservation) ToReservation() *entity.Reservation {
	checkIn, _ := time.Parse("2006-01-02", n.CheckInDate)
	checkOut, _ := time.Parse("2006-01-02", n.CheckOutDate)
	return &entity.Reservation{
		HouseId:      n.HouseId,
		CheckInDate:  checkIn,
		CheckOutDate: checkOut,
		TotalPrice:   n.TotalPrice,
	}
}
