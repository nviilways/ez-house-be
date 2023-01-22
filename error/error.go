package error

import "errors"

var (
	ErrInvalidCredential = errors.New("invalid user credential")
	ErrRecordNotFound = errors.New("record not found")
	ErrDuplicateEntry = errors.New("duplicate entry")
	ErrTokenNotExist = errors.New("token not exist")
	ErrInvalidToken = errors.New("invalid token")
	ErrInsufficientBalance = errors.New("insufficient balance")
)