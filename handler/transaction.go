package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UserTopUp(c *gin.Context) {
	claim, _ := c.Get("claim")
	parsedClaim := entity.Claim(claim.(entity.Claim))

	var reqTopUp dto.TopUp
	err := c.ShouldBindJSON(&reqTopUp)
	if err != nil {
		errorTag(c, err)
		return
	}

	newTopUp := reqTopUp.ToTransaction();
	result, err2 := h.transactionUsecase.TopUp(int(parsedClaim.ID), newTopUp)
	if err2 != nil {
		errorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	JSONResponse(c, http.StatusOK, result)
}

func (h *Handler) UserGetTransaction(c *gin.Context) {
	claim, _ := c.Get("claim")
	parsedClaim := entity.Claim(claim.(entity.Claim))

	result, err := h.transactionUsecase.GetTransactionByWalletId(parsedClaim.WalletID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	JSONResponse(c, http.StatusOK, result)
}