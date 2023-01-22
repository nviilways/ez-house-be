package handler

import (
	"errors"
	"net/http"
	"strconv"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UserGetHouseById(c *gin.Context) {
	id, errParse := strconv.Atoi(c.Param("id"))
	if errParse != nil {
		errorResponse(c, http.StatusBadRequest, errParse.Error())
		return
	}
	result, err := h.houseUsecase.GetHouseByID(uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			errorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	JSONResponse(c, http.StatusOK, result)
}

func (h *Handler) UserGetHouseList(c *gin.Context) {
	result, err := h.houseUsecase.GetHouseList()
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	JSONResponse(c, http.StatusOK, result)
}

func (h *Handler) UserGetHouseByVacancy(c *gin.Context) {
	var reqFilter dto.FilterHouse
	err := c.ShouldBindJSON(&reqFilter)
	if err != nil {
		errorTag(c, err)
		return
	}

	result, err2 := h.houseUsecase.GetHouseListByVacancy(reqFilter.CheckInDate, reqFilter.CheckOutDate)
	if err2 != nil {
		errorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	JSONResponse(c, http.StatusOK, result)
}

func (h *Handler) UserGetHouseByHost(c *gin.Context) {
	claim, _ := c.Get("claim")
	parsedClaim := entity.Claim(claim.(entity.Claim))

	result, err := h.houseUsecase.GetHouseByHost(parsedClaim.ID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	JSONResponse(c, http.StatusOK, result)
}

func (h *Handler) HostAddHouse(c *gin.Context) {
	claim, _ := c.Get("claim")
	parsedClaim := entity.Claim(claim.(entity.Claim))

	var reqNewHouse dto.NewHouse
	err := c.ShouldBindJSON(&reqNewHouse)
	if err != nil {
		errorTag(c, err)
		return
	}

	newHouse := reqNewHouse.ToHouse()
	newHouse.UserID = parsedClaim.ID
	result, err2 := h.houseUsecase.AddHouse(newHouse)
	if err2 != nil {
		errorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	JSONResponse(c, http.StatusCreated, result)
}

func (h *Handler) HostUpdateHouse(c *gin.Context) {
	claim, _ := c.Get("claim")
	parsedClaim := entity.Claim(claim.(entity.Claim))

	var reqUpdateHouse dto.UpdateHouse
	err := c.ShouldBindJSON(&reqUpdateHouse)
	if err != nil {
		errorTag(c, err)
		return
	}
	
	updateHouse := reqUpdateHouse.ToHouse()
	updateHouse.UserID = parsedClaim.ID
	result, err2 := h.houseUsecase.UpdateHouse(parsedClaim.ID, updateHouse)
	if err2 != nil {
		errorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	JSONResponse(c, http.StatusOK, result)
}

func (h *Handler) HostDeleteHouse(c *gin.Context) {
	id, errParse := strconv.Atoi(c.Param("id"))
	if errParse != nil {
		errorResponse(c, http.StatusBadRequest, errParse.Error())
		return
	}

	result, err := h.houseUsecase.DeleteHouse(uint(id))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	JSONResponse(c, http.StatusOK, result)
}