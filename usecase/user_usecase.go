package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/repository"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/utils"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
)

type UserUsecase interface {
	SignIn(string, *entity.User) (*entity.User, error)
	SignUp(*entity.User) (*entity.User, error)
	GetUserByID(*entity.User) (*entity.User, error)
}

type userUsecaseImpl struct {
	userRepository repository.UserRepository
}

type UserUConfig struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(cfg *UserUConfig) UserUsecase {
	return &userUsecaseImpl {
		userRepository: cfg.UserRepository,
	}
}

func (u *userUsecaseImpl) SignIn(pw string, user *entity.User) (*entity.User, error) {
	result, err := u.userRepository.SignIn(user)

	if(err != nil) {
		return nil, err
	}

	isValid := utils.ComparePassword(result.Password, pw)
	if(isValid) {
		return result, nil
	}

	return nil, errs.ErrInvalidCredential
}

func (u *userUsecaseImpl) SignUp(user *entity.User) (*entity.User, error) {
	hashedPw, _ := utils.HashAndSalt(user.Password)

	user.Password = hashedPw

	result, err := u.userRepository.SignUp(user)

	if(err != nil) {
		return nil, err
	}

	return result, nil
}

func (u *userUsecaseImpl) GetUserByID(user *entity.User) (*entity.User, error) {
	result, err := u.userRepository.GetUserByID(user)

	if(err != nil) {
		return nil, err
	}

	return result, nil
}