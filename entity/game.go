package entity

import "gorm.io/gorm"

type Game struct {
	ID uint `json:"id"`
	gorm.Model `json:"-"`
	UserID uint `json:"user_id"`
	User User `json:"user"`
	Chance int `json:"chance"`
	Count int `json:"count"`
}

func (Game) TableName() string {
	return "games_chance_tab"
}