package error

import "errors"

var (
	ErrInvalidCredential = errors.New("invalid user credential")
	ErrRecordNotFound = errors.New("record not found")
	ErrDuplicateEntry = errors.New("duplicate entry")
	ErrTokenNotExist = errors.New("token not exist")
	ErrInvalidToken = errors.New("invalid token")
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrInvalidCheckDate = errors.New("failed to check in between those dates")
	ErrBookForZeroDays = errors.New("cannot book under 1 day")
	ErrAlreadyReserved = errors.New("house already reserved")
	ErrHouseNotFound = errors.New("house not found")
	ErrReserveOwnHouse = errors.New("cannot reserve owned house")
)