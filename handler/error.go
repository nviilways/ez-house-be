package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"

	"github.com/gin-gonic/gin"
)

type ApiError struct {
	Field string
	Msg string
}

func errorTag(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, dto.JSONResponse{
		Code: http.StatusBadRequest,
		Message: errs.ErrorCode[http.StatusBadRequest],
		Data: nil,
	})
}

func errorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, dto.JSONResponse{
		Code: code,
		Message: message,
		Data: nil,
	})
}