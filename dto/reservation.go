package dto

import (
	"time"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
)

type NewReservation struct {
	HouseId      uint      `json:"house_id" binding:"required"`
	CheckInDate  time.Time `json:"check_in_date" binding:"required"`
	CheckOutDate time.Time `json:"check_out_date" binding:"required"`
	TotalPrice   int       `json:"total_price" binding:"required"`
}

func (n NewReservation) ToReservation() *entity.Reservation {
	return &entity.Reservation{
		HouseId:      n.HouseId,
		CheckInDate:  n.CheckInDate,
		CheckOutDate: n.CheckOutDate,
		TotalPrice:   n.TotalPrice,
	}
}
