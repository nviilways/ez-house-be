package entity

type PickupStatus struct {
	ID uint `json:"id"`
	Status string `json:"status"`
}

func (PickupStatus) TableName() string {
	return "pickup_status_tab"
}