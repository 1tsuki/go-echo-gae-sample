package service

import (
	"github.com/1tsuki/gl-market/domain/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/1tsuki/gl-market/domain/model/auth"
	"github.com/1tsuki/gl-market/domain/model/user"
)

type (
	AuthService struct {
		authInfoRepository repository.AuthInfoRepository
	}

	UserDescriptor struct {
		UserID user.UserID
		Email  string
		Role   user.Role
	}

	JwtCustomClaims struct {
		Email string
		Role  user.Role
		jwt.StandardClaims
	}
)

func NewAuthService(authInfoRepository repository.AuthInfoRepository) *AuthService {
	a := new(AuthService)
	a.authInfoRepository = authInfoRepository

	return a
}

func NewUserDescriptorFromAuthInfo(authInfo *auth.AuthInfo) *UserDescriptor {
	u := new(UserDescriptor)
	u.UserID = authInfo.UserID
	u.Email = authInfo.Email
	u.Role = authInfo.Role

	return u
}

func (a *AuthService) AuthenticateUser(email string, password string) (*UserDescriptor, error) {
	authInfo, err := a.authInfoRepository.Get(email)
	if err != nil {
		return nil, err
	}

	err = CompareHashAndPassword(authInfo.Password, password)
	if err != nil {
		return nil, err
	}

	return NewUserDescriptorFromAuthInfo(authInfo), nil
}

/**
 * issue JWT with UserDescriptor
 * TODO: make Token's available period to properties
 * TODO: change SignedString to a key
 */
func (a *AuthService) IssueJWT(descriptor *UserDescriptor) (string, error) {
	claims := &JwtCustomClaims{
		descriptor.Email,
		descriptor.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret"))
}

func (a *AuthService) RegisterUser(email string, password string, Role user.Role) (*UserDescriptor, error) {
	bytes, err := EncryptedValue(password)
	if err != nil {
		return nil, err
	}

	authInfo := auth.AuthInfo{"", email, Role, string(bytes)}
	a.authInfoRepository.Create(authInfo)

	return NewUserDescriptorFromAuthInfo(&authInfo), nil
}
