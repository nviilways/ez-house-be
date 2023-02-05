package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/handler"
	mocks "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/mocks/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserGetHouseByVacancy(t *testing.T) {
	t.Run("should return list of vacant house when called", func(t *testing.T) {
		var houses []*entity.House
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
		inDate, _ := time.Parse("2006-01-02", "0001-01-01")
		outDate, _ := time.Parse("2006-01-02", "0001-01-01")
		filter := &dto.FilterHouse{
			CheckInDate: inDate,
			CheckOutDate: outDate,
			SortColumn: "name",
			SortBy: "asc",
			SearchName: "",
			SearchCity: "",
			SearchGuest: 1,
		}
		result := dto.JSONResponse{
			Code: 200,
			Message: "OK",
			Data: pagination,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.HouseUsecase)
		mockUsecase.On("GetHouseListByVacancy", filter, pagination).Return(houses, 0, nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			HouseUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("GET", "/api/v1/houses", nil)

		handler.UserGetHouseByVacancy(c)

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
		inDate, _ := time.Parse("2006-01-02", "0001-01-01")
		outDate, _ := time.Parse("2006-01-02", "0001-01-01")
		filter := &dto.FilterHouse{
			CheckInDate: inDate,
			CheckOutDate: outDate,
			SortColumn: "name",
			SortBy: "asc",
			SearchName: "",
			SearchCity: "",
			SearchGuest: 1,
		}
		result := dto.JSONResponse{
			Code: 500,
			Message: errs.ErrorCode[500],
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.HouseUsecase)
		mockUsecase.On("GetHouseListByVacancy", filter, pagination).Return(nil, 0, errors.New("error"))
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			HouseUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("GET", "/api/v1/houses", nil)

		handler.UserGetHouseByVacancy(c)

		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})
}