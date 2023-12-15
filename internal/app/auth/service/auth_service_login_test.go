package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
	// Create mock repositories
	mockAuthRepo := new(mocks.AuthRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockPasswordComparer := new(mocks.PasswordComparer)
	mockFirebase := new(mocks.Firebase)
	mockNotification := new(mocks.NotificationRepository)
	validate := validator.New()


	// Create an AuthServiceImpl with the mock repositories
	authService := NewAuthService(mockAuthRepo, mockUserRepo, validate, mockPasswordComparer,mockNotification, mockFirebase )

	// Define test data
	userRequest := web.LoginUserRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Set up mock expectations
	mockUserRepo.On("UserAvailable", "", userRequest.Email).Return(&domain.User{ID: 1, Password: "hashedPassword"}, nil)
	mockAuthRepo.On("LoginUser", uint(1)).Return(&domain.Auth{}, nil)
	mockPasswordComparer.On("ComparePassword", "hashedPassword", userRequest.Password).Return(nil)

	// Call the function you want to test
	result, err := authService.LoginUser(userRequest)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// Add assertions for AuthResponse values

	// Assert that the expected calls were made
	mockUserRepo.AssertExpectations(t)
	mockAuthRepo.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

func TestComparWithError(t *testing.T) {
	// Create mock repositories
	mockAuthRepo := new(mocks.AuthRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockPasswordComparer := new(mocks.PasswordComparer)
	mockFirebase := new(mocks.Firebase)
	mockNotification := new(mocks.NotificationRepository)
	validate := validator.New()

	// Create an AuthServiceImpl with the mock repositories
	authService := NewAuthService(mockAuthRepo, mockUserRepo, validate, mockPasswordComparer,mockNotification, mockFirebase )

	// Define test data
	userRequest := web.LoginUserRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	user := domain.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Define an expected error during user retrieval
	expectedError := fmt.Errorf("error retrieving user")

	// Set up mock expectations
	mockUserRepo.On("UserAvailable", "", userRequest.Email).Return(&user, nil)
	mockPasswordComparer.On("ComparePassword", "password123", userRequest.Password).Return(expectedError)

	// Call the function you want to test
	result, err := authService.LoginUser(userRequest)

	// Assert that the result is nil (as there's an error)
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockUserRepo.AssertExpectations(t)
	mockAuthRepo.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

func TestLoginUserWithError(t *testing.T) {
	// Create mock repositories
	mockAuthRepo := new(mocks.AuthRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockPasswordComparer := new(mocks.PasswordComparer)
	mockFirebase := new(mocks.Firebase)
	mockNotification := new(mocks.NotificationRepository)
	validate := validator.New()

	// Create an AuthServiceImpl with the mock repositories
	authService := NewAuthService(mockAuthRepo, mockUserRepo, validate, mockPasswordComparer,mockNotification, mockFirebase )

	// Define test data
	userRequest := web.LoginUserRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Define an expected error during user retrieval
	expectedError := fmt.Errorf("error retrieving user")

	// Set up mock expectations
	mockUserRepo.On("UserAvailable", "", userRequest.Email).Return(nil, expectedError)

	// Call the function you want to test
	result, err := authService.LoginUser(userRequest)

	// Assert that the result is nil (as there's an error)
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockUserRepo.AssertExpectations(t)
	mockAuthRepo.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

