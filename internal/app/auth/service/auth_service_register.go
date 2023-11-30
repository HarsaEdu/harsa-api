package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	conversionResponse "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
	"github.com/HarsaEdu/harsa-api/internal/pkg/password"
	"github.com/labstack/echo/v4"
)

func (authService *AuthServiceImpl) RegisterUser(ctx echo.Context, r web.RegisterUserRequest) (*web.AuthResponse, error) {
	// validate the request
	err := authService.Validate.Struct(r)

	// check errors when validate the request
	if err != nil {
		return nil, err
	}

	// check available username and email
	existingUser, _ := authService.UserRepository.UserAvailable(r.Username, r.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("already exist")
	}

	// convert request to model
	user := conversionRequest.RegisterUserRequestToUserModel(r)

	// hash password
	user.Password = password.HashPassword(user.Password)

	// insert data and get back user data with id and role name
	res, err := authService.AuthRepository.RegisterUser(user)

	// check if error when insert data
	if err != nil {
		return nil, fmt.Errorf("error when creating user %s:", err.Error())
	}

	// convert user data to auth response
	userResponse := conversionResponse.AuthDomainToAuthResponse(res)

	return userResponse, nil
}