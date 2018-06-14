package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/1tsuki/gl-market/domain/model/user"
)

type AuthInfo struct {
	UserID   user.UserID
	Email    string `json:"email" datastore:"email" validate:"required"`
	Role     user.Role
	Password string `json:"password" datastore:"password" validate:"required"`
}

func (a AuthInfo) IssueJWTToken() (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["UserID"] = a.UserID
	claims["Email"] = a.Email
	claims["Role"] = a.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	return token.SignedString([]byte("secret")) // TODO: replace with secure key
}
