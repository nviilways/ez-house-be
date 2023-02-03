package dto

import "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"

type PickupRequest struct {
	UserID uint `json:"user_id" binding:"required"`
	ReservationID uint `json:"reservation_id" binding:"required"`
}

func (p *PickupRequest) ToPickup() *entity.Pickup {
	return &entity.Pickup{
		UserID: p.UserID,
		ReservationID: p.ReservationID,
	}
}