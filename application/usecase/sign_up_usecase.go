package usecase

import (
	"github.com/1tsuki/gl-market/domain/repository"
	"github.com/1tsuki/gl-market/domain/service"
	"github.com/1tsuki/gl-market/domain/model/user"
)

type signUpUsecase struct {
	authInfoRepository repository.AuthInfoRepository
}

func NewSignUpUsecase(authInfoRepository repository.AuthInfoRepository) *signUpUsecase {
	s := new(signUpUsecase)
	s.authInfoRepository = authInfoRepository
	return s
}

func (u *signUpUsecase) SignUp(email string, password string) (*service.UserDescriptor, error) {
	authService := service.NewAuthService(u.authInfoRepository)
	return authService.RegisterUser(email, password, user.UserRole)
}
