package dto

type Token struct {
	AccessToken string `json:"authorization" binding:"required"`
}