package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {
	// Create mock repositories and dependencies
	mockAuthRepo := new(mocks.AuthRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockPasswordComparer := new(mocks.PasswordComparer)
	mockFirebase := new(mocks.Firebase)
	mockNotification := new(mocks.NotificationRepository)
	validate := validator.New()

	// Create an AuthServiceImpl with the mock repositories
	authService := NewAuthService(mockAuthRepo, mockUserRepo, validate, mockPasswordComparer, mockNotification, mockFirebase)

	// Define test data
	userRequest := web.RegisterUserRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	// Mock expectations for successful user availability check
	mockUserRepo.On("UserAvailable", userRequest.Username, userRequest.Email).Return(nil, nil)

	// Mock expectations for successful password hashing
	mockPasswordComparer.On("HashPassword", userRequest.Password).Return("hashedPassword", nil)

	// Mock expectations for successful registration
	mockAuthRepo.On("RegisterWithFreeSubscibe", mock.AnythingOfType("*domain.User")).Return(&domain.Auth{
		// Populate with relevant data for AuthResponse
		ID:   uint(1),
		RoleName: "user",
		RegistrationToken: "sadas",
	}, nil)

	// Mock expectations for successful notification creation
	mockNotification.On("Create", mock.AnythingOfType("*domain.Notification")).Return(nil)

	// Mock expectations for Firebase notification
	mockFirebase.On("SendNotificationPersonal", mock.AnythingOfType("*web.NotificationPersonal")).Return(nil)

	// Call the function you want to test
	result, err := authService.RegisterUser(userRequest)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Add assertions for AuthResponse values
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, web.Role("user"), result.RoleName)
	// ... assert other AuthResponse fields

	// Assert that the expected calls were made
	mockUserRepo.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
	mockAuthRepo.AssertExpectations(t)
	mockNotification.AssertExpectations(t)
	mockFirebase.AssertExpectations(t)
}


func TestRegisterUserValidationError(t *testing.T) {
	// Create mock repositories and dependencies
	// Create mock repositories
	mockAuthRepo := new(mocks.AuthRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockPasswordComparer := new(mocks.PasswordComparer)
	mockFirebase := new(mocks.Firebase)
	mockNotification := new(mocks.NotificationRepository)
	validate := validator.New()

	// Create an AuthServiceImpl with the mock repositories
	authService := NewAuthService(mockAuthRepo, mockUserRepo, validate, mockPasswordComparer,mockNotification, mockFirebase )

	// Define test data with invalid input (simulate validation error)
	userRequest := web.RegisterUserRequest{
		// ... incomplete or invalid data
	}

	// Call the function you want to test
	result, err := authService.RegisterUser(userRequest)

	// Assert that the result is nil (as there's a validation error)
	require.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockUserRepo.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
	mockAuthRepo.AssertExpectations(t)
}

func TestRegisterUserExistingUser(t *testing.T) {
	// Create mock repositories and dependencies
	// Create mock repositories
	mockAuthRepo := new(mocks.AuthRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockPasswordComparer := new(mocks.PasswordComparer)
	mockFirebase := new(mocks.Firebase)
	mockNotification := new(mocks.NotificationRepository)
	validate := validator.New()

	// Create an AuthServiceImpl with the mock repositories
	authService := NewAuthService(mockAuthRepo, mockUserRepo, validate, mockPasswordComparer,mockNotification, mockFirebase )

	// Define test data with an existing user
	userRequest := web.RegisterUserRequest{
		Username: "existingUser",
		Email:    "existing@example.com",
		Password: "password123",
	}


	// Mock expectations for an existing user
	mockUserRepo.On("UserAvailable", userRequest.Username, userRequest.Email).Return(&domain.User{}, nil)

	// Call the function you want to test
	result, err := authService.RegisterUser(userRequest)

	// Assert that the result is nil (as there's an existing user)
	require.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockUserRepo.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
	mockAuthRepo.AssertExpectations(t)
}

func TestComparRegisterWithError(t *testing.T) {
	// Create mock repositories
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


	user := web.RegisterUserRequest{
		Username: "existingUser",
		Email:    "existing@example.com",
		Password: "password123",
	}

	// Define an expected error during user retrieval
	expectedError := fmt.Errorf("error retrieving user")

	// Set up mock expectations
	mockUserRepo.On("UserAvailable", user.Username, user.Email).Return(nil, nil)
	mockPasswordComparer.On("HashPassword", user.Password).Return("hashedPassword", expectedError)

	// Call the function you want to test
	result, err := authService.RegisterUser(user)

	// Assert that the result is nil (as there's an error)
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockUserRepo.AssertExpectations(t)
	mockAuthRepo.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

func TestRegisterUserNotificationError(t *testing.T) {
	// Create mock repositories and dependencies
	mockAuthRepo := new(mocks.AuthRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockPasswordComparer := new(mocks.PasswordComparer)
	mockFirebase := new(mocks.Firebase)
	mockNotification := new(mocks.NotificationRepository)
	validate := validator.New()

	// Create an AuthServiceImpl with the mock repositories
	authService := NewAuthService(mockAuthRepo, mockUserRepo, validate, mockPasswordComparer, mockNotification, mockFirebase)

	// Define test data
	userRequest := web.RegisterUserRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	// Mock expectations for successful user availability check
	mockUserRepo.On("UserAvailable", userRequest.Username, userRequest.Email).Return(nil, nil)

	// Mock expectations for successful password hashing
	mockPasswordComparer.On("HashPassword", userRequest.Password).Return("hashedPassword", nil)

	// Mock expectations for successful registration
	mockAuthRepo.On("RegisterWithFreeSubscibe", mock.AnythingOfType("*domain.User")).Return(&domain.Auth{
		ID:                1,
		RegistrationToken: "registrationToken123",
	}, nil)

	// Mock expectations for notification creation error
	mockNotification.On("Create", mock.AnythingOfType("*domain.Notification")).Return(fmt.Errorf("error creating notification"))

	// Call the function you want to test
	result, err := authService.RegisterUser(userRequest)

	// Assert that the result is nil (as there's an error in creating a notification)
	require.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockUserRepo.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
	mockAuthRepo.AssertExpectations(t)
	mockNotification.AssertExpectations(t)
	mockFirebase.AssertExpectations(t)
}
