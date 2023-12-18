package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestUserCreate_Success tests the case when creating a user is successful.
func TestUserCreate_Success(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailable method with no existing user
	mockRepository.On("UserAvailable", mock.Anything, mock.Anything).Return(nil, nil)

	// Mock the HashPassword method with success
	mockPasswordComparer.On("HashPassword", mock.Anything).Return("hashedPassword", nil)

	// Mock the HandleTrx method with success
	mockRepository.On("HandleTrx", mock.Anything, mock.Anything).Return(nil)

	// Mock the time.Parse method with success
	mockTime, _ := time.Parse("2006-01-02", "1990-01-01")
	timeParseMock := func(layout, value string) (time.Time, error) {
		return mockTime, nil
	}

	timeParse := timeParseMock
	timeParseOriginal := timeParse
	defer func() { timeParse = timeParseOriginal }()

	// Call the function you want to test
	err := userService.UserCreate(nil, web.UserCreateRequest{
		Username:    "john_doe",
		Email:       "john@example.com",
		Password:    "password123",
		FirstName:   "John",
		LastName:    "Doe",
		DateBirth:   "2006-01-02",
		Bio:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		Gender:      "m",
		PhoneNumber: "123456789",
		City:        "Example City",
		Address:     "123 Example Street",
		Job:         "Software Developer",
		RoleID: 1,
		// Add other necessary fields
	})

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

// TestUserCreate_ExistingUser tests the case when the user already exists.
func TestUserCreate_ExistingUser(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailable method with an existing user
	mockRepository.On("UserAvailable", mock.Anything, mock.Anything).Return(&domain.User{ID :uint(1)}, nil)

	// Call the function you want to test
	err := userService.UserCreate(nil, web.UserCreateRequest{
		Username:    "john_doe",
		Email:       "john@example.com",
		Password:    "password123",
		FirstName:   "John",
		LastName:    "Doe",
		DateBirth:   "2006-01-02",
		Bio:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		Gender:      "m",
		PhoneNumber: "123456789",
		City:        "Example City",
		Address:     "123 Example Street",
		Job:         "Software Developer",
		RoleID: 1,
		// Add other necessary fields
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already exists")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

// TestUserCreate_HashPasswordError tests the case when hashing the password fails.
func TestUserCreate_HashPasswordError(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailable method with no existing user
	mockRepository.On("UserAvailable", mock.Anything, mock.Anything).Return(nil, nil)

	// Mock the HashPassword method with an error
	mockPasswordComparer.On("HashPassword", mock.Anything).Return("", fmt.Errorf("hashing error"))

	// Call the function you want to test
	err := userService.UserCreate(nil, web.UserCreateRequest{
		Username:    "john_doe",
		Email:       "john@example.com",
		Password:    "password123",
		FirstName:   "John",
		LastName:    "Doe",
		DateBirth:   "2006-01-02",
		Bio:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		Gender:      "m",
		PhoneNumber: "123456789",
		City:        "Example City",
		Address:     "123 Example Street",
		Job:         "Software Developer",
		RoleID: 1,
		// Add other necessary fields
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "hashing error")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

