package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/1tsuki/gl-market/domain/model/user"
)

type jwtCustomClaims struct {
	Email string    `json:"email"`
	Role  user.Role `json:"role"`
	jwt.StandardClaims
}
