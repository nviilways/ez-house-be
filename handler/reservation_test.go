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

func TestUserAddReservation(t *testing.T) {
	t.Run("should return created reservation when successful", func(t *testing.T) {
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		reserve := dto.NewReservation{
			HouseId: 1,
			CheckInDate: "2006-02-01",
			CheckOutDate: "2006-02-01",
		}
		toReserve := reserve.ToReservation()
		toReserve.UserId = claim.ID
		result := dto.JSONResponse{
			Code: 201,
			Message: "CREATED",
			Data: toReserve,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.ReservationUsecase)
		mockUsecase.On("AddReservation", toReserve).Return(toReserve, nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			ReservationUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("POST", "/api/v1/reservations", testutils.MakeRequestBody(reserve))

		handler.UserAddReservation(c)

		assert.Equal(t, 201, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return bad request when required input doenst met", func(t *testing.T) {
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		reserve := dto.NewReservation{
			HouseId: 1,
			CheckInDate: "2006-02-01",
		}
		toReserve := reserve.ToReservation()
		toReserve.UserId = claim.ID
		result := dto.JSONResponse{
			Code: 400,
			Message: errs.ErrorCode[400],
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.ReservationUsecase)
		mockUsecase.On("AddReservation", toReserve).Return(toReserve, nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			ReservationUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("POST", "/api/v1/reservations", testutils.MakeRequestBody(reserve))

		handler.UserAddReservation(c)

		assert.Equal(t, 400, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return bad request when duplicate entry", func(t *testing.T) {
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		reserve := dto.NewReservation{
			HouseId: 1,
			CheckInDate: "2006-02-01",
			CheckOutDate: "2006-02-01",
		}
		toReserve := reserve.ToReservation()
		toReserve.UserId = claim.ID
		result := dto.JSONResponse{
			Code: 400,
			Message: errs.ErrorCode[400],
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.ReservationUsecase)
		mockUsecase.On("AddReservation", toReserve).Return(nil, errs.ErrDuplicateEntry)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			ReservationUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("POST", "/api/v1/reservations", testutils.MakeRequestBody(reserve))

		handler.UserAddReservation(c)

		assert.Equal(t, 400, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return error when internal server error", func(t *testing.T) {
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		reserve := dto.NewReservation{
			HouseId: 1,
			CheckInDate: "2006-02-01",
			CheckOutDate: "2006-02-01",
		}
		toReserve := reserve.ToReservation()
		toReserve.UserId = claim.ID
		result := dto.JSONResponse{
			Code: 500,
			Message: errs.ErrorCode[500],
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.ReservationUsecase)
		mockUsecase.On("AddReservation", toReserve).Return(nil, errors.New("error"))
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			ReservationUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("POST", "/api/v1/reservations", testutils.MakeRequestBody(reserve))

		handler.UserAddReservation(c)

		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})
}

func TestUserGetReservationById(t *testing.T) {
	t.Run("should return selected reservation when called", func(t *testing.T) {
		reserve := &entity.Reservation{
			ID: 1,
		}
		id := uint(1)
		result := dto.JSONResponse{
			Code: 200,
			Message: "OK",
			Data: reserve,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.ReservationUsecase)
		mockUsecase.On("GetReservationById", id).Return(reserve, nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Params = append(c.Params, gin.Param{
			Key: "id",
			Value: "1",
		})
		cfg := handler.Config{
			ReservationUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("GET", "/api/v1/reservations/1", nil)

		handler.UserGetReservationById(c)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return bad request when url param not valid", func(t *testing.T) {
		reserve := &entity.Reservation{
			ID: 1,
		}
		id := uint(1)
		result := dto.JSONResponse{
			Code: 400,
			Message: errs.ErrorCode[400],
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.ReservationUsecase)
		mockUsecase.On("GetReservationById", id).Return(reserve, nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		cfg := handler.Config{
			ReservationUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("GET", "/api/v1/reservations/1", nil)

		handler.UserGetReservationById(c)

		assert.Equal(t, 400, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return bad request when id is invalid", func(t *testing.T) {
		id := uint(1)
		result := dto.JSONResponse{
			Code: 400,
			Message: errs.ErrorCode[400],
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.ReservationUsecase)
		mockUsecase.On("GetReservationById", id).Return(nil, errs.ErrRecordNotFound)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Params = append(c.Params, gin.Param{
			Key: "id",
			Value: "1",
		})
		cfg := handler.Config{
			ReservationUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("GET", "/api/v1/reservations/1", nil)

		handler.UserGetReservationById(c)

		assert.Equal(t, 400, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should error when internal server error", func(t *testing.T) {
		id := uint(1)
		result := dto.JSONResponse{
			Code: 500,
			Message: errs.ErrorCode[500],
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.ReservationUsecase)
		mockUsecase.On("GetReservationById", id).Return(nil, errors.New("error"))
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Params = append(c.Params, gin.Param{
			Key: "id",
			Value: "1",
		})
		cfg := handler.Config{
			ReservationUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("GET", "/api/v1/reservations/1", nil)

		handler.UserGetReservationById(c)

		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})
}

func TestUserGetReservationByUserId(t *testing.T) {
	t.Run("should return user booking history when called", func(t *testing.T) {
		var reservation []*entity.Reservation
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
			Code: 200,
			Message: "OK",
			Data: pagination,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.ReservationUsecase)
		mockUsecase.On("GetReservationListByUserId", claim.ID, pagination).Return(reservation, 0, nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			ReservationUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("GET", "/api/v1/reservations/history", nil)

		handler.UserGetReservationByUserId(c)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return error when internal server error", func(t *testing.T) {
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
		mockUsecase := new(mocks.ReservationUsecase)
		mockUsecase.On("GetReservationListByUserId", claim.ID, pagination).Return(nil, 0, errors.New("error"))
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			ReservationUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("GET", "/api/v1/reservations/history", nil)

		handler.UserGetReservationByUserId(c)

		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})
}