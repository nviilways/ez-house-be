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

func TestAdminSignIn(t *testing.T) {
	t.Run("should return token when successfully logged in", func(t *testing.T) {
		hashedPw, _ := utils.HashAndSalt("password")
		admin := &entity.Admin{
			Email:    "user@mail.com",
			Password: hashedPw,
		}
		mockRepository := new(mocks.AdminRepository)
		usecase := usecase.NewAdminUsecase(&usecase.AdminUConfig{
			AdminRepository: mockRepository,
		})
		mockRepository.On("SignIn", admin).Return(admin, nil)

		result, _ := usecase.SignIn("password", admin)

		assert.NotNil(t, result)
	})

	t.Run("should return error when inputting invalid credential", func(t *testing.T) {
		hashedPw, _ := utils.HashAndSalt("passwo")
		admin := &entity.Admin{
			Email:    "user@mail.com",
			Password: hashedPw,
		}
		mockRepository := new(mocks.AdminRepository)
		usecase := usecase.NewAdminUsecase(&usecase.AdminUConfig{
			AdminRepository: mockRepository,
		})
		mockRepository.On("SignIn", admin).Return(admin, nil)

		_ , err := usecase.SignIn("password", admin)

		assert.Error(t, errs.ErrInvalidCredential, err)
	})

	t.Run("should return error when server error", func(t *testing.T) {
		hashedPw, _ := utils.HashAndSalt("password")
		admin := &entity.Admin{
			Email:    "user@mail.com",
			Password: hashedPw,
		}
		mockRepository := new(mocks.AdminRepository)
		usecase := usecase.NewAdminUsecase(&usecase.AdminUConfig{
			AdminRepository: mockRepository,
		})
		mockRepository.On("SignIn", admin).Return(nil, errors.New("error"))
		expectedErr := errors.New("error")

		_ , err := usecase.SignIn("password", admin)

		assert.Error(t, expectedErr, err)
	})
}

func TestAdminSignUp(t *testing.T) {
	t.Run("should return created admin information", func(t *testing.T) {
		admin := &entity.Admin{
			ID: 1,
			Email: "admin@min.com",
			Password: "password",
		}
		mockRepository := new(mocks.AdminRepository)
		usecase := usecase.NewAdminUsecase(&usecase.AdminUConfig{
			AdminRepository: mockRepository,
		})
		mockRepository.On("SignUp", admin).Return(admin, nil)

		result, _ := usecase.SignUp(admin)
		
		assert.Equal(t, admin, result)
	})

	t.Run("should return error when server error", func(t *testing.T) {
		admin := &entity.Admin{
			ID: 1,
			Email: "admin@min.com",
			Password: "password",
		}
		mockRepository := new(mocks.AdminRepository)
		usecase := usecase.NewAdminUsecase(&usecase.AdminUConfig{
			AdminRepository: mockRepository,
		})
		mockRepository.On("SignUp", admin).Return(nil, errors.New("error"))
		expectedErr := errors.New("error")

		_, err := usecase.SignUp(admin)
		
		assert.Error(t, expectedErr, err)
	})
}