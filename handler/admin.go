package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AdminSignIn(c *gin.Context) {
	var reqLogin dto.UserLogin
	err := c.ShouldBindJSON(&reqLogin)
	if err != nil {
		errorTag(c, err)
		return
	}

	newAdminLogin := reqLogin.ToAdmin()
	result, err2 := h.adminUsecase.SignIn(newAdminLogin.Password, newAdminLogin)
	if err2 != nil {
		errorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	JSONResponse(c, http.StatusOK, result)
}

func (h *Handler) AdminSignUp(c *gin.Context) {
	var reqRegister dto.AdminRegister
	err := c.ShouldBindJSON(&reqRegister)
	if err != nil {
		errorTag(c, err)
		return
	}

	newAdminRegister := reqRegister.ToAdmin()
	result, err2 := h.adminUsecase.SignUp(newAdminRegister)
	if err2 != nil {
		errorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	JSONResponse(c, http.StatusCreated, result)
}