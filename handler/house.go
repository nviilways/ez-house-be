package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

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
	var inDate time.Time
	checkIn := c.Query("checkin")
	if checkIn != "" {
		inDate, _ = time.Parse("2006-01-02", checkIn)
	}

	var outDate time.Time
	checkOut := c.Query("checkout")
	if checkOut != "" {
		outDate, _ = time.Parse("2006-01-02", checkOut)
	}

	sortColumn := c.Query("sort")
	if sortColumn == "" {
		sortColumn = "name"
	}

	sortBy := c.Query("sortby")
	if sortBy == "" {
		sortBy = "asc"
	}

	searchName := c.Query("searchname")
	searchCity := c.Query("searchcity")

	reqFilter := &dto.FilterHouse{
		CheckInDate:  inDate,
		CheckOutDate: outDate,
		SortColumn:   sortColumn,
		SortBy:       sortBy,
		SearchName:   searchName,
		SearchCity:   searchCity,
	}

	result, err2 := h.houseUsecase.GetHouseListByVacancy(reqFilter)
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
	err := c.ShouldBind(&reqNewHouse)
	if err != nil {
		errorTag(c, err)
		return
	}

	newHouse := reqNewHouse.ToHouse()
	newHouse.UserID = parsedClaim.ID
	result, err2 := h.houseUsecase.AddHouse(newHouse, reqNewHouse.Photos)
	if err2 != nil {
		errorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	JSONResponse(c, http.StatusCreated, result)
}

func (h *Handler) HostUpdateHouse(c *gin.Context) {
	claim, _ := c.Get("claim")
	parsedClaim := entity.Claim(claim.(entity.Claim))

	id, errParse := strconv.Atoi(c.Param("id"))
	if errParse != nil {
		errorResponse(c, http.StatusBadRequest, errParse.Error())
	}

	var reqUpdateHouse dto.UpdateHouse
	err := c.ShouldBind(&reqUpdateHouse)
	if err != nil {
		errorTag(c, err)
		return
	}

	updateHouse := reqUpdateHouse.ToHouse()
	updateHouse.UserID = parsedClaim.ID
	updateHouse.ID = uint(id)
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

	claim, _ := c.Get("claim")
	parsedClaim := entity.Claim(claim.(entity.Claim))

	result, err := h.houseUsecase.DeleteHouse(uint(id), uint(parsedClaim.ID))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	JSONResponse(c, http.StatusOK, result)
}

func (h *Handler) HostAddPhotoHouse(c *gin.Context) {

}
