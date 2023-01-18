package error

import "errors"

var (
	ErrInvalidCredential = errors.New("invalid user credential")
	ErrDuplicateEntry = errors.New("duplicate entry")
	ErrTokenNotExist = errors.New("token not exist")
	ErrInsufficientBalance = errors.New("insufficient balance")
)