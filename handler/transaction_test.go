package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/handler"
	mocks "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/mocks/usecase"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/testutils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserTopUp(t *testing.T) {
	t.Run("should return the transaction after top up", func(t *testing.T) {
		tx := &entity.Transaction{
			ID: 1,
			Balance: 100000,
			TransactionTypeID: 1,
			WalletID: 1,
		}
		dtoTx := dto.TopUp {
			Amount: 100000,
		}
		toTx := dtoTx.ToTransaction()
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		result := dto.JSONResponse{
			Code: 200,
			Message: "OK",
			Data: tx,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.TransactionUsecase)
		mockUsecase.On("TopUp", int(claim.ID), toTx).Return(tx, nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			TransactionUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("POST", "/api/v1/transactions/topup", testutils.MakeRequestBody(dtoTx))

		handler.UserTopUp(c)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return bad request when required input doesnt met", func(t *testing.T) {
		tx := &entity.Transaction{
			ID: 1,
			Balance: 100000,
			TransactionTypeID: 1,
			WalletID: 1,
		}
		dtoTx := dto.TopUp {
			Amount: 1000,
		}
		toTx := dtoTx.ToTransaction()
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		result := dto.JSONResponse{
			Code: 400,
			Message: errs.ErrorCode[400],
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.TransactionUsecase)
		mockUsecase.On("TopUp", int(claim.ID), toTx).Return(tx, nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			TransactionUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("POST", "/api/v1/transactions/topup", testutils.MakeRequestBody(dtoTx))

		handler.UserTopUp(c)

		assert.Equal(t, 400, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return error when internal server error", func(t *testing.T) {
		dtoTx := dto.TopUp {
			Amount: 100000,
		}
		toTx := dtoTx.ToTransaction()
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		result := dto.JSONResponse{
			Code: 500,
			Message: errs.ErrorCode[500],
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.TransactionUsecase)
		mockUsecase.On("TopUp", int(claim.ID), toTx).Return(nil, errors.New("error"))
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			TransactionUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("POST", "/api/v1/transactions/topup", testutils.MakeRequestBody(dtoTx))

		handler.UserTopUp(c)

		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})
}

func TestUserGetTransaction(t *testing.T) {
	t.Run("should return user transaction history when called", func(t *testing.T) {
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		var transaction []*entity.Transaction
		pagination := &dto.Pagination{
			Page: 1,
			Limit: 10,
			Data: nil,
		}
		result := dto.JSONResponse{
			Code: 200,
			Message: "OK",
			Data: pagination,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.TransactionUsecase)
		mockUsecase.On("GetTransactionByWalletId", claim.WalletID, pagination).Return(transaction, 0, nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			TransactionUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("GET", "/api/v1/transactions", nil)

		handler.UserGetTransaction(c)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return user transaction history when called", func(t *testing.T) {
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		pagination := &dto.Pagination{
			Page: 1,
			Limit: 10,
			Data: nil,
		}
		result := dto.JSONResponse{
			Code: 500,
			Message: errs.ErrorCode[500],
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.TransactionUsecase)
		mockUsecase.On("GetTransactionByWalletId", claim.WalletID, pagination).Return(nil, 0, errors.New("error"))
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			TransactionUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("GET", "/api/v1/transactions", nil)

		handler.UserGetTransaction(c)

		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})
}