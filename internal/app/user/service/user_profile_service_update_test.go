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

// TestUserProfileUpdate_Success tests the case when updating user profile is successful.
func TestUserProfileUpdate_Success(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)
	user:=web.UserProfileUpdateRequest{
		ID:          1,
		FirstName:   "John",
		LastName:    "Doe",
		DateBirth:   "2006-01-02",
		Bio:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		Gender:      "m",
		PhoneNumber: "123456789",
		City:        "Example City",
		Address:     "123 Example Street",
		Job:         "Software Developer",
	}
	// Mock the UserRepository UserProfileAvailableByID method with valid results
	expectedResult := &domain.UserProfile{
		ID: uint(1),
	} // Replace with your actual struct
	mockRepository.On("UserProfileAvailableByID", mock.Anything).Return(expectedResult, nil)

	// Mock the UserRepository UserProfileUpdate method with success
	mockRepository.On("UserProfileUpdate", mock.Anything).Return(nil)

	// Call the function you want to test
	err := userService.UserProfileUpdate(user) // Replace 1 with an actual ID

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestUserProfileUpdate_ValidationError tests the case when the request validation fails.
func TestUserProfileUpdate_ValidationError(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)
	user := web.UserProfileUpdateRequest{
		// ... (provide invalid data here to trigger validation error)
	}

	// Call the function you want to test
	err := userService.UserProfileUpdate(user) // Replace 1 with an actual ID

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'UserProfileUpdateRequest.ID")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestUserProfileUpdate_UserProfileNotFound tests the case when the user profile is not found.
func TestUserProfileUpdate_UserProfileNotFound(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)
	user := web.UserProfileUpdateRequest{
		ID:          1,
		FirstName:   "John",
		LastName:    "Doe",
		DateBirth:   "2006-01-02",
		Bio:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		Gender:      "m",
		PhoneNumber: "123456789",
		City:        "Example City",
		Address:     "123 Example Street",
		Job:         "Software Developer",
	}

	// Mock the UserRepository UserProfileAvailableByID method with no profile found
	mockRepository.On("UserProfileAvailableByID", mock.Anything).Return(nil, fmt.Errorf("profile not found"))

	// Call the function you want to test
	err := userService.UserProfileUpdate(user) // Replace 1 with an actual ID

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user profile not found")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestUserProfileUpdate_FailedToUpdateUserProfile tests the case when updating user profile fails.
func TestUserProfileUpdate_FailedToUpdateUserProfile(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)
	user := web.UserProfileUpdateRequest{
		ID:          1,
		FirstName:   "John",
		LastName:    "Doe",
		DateBirth:   "2006-01-02",
		Bio:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		Gender:      "m",
		PhoneNumber: "123456789",
		City:        "Example City",
		Address:     "123 Example Street",
		Job:         "Software Developer",
	}

	// Mock the UserRepository UserProfileAvailableByID method with valid results
	expectedResult := &domain.UserProfile{
		ID: uint(1),
	} // Replace with your actual struct
	mockRepository.On("UserProfileAvailableByID", mock.Anything).Return(expectedResult, nil)

	// Mock the UserRepository UserProfileUpdate method with an update failure
	mockRepository.On("UserProfileUpdate", mock.Anything).Return(fmt.Errorf("update failed"))

	// Call the function you want to test
	err := userService.UserProfileUpdate(user) // Replace 1 with an actual ID

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when update user profile")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}
