package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/repository"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/utils"
)

type AdminUsecase interface {
	SignIn(string, *entity.Admin) (*dto.Token, error)
	SignUp(*entity.Admin) (*entity.Admin, error)
}

type adminUsecaseImpl struct {
	adminRepository repository.AdminRepository
}

type AdminUConfig struct {
	AdminRepository repository.AdminRepository
}

func NewAdminUsecase(cfg *AdminUConfig) AdminUsecase {
	return &adminUsecaseImpl {
		adminRepository: cfg.AdminRepository,
	}
}

func (a *adminUsecaseImpl) SignIn(pw string, admin *entity.Admin) (*dto.Token, error) {
	result, err := a.adminRepository.SignIn(admin)
	if err != nil {
		return nil, err
	}

	isValid := utils.ComparePassword(result.Password, pw)
	if isValid {
		toUser := &entity.User{
			ID: result.ID,
			Email: result.Email,
			RoleID: result.RoleID,
		}
		result, _ := utils.GenerateAccessToken(toUser)
		return result, nil
	}

	return nil, errs.ErrInvalidCredential
}

func (a *adminUsecaseImpl) SignUp(admin *entity.Admin) (*entity.Admin, error) {
	hashedPw, _ := utils.HashAndSalt(admin.Password)

	admin.Password = hashedPw

	result, err := a.adminRepository.SignUp(admin)
	if err != nil {
		return nil, err
	}

	return result, nil
}