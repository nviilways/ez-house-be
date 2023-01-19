package handler

import (
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UserRegister(c *gin.Context) {
	var reqRegister dto.UserRegister
	err := c.ShouldBindJSON(&reqRegister)
	
	if(err != nil) {
		errorTag(c, err)
		return
	}

	newRegister := reqRegister.ToUser()
	newUser, err2 := h.userUsecase.SignUp(newRegister)
	if(err2 != nil) {
		if errors.Is(err2, errs.ErrDuplicateEntry){
			errorResponse(c, http.StatusConflict, err2.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	c.JSON(http.StatusCreated, newUser)
}