package usecase

import (
	"github.com/1tsuki/gl-market/domain/repository"
	"github.com/1tsuki/gl-market/domain/service"
)

type signInUsecase struct {
	authInfoRepository repository.AuthInfoRepository
}

func NewSignInUsecase(authInfoRepository repository.AuthInfoRepository) *signInUsecase {
	s := new(signInUsecase)
	s.authInfoRepository = authInfoRepository
	return s
}

func (u *signInUsecase) SignIn(email string, password string) (string, error) {
	authService := service.NewAuthService(u.authInfoRepository)
	descriptor, err := authService.AuthenticateUser(email, password)
	if err != nil {
		return "", err
	}

	return authService.IssueJWT(descriptor)
}
