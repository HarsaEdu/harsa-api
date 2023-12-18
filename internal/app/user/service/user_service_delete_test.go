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
)

// TestUserDelete_Success tests the case when deleting a user is successful.
func TestUserDelete_Success(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with an existing user
	mockRepository.On("UserAvailableByID", mock.Anything).Return(&domain.User{}, nil)

	// Mock the UserRepository UserDelete method with success
	mockRepository.On("UserDelete", mock.Anything).Return(nil)

	// Call the function you want to test
	err := userService.UserDelete(web.UserDeleteRequest{ID: 1}) // Replace 1 with an actual ID

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	
}

// TestUserDelete_UserNotFound tests the case when the user is not found.
func TestUserDelete_UserNotFound(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with a non-existing user
	mockRepository.On("UserAvailableByID", mock.Anything).Return(nil, nil)

	// Call the function you want to test
	err := userService.UserDelete(web.UserDeleteRequest{ID: 1}) // Replace 1 with an actual ID

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	
}

// TestUserDelete_DeleteError tests the case when deleting a user encounters an error.
func TestUserDelete_DeleteError(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with an existing user
	mockRepository.On("UserAvailableByID", mock.Anything).Return(&domain.User{}, nil)

	// Mock the UserRepository UserDelete method with an error
	mockRepository.On("UserDelete", mock.Anything).Return(fmt.Errorf("delete error"))

	// Call the function you want to test
	err := userService.UserDelete(web.UserDeleteRequest{ID: 1}) // Replace 1 with an actual ID

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "delete error")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	
}
