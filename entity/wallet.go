package entity

import "gorm.io/gorm"

type Wallet struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	UserID uint `json:"user_id"`
	Balance int `json:"balance"`
}

func (Wallet) TableName() string {
	return "wallets_tab"
}