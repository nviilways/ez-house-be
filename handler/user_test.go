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
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/server"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/testutils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserRegister(t *testing.T) {
	t.Run("should return registered user when successful", func(t *testing.T) {
		newUser := &entity.User{
			Email: "user@mail.com",
			Password: "test",
			CityID: 1,
			RoleID: 1,
		}
		result := &dto.JSONResponse{
			Code: 201,
			Message: "CREATED",
			Data: newUser,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("SignUp", newUser).Return(newUser, nil)
		cfg := &server.RouterConfig{
			UserUsecase: mockUsecase,
		}
		req, _ := http.NewRequest("POST", "/api/v1/register", testutils.MakeRequestBody(newUser))

		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, 201, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return bad request when required input doesnt met", func(t *testing.T) {
		newUser := &entity.User{
			Email: "user@mail.com",
			Password: "",
			CityID: 1,
			RoleID: 1,
		}
		result := &dto.JSONResponse{
			Code: 400,
			Message: errs.ErrorCode[400],
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("SignUp", newUser).Return(nil, errors.New("error"))
		cfg := &server.RouterConfig{
			UserUsecase: mockUsecase,
		}
		req, _ := http.NewRequest("POST", "/api/v1/register", testutils.MakeRequestBody(newUser))

		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, 400, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return error when duplicate entry", func(t *testing.T) {
		newUser := &entity.User{
			Email: "user@mail.com",
			Password: "test",
			CityID: 1,
			RoleID: 1,
		}
		result := &dto.JSONResponse{
			Code: 409,
			Message: errs.ErrDuplicateEntry.Error(),
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("SignUp", newUser).Return(nil, errs.ErrDuplicateEntry)
		cfg := &server.RouterConfig{
			UserUsecase: mockUsecase,
		}
		http.NewRequest("POST", "/api/v1/register", testutils.MakeRequestBody(newUser))
		req, _ := http.NewRequest("POST", "/api/v1/register", testutils.MakeRequestBody(newUser))

		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, 409, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return error when internal server error", func(t *testing.T) {
		newUser := &entity.User{
			Email: "user@mail.com",
			Password: "test",
			CityID: 1,
			RoleID: 1,
		}
		result := &dto.JSONResponse{
			Code: 500,
			Message: "error",
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("SignUp", newUser).Return(nil, errors.New("error"))
		cfg := &server.RouterConfig{
			UserUsecase: mockUsecase,
		}
		req, _ := http.NewRequest("POST", "/api/v1/register", testutils.MakeRequestBody(newUser))

		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})
}

func TestUserLogin(t *testing.T) {
	t.Run("should return login token when successfully logged in", func(t *testing.T) {
		user := &entity.User{
			Email: "user@mail.com",
			Password: "test",
		}
		token := &dto.Token{}
		result := &dto.JSONResponse{
			Code: 200,
			Message: "OK",
			Data: token,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("SignIn", user.Password, user).Return(token, nil)
		cfg := &server.RouterConfig{
			UserUsecase: mockUsecase,
		}
		req, _ := http.NewRequest("POST", "/api/v1/login", testutils.MakeRequestBody(user))

		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})
	t.Run("should return bad request when required input doesnt met", func(t *testing.T) {
		user := &entity.User{
			Email: "user@mail.com",
			Password: "",
		}
		result := &dto.JSONResponse{
			Code: 400,
			Message: errs.ErrorCode[400],
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("SignIn", user.Password, user).Return(nil, errors.New("error"))
		cfg := &server.RouterConfig{
			UserUsecase: mockUsecase,
		}
		req, _ := http.NewRequest("POST", "/api/v1/login", testutils.MakeRequestBody(user))

		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, 400, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})
	t.Run("should return bad request when credential invalid", func(t *testing.T) {
		user := &entity.User{
			Email: "user@mail.com",
			Password: "test",
		}
		result := &dto.JSONResponse{
			Code: 400,
			Message: errs.ErrInvalidCredential.Error(),
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("SignIn", user.Password, user).Return(nil, errs.ErrInvalidCredential)
		cfg := &server.RouterConfig{
			UserUsecase: mockUsecase,
		}
		req, _ := http.NewRequest("POST", "/api/v1/login", testutils.MakeRequestBody(user))

		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, 400, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})
	t.Run("should return error when internal server error", func(t *testing.T) {
		user := &entity.User{
			Email: "user@mail.com",
			Password: "test",
		}
		result := &dto.JSONResponse{
			Code: 500,
			Message: "error",
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("SignIn", user.Password, user).Return(nil, errors.New("error"))
		cfg := &server.RouterConfig{
			UserUsecase: mockUsecase,
		}
		req, _ := http.NewRequest("POST", "/api/v1/login", testutils.MakeRequestBody(user))

		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})
}

func TestUserLogout(t *testing.T) {
	t.Run("should return no content when logged out", func(t *testing.T) {
		token := ""
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("SignOut", token).Return(nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			UserUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("POST", "/api/v1/logout", nil)

		handler.UserLogout(c)

		assert.Equal(t, 204, rec.Code)
	})

	t.Run("should return error when internal server error", func(t *testing.T) {
		token := ""
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("SignOut", token).Return(errors.New("error"))
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			UserUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("POST", "/api/v1/logout", nil)

		handler.UserLogout(c)

		assert.Equal(t, 500, rec.Code)
	})
}

func TestUserDetails(t *testing.T) {
	t.Run("should return user details when called", func(t *testing.T) {
		token := ""
		user := &entity.User{
			ID: 1,
		}
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		result := dto.JSONResponse{
			Code: 200,
			Message: "OK",
			Data: user,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("TokenCheck", token).Return(nil)
		mockUsecase.On("GetUserByID", claim.ID).Return(user, nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			UserUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("POST", "/api/v1/me", nil)

		handler.UserDetails(c)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return error when token invalid", func(t *testing.T) {
		token := ""
		user := &entity.User{
			ID: 1,
		}
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		result := dto.JSONResponse{
			Code: 401,
			Message: "error",
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("TokenCheck", token).Return(errors.New("error"))
		mockUsecase.On("GetUserByID", claim.ID).Return(user, nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			UserUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("POST", "/api/v1/me", nil)

		handler.UserDetails(c)

		assert.Equal(t, 401, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return error internal server error", func(t *testing.T) {
		token := ""
		user := &entity.User{
			ID: 1,
		}
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		result := dto.JSONResponse{
			Code: 500,
			Message: "error",
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("TokenCheck", token).Return(nil)
		mockUsecase.On("GetUserByID", claim.ID).Return(user, errors.New("error"))
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			UserUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("POST", "/api/v1/me", nil)

		handler.UserDetails(c)

		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})	
}

func TestUserUpdate(t *testing.T) {
	t.Run("should return updated user data when successful", func(t *testing.T) {
		token := ""
		user := &entity.User{
			ID: 1,
			FullName: "Asep",
		}
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		dtoUser := dto.UserUpdate{
			FullName: "Asep",
		}
		result := dto.JSONResponse{
			Code: 200,
			Message: "OK",
			Data: dtoUser,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("TokenCheck", token).Return(nil)
		mockUsecase.On("Update", user).Return(user, nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		c.Params = append(c.Params, gin.Param{
			Key: "id",
			Value: "1",
		})
		cfg := handler.Config{
			UserUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("PATCH", "/api/v1/users/1", testutils.MakeRequestBody(user))

		handler.UserUpdate(c)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return unauthorized when claim id not match with param id", func(t *testing.T) {
		token := ""
		user := &entity.User{
			ID: 1,
		}
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		result := dto.JSONResponse{
			Code: 401,
			Message: "invalid token",
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("TokenCheck", token).Return(nil)
		mockUsecase.On("Update", user).Return(user, nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		c.Params = append(c.Params, gin.Param{
			Key: "id",
			Value: "2",
		})
		cfg := handler.Config{
			UserUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("PATCH", "/api/v1/users/1", testutils.MakeRequestBody(user))

		handler.UserUpdate(c)

		assert.Equal(t, 401, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return unauthorized when token check error", func(t *testing.T) {
		token := ""
		user := &entity.User{
			ID: 1,
		}
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		result := dto.JSONResponse{
			Code: 401,
			Message: "error",
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("TokenCheck", token).Return(errors.New("error"))
		mockUsecase.On("Update", user).Return(user, nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		c.Params = append(c.Params, gin.Param{
			Key: "id",
			Value: "2",
		})
		cfg := handler.Config{
			UserUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("PATCH", "/api/v1/users/1", testutils.MakeRequestBody(user))

		handler.UserUpdate(c)

		assert.Equal(t, 401, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return error when internal server error", func(t *testing.T) {
		token := ""
		user := &entity.User{
			ID: 1,
		}
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		result := dto.JSONResponse{
			Code: 500,
			Message: "error",
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("TokenCheck", token).Return(nil)
		mockUsecase.On("Update", user).Return(user, errors.New("error"))
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		c.Params = append(c.Params, gin.Param{
			Key: "id",
			Value: "1",
		})
		cfg := handler.Config{
			UserUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("PATCH", "/api/v1/users/1", testutils.MakeRequestBody(user))

		handler.UserUpdate(c)

		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})
}

func TestUserUpdateRole(t *testing.T) {
	t.Run("should return no content when change role success", func(t *testing.T) {
		token := ""
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("TokenCheck", token).Return(nil)
		mockUsecase.On("UpdateRole", claim.ID).Return(nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			UserUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("POST", "/api/v1/users/host", nil)

		handler.UserUpdateRole(c)

		assert.Equal(t, 204, rec.Code)
	})

	t.Run("should return unauthorized when token invalid", func(t *testing.T) {
		token := ""
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		result := dto.JSONResponse{
			Code: 401,
			Message: "error",
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("TokenCheck", token).Return(errors.New("error"))
		mockUsecase.On("UpdateRole", claim.ID).Return(nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			UserUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("POST", "/api/v1/users/host", nil)

		handler.UserUpdateRole(c)

		assert.Equal(t, 401, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return error when internal server error", func(t *testing.T) {
		token := ""
		claim := entity.Claim{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 1,
			WalletID: 1,
		}
		result := dto.JSONResponse{
			Code: 500,
			Message: "error",
			Data: nil,
		}
		expectedBody, _ := json.Marshal(result)
		mockUsecase := new(mocks.UserUsecase)
		mockUsecase.On("TokenCheck", token).Return(nil)
		mockUsecase.On("UpdateRole", claim.ID).Return(errors.New("error"))
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("claim", claim)
		cfg := handler.Config{
			UserUsecase: mockUsecase,
		}
		handler := handler.New(&cfg)
		c.Request, _ = http.NewRequest("POST", "/api/v1/users/host", nil)

		handler.UserUpdateRole(c)

		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})
}