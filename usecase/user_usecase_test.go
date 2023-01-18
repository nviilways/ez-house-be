package usecase_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	mocks "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/mocks/repository"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/usecase"
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

// func TestSignIn(t *testing.T) {
// 	t.Run("should return user credential when called", func(t *testing.T) {
// 		user := &enti
// 	})
// }