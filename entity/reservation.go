package entity

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	HouseId uint `json:"house_id"`
	House *House `json:"house,omitempty"`
	UserId uint `json:"user_id"`
	User *User `json:"user,omitempty"`
	CheckInDate time.Time `json:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date"`
	TotalPrice int `json:"total_price"`
	BookingCode string `json:"booking_code"`
	Pickup *Pickup `json:"pickup,omitempty"`
}

func (Reservation) TableName() string {
	return "reservation_tab"
}