package entity

import "gorm.io/gorm"

type Pickup struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	UserID uint `json:"user_id"`
	ReservationID uint `json:"reservation_id"`
	PickupStatusID uint `json:"pickup_status_id"`
}

func (Pickup) TableName() string {
	return "pickup_tab"
}