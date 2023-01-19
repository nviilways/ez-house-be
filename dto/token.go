package dto

type Token struct {
	AccessToken string `json:"token" binding:"required"`
}