package service

import (
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/jwt"
)

func (authService *AuthServiceImpl) GetAccessToken(tokenRequest web.AccessTokenRequest) (*web.AuthResponse, error) {
	err := authService.Validate.Struct(tokenRequest)

	if err != nil {
		return nil, err
	}

	data, err := jwt.ExtractToken(tokenRequest.RefreshToken)

	if err != nil {
		return nil, err
	}

	userData := &web.AuthResponse{
		ID:       data.ID,
		Username: data.Username,
		Email:    data.Email,
		RoleName: data.RoleName,
	}

	return userData, nil
}
