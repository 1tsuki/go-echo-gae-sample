package mock

import (
	"github.com/1tsuki/gl-market/domain/service"
	"github.com/1tsuki/gl-market/domain/model/auth"
	"github.com/1tsuki/gl-market/domain/model/user"
)

type AuthInfoRepository struct{}

func (r *AuthInfoRepository) Get(email string) (*auth.AuthInfo, error) {
	hashValue, err := service.EncryptedValue("password")
	if err != nil {
		return nil, err
	}

	return &auth.AuthInfo{UserID: "1234", Email: email, Role: user.UserRole, Password: hashValue}, nil
}

func (r *AuthInfoRepository) Create(auth.AuthInfo) error {
	return nil
}

func (r *AuthInfoRepository) Update() error {
	return nil
}
