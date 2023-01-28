package dto

type Pagination struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
	Count int `json:"count"`
	Data any `json:"data"`
}