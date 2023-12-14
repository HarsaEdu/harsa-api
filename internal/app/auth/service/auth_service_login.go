package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionResponse "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (authService *AuthServiceImpl) LoginUser(userRequest web.LoginUserRequest) (*web.AuthResponse, error) {
	// check available username and email
	existingUser, _ := authService.UserRepository.UserAvailable("", userRequest.Email)

	if existingUser == nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	err :=  authService.Password.ComparePassword(existingUser.Password, userRequest.Password)

	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	response, err := authService.AuthRepository.LoginUser(existingUser.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to get data user: %s", err.Error())
	}

	// convert user data to auth response
	userResponse := conversionResponse.AuthDomainToAuthResponse(response)

	return userResponse, nil
}
