package handler

import (
	"errors"
	"net/http"
	"strconv"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/config"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/utils"
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
		if(errors.Is(err2, errs.ErrInvalidCredential)) {
			errorResponse(c, http.StatusBadRequest, err2.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	JSONResponse(c, http.StatusOK, result)
}

func (h *Handler) UserLogout(c *gin.Context) {
	token := utils.GetToken(c)

	err := h.userUsecase.SignOut(token)
	if(err != nil) {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	JSONResponse(c, http.StatusNoContent, nil)
}

func (h *Handler) UserDetails(c *gin.Context) {
	token := utils.GetToken(c)
	tokenErr := h.userUsecase.TokenCheck(token)
	if tokenErr != nil {
		errorResponse(c, http.StatusUnauthorized, tokenErr.Error())
		return
	}
	
	claim, _ := c.Get("claim")
	parsedClaim := entity.Claim(claim.(entity.Claim))

	result, err := h.userUsecase.GetUserByID(parsedClaim.ID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	JSONResponse(c, http.StatusOK, result)
}

func (h *Handler) UserUpdate(c *gin.Context) {
	token := utils.GetToken(c)
	tokenErr := h.userUsecase.TokenCheck(token)
	if tokenErr != nil {
		errorResponse(c, http.StatusUnauthorized, tokenErr.Error())
		return
	}

	claim, _ := c.Get("claim")
	parsedClaim := entity.Claim(claim.(entity.Claim))
	var reqUpdate dto.UserUpdate
	err := c.ShouldBindJSON(&reqUpdate)
	if(err != nil) {
		errorTag(c, err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	if(id != int(parsedClaim.ID)) {
		errorResponse(c, http.StatusUnauthorized, errs.ErrInvalidToken.Error())
		return
	}

	newUpdate := reqUpdate.ToUser()
	newUpdate.ID = parsedClaim.ID

	result, err2 := h.userUsecase.Update(newUpdate)
	if(err2 != nil) {
		errorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	reqUpdate.FromUser(result)

	JSONResponse(c, http.StatusOK, reqUpdate)
}

func (h *Handler) UserUpdateRole(c *gin.Context) {
	token := utils.GetToken(c)
	tokenErr := h.userUsecase.TokenCheck(token)
	if tokenErr != nil {
		errorResponse(c, http.StatusUnauthorized, tokenErr.Error())
		return
	}

	claim, _ := c.Get("claim")
	parsedClaim := entity.Claim(claim.(entity.Claim))

	err := h.userUsecase.UpdateRole(parsedClaim.ID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	JSONResponse(c, http.StatusNoContent, nil)
}