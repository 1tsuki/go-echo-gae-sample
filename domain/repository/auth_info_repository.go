package repository

import (
	"github.com/1tsuki/gl-market/domain/model/auth"
)

type AuthInfoRepository interface {
	Get(email string) (*auth.AuthInfo, error)
	Create(model auth.AuthInfo) error
}
