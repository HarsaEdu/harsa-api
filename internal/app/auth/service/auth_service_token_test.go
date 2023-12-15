package service

import (
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestGetAccessToken(t *testing.T) {
	// Create an AuthServiceImpl with a real validator
	authService := NewAuthService(nil, nil, validator.New(), nil,nil,nil)

	// Define test data
	tokenRequest := web.AccessTokenRequest{
		RefreshToken: "validRefreshToken",
	}

	// Call the function you want to test
	result, err := authService.GetAccessToken(tokenRequest)

	assert.Error(t, err)
	assert.Nil(t, result)
	// Add assertions for AuthResponse values

	// Additional assertions if needed
}

