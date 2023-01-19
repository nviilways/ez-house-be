package handler

import (
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"github.com/gin-gonic/gin"
)

var ResponseCode = map[int]string {
	200: "OK",
	201: "CREATED",
}

func JSONResponse(c *gin.Context ,code int, data any) {
	c.JSON(code, dto.JSONResponse {
		Code: code,
		Message: ResponseCode[code],
		Data: data,
	})
}