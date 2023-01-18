package entity

import "gorm.io/gorm"

type Transaction struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	WalletID uint `json:"wallet_id"`
	TransactionTypeID uint `json:"transaction_type_id"`
	TransactionType TransactionType `json:"transaction_type"`
	Balance int `json:"balance"`
 }

func (Transaction) TableName() string {
	return "transactions_tab"
}