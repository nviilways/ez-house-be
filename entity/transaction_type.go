package entity

import "gorm.io/gorm"

type TransactionType struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	Type string `json:"type"`
}

func (TransactionType) TableName() string {
	return "transactions_type_tab"
}