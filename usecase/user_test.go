package usecase_test

import (
	"errors"
	"testing"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	mocks "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/mocks/repository"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/usecase"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/utils"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	t.Run("should return registered user credential when called", func(t *testing.T) {
		user := &entity.User {
			Email: "user@mail.com",
			Password: "user",
		}
		mockRepository := new(mocks.UserRepository)
		usecase := usecase.NewUserUsecase(&usecase.UserUConfig{
			UserRepository: mockRepository,
		})
		mockRepository.On("SignUp", user).Return(user, nil)

		result, _ := usecase.SignUp(user)

		assert.Equal(t, user, result);
	})

	t.Run("should return duplicate error when registering with same email", func(t *testing.T) {
		user := &entity.User {
			Email: "user@mail.com",
			Password: "user",
		}
		mockRepository := new(mocks.UserRepository)
		usecase := usecase.NewUserUsecase(&usecase.UserUConfig{
			UserRepository: mockRepository,
		})
		mockRepository.On("SignUp", user).Return(nil, errs.ErrDuplicateEntry)
		
		usecase.SignUp(user)
		_, err := usecase.SignUp(user)

		assert.ErrorIs(t, err, errs.ErrDuplicateEntry)
	})
}

func TestSignIn(t *testing.T) {
	t.Run("should return access token when called", func(t *testing.T) {
		pw, _ := utils.HashAndSalt("test")
		user := &entity.User{
			Email: "user@mail.com",
			Password: pw,
		}
		mockRepository := new(mocks.UserRepository)
		usecase := usecase.NewUserUsecase(&usecase.UserUConfig{
			UserRepository: mockRepository,
		})
		mockRepository.On("SignIn", user).Return(user, nil)

		result, _ := usecase.SignIn("test", user)

		assert.NotNil(t, result)
	})
	t.Run("should return error invalid credential when inserting wrong password or user not exist", func(t *testing.T) {
		pw, _ := utils.HashAndSalt("test")
		user := &entity.User{
			Email: "user@mail.com",
			Password: pw,
		}
		mockRepository := new(mocks.UserRepository)
		usecase := usecase.NewUserUsecase(&usecase.UserUConfig{
			UserRepository: mockRepository,
		})
		mockRepository.On("SignIn", user).Return(user, nil)

		_, err := usecase.SignIn("teste", user)

		assert.ErrorIs(t, err, errs.ErrInvalidCredential)
	})
	t.Run("should return error when server is down", func(t *testing.T) {
		pw, _ := utils.HashAndSalt("test")
		user := &entity.User{
			Email: "user@mail.com",
			Password: pw,
		}
		mockRepository := new(mocks.UserRepository)
		usecase := usecase.NewUserUsecase(&usecase.UserUConfig{
			UserRepository: mockRepository,
		})
		mockRepository.On("SignIn", user).Return(nil, errors.New("error"))
		expectedErr := errors.New("error")

		_, err := usecase.SignIn("test", user)

		assert.Error(t, err, expectedErr)
	})
}

func TestGetUserByID(t *testing.T) {
	t.Run("should return user details when called", func(t *testing.T) {
		user := &entity.User{
			ID: 1,
			Email: "user@mail.com",
		}
		mockRepository := new(mocks.UserRepository)
		usecase := usecase.NewUserUsecase(&usecase.UserUConfig{
			UserRepository: mockRepository,
		})
		mockRepository.On("GetUserByID", uint(1)).Return(user, nil)

		result, _ := usecase.GetUserByID(uint(1))

		assert.Equal(t, result, user);
	})
	t.Run("should return error when error occured", func(t *testing.T) {
		mockRepository := new(mocks.UserRepository)
		usecase := usecase.NewUserUsecase(&usecase.UserUConfig{
			UserRepository: mockRepository,
		})
		mockRepository.On("GetUserByID", uint(1)).Return(nil, errors.New("error"))
		expectedErr := errors.New("error")

		_, err := usecase.GetUserByID(uint(1))

		assert.Error(t, err, expectedErr)
	})
}