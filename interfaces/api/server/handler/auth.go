package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/1tsuki/gl-market/infrastructure/persistence/mock"
	"github.com/1tsuki/gl-market/application/usecase"
	"github.com/1tsuki/gl-market/interfaces/api/server/response"
)

type SignUpInterface struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func SignUp(c echo.Context) error {
	// bind request
	signUpInterface := new(SignUpInterface)
	if err := c.Bind(signUpInterface); err != nil {
		return c.JSON(response.NewError(err, response.INVALID_ARGUMENT, nil))
	}

	// validation
	if err := c.Validate(signUpInterface); err != nil {
		return c.JSON(response.NewError(err, response.INVALID_ARGUMENT, nil))
	}

	// prepare repositories
	authInfoRepository := new(mock.AuthInfoRepository)

	// execute usecase
	signUpUsecase := usecase.NewSignUpUsecase(authInfoRepository)
	userDescriptor, err := signUpUsecase.SignUp(signUpInterface.Email, signUpInterface.Password)
	if err != nil {
		return c.JSON(response.NewError(err, response.UNAUTHENTICATED, nil))
	}

	// return value
	return c.JSON(http.StatusOK, map[string]string{
		"email":  userDescriptor.Email,
		"userID": string(userDescriptor.UserID),
	})
}

type SignInInterface struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func SignIn(c echo.Context) error {
	// bind request
	signUpInterface := new(SignUpInterface)
	if err := c.Bind(signUpInterface); err != nil {
		return c.JSON(response.NewError(err, response.INVALID_ARGUMENT, nil))
	}

	// validation
	if err := c.Validate(signUpInterface); err != nil {
		return c.JSON(response.NewError(err, response.INVALID_ARGUMENT, nil))
	}

	// prepare repositories
	authInfoRepository := new(mock.AuthInfoRepository)
	signInUsecase := usecase.NewSignInUsecase(authInfoRepository)
	jwtToken, err := signInUsecase.SignIn(signUpInterface.Email, signUpInterface.Password)
	if err != nil {
		return c.JSON(response.NewError(err, response.UNAUTHENTICATED, nil))
	}

	// return value
	return c.JSON(http.StatusOK, map[string]string{
		"token": jwtToken,
	})
}
