package entity

import "gorm.io/gorm"

type TransactionType struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	Name string `json:"name"`
}

func (TransactionType) TableName() string {
	return "transaction_type_tab"
}