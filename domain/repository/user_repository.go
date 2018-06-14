package repository

import "github.com/1tsuki/gl-market/domain/model/user"

type UserRepository interface {
	Get(userId int) (*user.User, error)
	GetAll(limit int) ([]*user.User, error)
	Save(*user.User) error
}
