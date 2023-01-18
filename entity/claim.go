package entity

import "github.com/golang-jwt/jwt/v4"

type Claim struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	WalletID uint `json:"wallet_id"`
	RoleID uint `json:"role_id"`
	jwt.RegisteredClaims
}