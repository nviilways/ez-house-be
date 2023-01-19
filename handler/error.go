package handler

import (
	"net/http"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"

	"github.com/gin-gonic/gin"
)

type ApiError struct {
	Field string
	Msg string
}

func errorTag(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H {
		"error" : err.Error(),
		"code" : errs.ErrorCode[400],
		"message" : "invalid input",
	})
}

func errorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    errs.ErrorCode[code],
		"message": message,
	})
}