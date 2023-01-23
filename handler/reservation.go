package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UseAddReservation(c *gin.Context) {
	var reqReserve dto.NewReservation
	err := c.ShouldBindJSON(&reqReserve)
	if err != nil {
		errorTag(c, err)
		return
	}

	claim, _ := c.Get("claim")
	parsedClaim := entity.Claim(claim.(entity.Claim))

	newReserve := reqReserve.ToReservation()
	newReserve.UserId = parsedClaim.ID

	result, err2 := h.reservationUsecase.AddReservation(newReserve)
	if err2 != nil {
		if err2 == errs.ErrDuplicateEntry {
			errorResponse(c, http.StatusBadRequest, err2.Error())
			return
		}
		
		errorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	JSONResponse(c, http.StatusCreated, result)
}