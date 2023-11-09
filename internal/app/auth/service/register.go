package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	conversionResponse "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
	"github.com/HarsaEdu/harsa-api/internal/pkg/password"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (s *AuthServiceImpl) RegisterUser(ctx echo.Context, r web.RegisterUserRequest) (*web.AuthResponse, error) {
	// validate the request
	err := s.Validate.Struct(r)

	// check errors when validate the request
	if err != nil {
		return nil, validation.ValidationError(ctx, err)
	}

	// check available username and email
	existingUser, _ := s.UserRepository.UserAvailable(r.Username, r.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("already exists")
	}

	// convert request to model
	user := conversionRequest.RegisterUserRequestToUserModel(r)

	// hash password
	user.Password = password.HashPassword(user.Password)

	// insert data and get back user data with id and role name
	res, err := s.AuthRepository.RegisterUser(user)

	// check if error when insert data
	if err != nil {
		return nil, fmt.Errorf("Error When Creating User %s:", err.Error())
	}

	// convert user data to auth response
	userResponse := conversionResponse.ConvertToAuthResponse(res)

	return userResponse, nil
}
