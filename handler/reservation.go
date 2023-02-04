package handler

import (
	"net/http"
	"strconv"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UserAddReservation(c *gin.Context) {
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

func (h *Handler) UserGetReservationById(c *gin.Context) {
	id, errParse := strconv.Atoi(c.Param("id"))
	if errParse != nil {
		errorResponse(c, http.StatusBadRequest, errParse.Error())
		return
	}

	result, err := h.reservationUsecase.GetReservationById(uint(id))
	if err != nil {
		if err == errs.ErrRecordNotFound {
			errorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	JSONResponse(c, http.StatusOK, result)
}

func (h *Handler) UserGetReservationByUserId(c *gin.Context) {
	claim, _ := c.Get("claim")
	parsedClaim := entity.Claim(claim.(entity.Claim))

	page, err := strconv.Atoi(c.Query("page"))
	if page == 0 || err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if limit == 0 || err != nil {
		limit = 10
	}

	pagination := &dto.Pagination{
		Limit: limit,
		Page: page,
	}

	result, count, err := h.reservationUsecase.GetReservationListByUserId(parsedClaim.ID, pagination)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	pagination.Count = count
	pagination.Data = result

	JSONResponse(c, http.StatusOK, pagination)
}