package entity

import "gorm.io/gorm"

type Token struct {
	ID uint
	gorm.Model
	Token string
}

func (Token) TableName() string {
	return "invalid_token_tab"
}