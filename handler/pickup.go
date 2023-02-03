package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePickup(c *gin.Context) {
	var reqPickup dto.PickupRequest
	err := c.ShouldBindJSON(&reqPickup)
	if err != nil {
		errorTag(c, err)
		return
	}

	newPickup := reqPickup.ToPickup()
	result, err2 := h.pickupUsecase.RequestPickup(newPickup)
	if err2 != nil {
		if err2 == errs.ErrDoublePickup {
			errorResponse(c, http.StatusBadRequest, err2.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	JSONResponse(c, http.StatusCreated, result)
}