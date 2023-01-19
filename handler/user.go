package handler

import (
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/config"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
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
	newRegister.RoleID = uint(2)

	isAdmin := c.GetHeader("SecretAdmin")
	if(isAdmin == config.AdminKey) {
		newRegister.RoleID = 1
	}
	newUser, err2 := h.userUsecase.SignUp(newRegister)
	if(err2 != nil) {
		if errors.Is(err2, errs.ErrDuplicateEntry){
			errorResponse(c, http.StatusConflict, err2.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	JSONResponse(c, http.StatusCreated, newUser)
}

func (h *Handler) UserLogin(c *gin.Context) {
	var reqLogin dto.UserLogin
	err := c.ShouldBindJSON(&reqLogin)

	if(err != nil) {
		errorTag(c, err)
		return
	}

	newLogin := reqLogin.ToUser()
	result, err2 := h.userUsecase.SignIn(newLogin.Password, newLogin)
	if(err2 != nil) {
		errorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	JSONResponse(c, http.StatusOK, result)
}

func (h *Handler) UserDetails(c *gin.Context) {
	claim, _ := c.Get("claim")
	parsedClaim := entity.Claim(claim.(entity.Claim))

	result, err := h.userUsecase.GetUserByID(parsedClaim.ID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	JSONResponse(c, http.StatusOK, result)
}