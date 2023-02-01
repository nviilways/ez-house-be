package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	WalletID uint `json:"wallet_id"`
	TransactionTypeID uint `json:"transaction_type_id"`
	TransactionType *TransactionType `json:"transaction_type,omitempty"`
	Balance int `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
 }

func (Transaction) TableName() string {
	return "transaction_tab"
}