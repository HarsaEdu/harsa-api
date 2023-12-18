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

// TestUserUpdate_Success tests the case when updating user is successful.
func TestUserUpdate_Success(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with an existing user
	expectedUser := &domain.User{ID: 1, Username: "user1", Email: "user1@example.com"}
	mockRepository.On("UserAvailableByID", mock.Anything).Return(expectedUser, nil)

	// Mock the UserAvailableUsername and UserAvailableEmail methods with no existing user
	mockRepository.On("UserAvailableUsername", mock.Anything).Return(nil, nil)
	mockRepository.On("UserAvailableEmail", mock.Anything).Return(nil, nil)

	// Mock the UserUpdate method with success
	mockRepository.On("UserUpdate", mock.Anything).Return(nil)

	// Mock the HashPassword method with success
	mockPasswordComparer.On("HashPassword", mock.Anything).Return("hashedPassword", nil)

	// Call the function you want to test
	err := userService.UserUpdate(web.UserUpdateRequest{
		ID:       1,
		Username: "updated_user",
		Email:    "updated_user@example.com",
		Password: "new_password",
		RoleID:   1,
		// Add other necessary fields
	})

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

// TestUserUpdate_UserNotFound tests the case when the user is not found.
func TestUserUpdate_UserNotFound(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with no existing user
	mockRepository.On("UserAvailableByID", mock.Anything).Return(nil, nil)

	// Call the function you want to test
	err := userService.UserUpdate(web.UserUpdateRequest{
		ID:       1,
		Username: "updated_user",
		Email:    "updated_user@example.com",
		Password: "new_password",
		RoleID:   1,
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

// TestUserUpdate_UsernameAlreadyExists tests the case when the updated username already exists.
func TestUserUpdate_UsernameAlreadyExists(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with an existing user
	expectedUser := &domain.User{ID: 1, Username: "user1", Email: "user1@example.com"}
	mockRepository.On("UserAvailableByID", mock.Anything).Return(expectedUser, nil)

	// Mock the UserAvailableUsername method with an existing user
	mockRepository.On("UserAvailableUsername", mock.Anything).Return(&domain.User{ID: 2, Username: "updated_user"}, nil)

	// Call the function you want to test
	err := userService.UserUpdate(web.UserUpdateRequest{
		ID:       1,
		Username: "updated_user",
		Email:    "updated_user@example.com",
		Password: "new_password",
		RoleID:   1,
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "username already exists")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

// TestUserUpdate_EmailAlreadyExists tests the case when the updated email already exists.
func TestUserUpdate_EmailAlreadyExists(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with an existing user
	expectedUser := &domain.User{ID: 1, Username: "user1", Email: "user1@example.com"}
	mockRepository.On("UserAvailableByID", mock.Anything).Return(expectedUser, nil)

	// Mock the UserAvailableUsername and UserAvailableEmail methods with no existing user
	mockRepository.On("UserAvailableUsername", mock.Anything).Return(nil, nil)
	mockRepository.On("UserAvailableEmail", mock.Anything).Return(&domain.User{ID: 2, Email: "updated_user@example.com"}, nil)

	// Call the function you want to test
	err := userService.UserUpdate(web.UserUpdateRequest{
		ID:       1,
		Username: "updated_user",
		Email:    "updated_user@example.com",
		Password: "new_password",
		RoleID:   1,
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "email already exists")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

// TestUserUpdate_UpdateError tests the case when updating user encounters an error.
func TestUserUpdate_UpdateError(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with an existing user
	expectedUser := &domain.User{ID: 1, Username: "user1", Email: "user1@example.com"}
	mockRepository.On("UserAvailableByID", mock.Anything).Return(expectedUser, nil)

	// Mock the UserAvailableUsername and UserAvailableEmail methods with no existing user
	mockRepository.On("UserAvailableUsername", mock.Anything).Return(nil, nil)
	mockRepository.On("UserAvailableEmail", mock.Anything).Return(nil, nil)

	// Mock the UserUpdate method with an error
	mockRepository.On("UserUpdate", mock.Anything).Return(fmt.Errorf("update error"))

	// Mock the HashPassword method with success
	mockPasswordComparer.On("HashPassword", mock.Anything).Return("hashedPassword", nil)

	// Call the function you want to test
	err := userService.UserUpdate(web.UserUpdateRequest{
		ID:       1,
		Username: "updated_user",
		Email:    "updated_user@example.com",
		Password: "new_password",
		RoleID:   1,
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "update error")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

// TestUserUpdate_PasswordHashError tests the case when hashing the password encounters an error.
func TestUserUpdate_PasswordHashError(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with an existing user
	expectedUser := &domain.User{ID: 1, Username: "user1", Email: "user1@example.com"}
	mockRepository.On("UserAvailableByID", mock.Anything).Return(expectedUser, nil)

	// Mock the UserAvailableUsername and UserAvailableEmail methods with no existing user
	mockRepository.On("UserAvailableUsername", mock.Anything).Return(nil, nil)
	mockRepository.On("UserAvailableEmail", mock.Anything).Return(nil, nil)

	// Mock the HashPassword method with an error
	mockPasswordComparer.On("HashPassword", mock.Anything).Return("", fmt.Errorf("hash error"))

	// Call the function you want to test
	err := userService.UserUpdate(web.UserUpdateRequest{
		ID:       1,
		Username: "updated_user",
		Email:    "updated_user@example.com",
		Password: "new_password",
		RoleID:   1,
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "hash error")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

func TestUserUpdate_ValidationError(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Call the function you want to test
	err := userService.UserUpdate(web.UserUpdateRequest{
		ID:       1,
		Username: "updated_user",
		Email:    "updated_user@example.com",
		Password: "new_password",
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'UserUpdateRequest.RoleID")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

func TestUserUpdateMobile_Success(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with an existing user
	expectedUser := &domain.User{ID: 1, Username: "user1", Email: "user1@example.com"}
	mockRepository.On("UserAvailableByID", mock.Anything).Return(expectedUser, nil)

	// Mock the UserAvailableUsername and UserAvailableEmail methods with no existing user
	mockRepository.On("UserAvailableUsername", mock.Anything).Return(nil, nil)
	mockRepository.On("UserAvailableEmail", mock.Anything).Return(nil, nil)

	// Mock the UserUpdate method with success
	mockRepository.On("UserUpdate", mock.Anything).Return(nil)

	// Call the function you want to test
	err := userService.UserUpdateMobile(web.UserUpdateRequestMobile{
		ID:       1,
		Username: "updated_user",
		Email:    "updated_user@example.com",

		// Add other necessary fields
	})

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

// TestUserUpdateMobile_UserNotFound tests the case when the user is not found.
func TestUserUpdateMobile_UserNotFound(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with no existing user
	mockRepository.On("UserAvailableByID", mock.Anything).Return(nil, nil)

	// Call the function you want to test
	err := userService.UserUpdateMobile(web.UserUpdateRequestMobile{
		ID:       1,
		Username: "updated_user",
		Email:    "updated_user@example.com",
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

// TestUserUpdateMobile_UsernameAlreadyExists tests the case when the updated username already exists.
func TestUserUpdateMobile_UsernameAlreadyExists(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with an existing user
	expectedUser := &domain.User{ID: 1, Username: "user1", Email: "user1@example.com"}
	mockRepository.On("UserAvailableByID", mock.Anything).Return(expectedUser, nil)

	// Mock the UserAvailableUsername method with an existing user
	mockRepository.On("UserAvailableUsername", mock.Anything).Return(&domain.User{ID: 2, Username: "updated_user"}, nil)

	// Call the function you want to test
	err := userService.UserUpdateMobile(web.UserUpdateRequestMobile{
		ID:       1,
		Username: "updated_user",
		Email:    "updated_user@example.com",
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "username already exists")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

// TestUserUpdateMobile_EmailAlreadyExists tests the case when the updated email already exists.
func TestUserUpdateMobile_EmailAlreadyExists(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with an existing user
	expectedUser := &domain.User{ID: 1, Username: "user1", Email: "user1@example.com"}
	mockRepository.On("UserAvailableByID", mock.Anything).Return(expectedUser, nil)

	// Mock the UserAvailableUsername and UserAvailableEmail methods with no existing user
	mockRepository.On("UserAvailableUsername", mock.Anything).Return(nil, nil)
	mockRepository.On("UserAvailableEmail", mock.Anything).Return(&domain.User{ID: 2, Email: "updated_user@example.com"}, nil)

	// Call the function you want to test
	err := userService.UserUpdateMobile(web.UserUpdateRequestMobile{
		ID:       1,
		Username: "updated_user",
		Email:    "updated_user@example.com",
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "email already exists")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

// TestUserUpdateMobile_UpdateError tests the case when updating user encounters an error.
func TestUserUpdateMobile_UpdateError(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with an existing user
	expectedUser := &domain.User{ID: 1, Username: "user1", Email: "user1@example.com"}
	mockRepository.On("UserAvailableByID", mock.Anything).Return(expectedUser, nil)

	// Mock the UserAvailableUsername and UserAvailableEmail methods with no existing user
	mockRepository.On("UserAvailableUsername", mock.Anything).Return(nil, nil)
	mockRepository.On("UserAvailableEmail", mock.Anything).Return(nil, nil)

	// Mock the UserUpdate method with an error
	mockRepository.On("UserUpdate", mock.Anything).Return(fmt.Errorf("update error"))

	// Call the function you want to test
	err := userService.UserUpdateMobile(web.UserUpdateRequestMobile{
		ID:       1,
		Username: "updated_user",
		Email:    "updated_user@example.com",
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "update error")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}


// TestUserUpdateMobile_ValidationError tests the case when there are validation errors in the request.
func TestUserUpdateMobile_ValidationError(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Call the function you want to test with an invalid request (missing RoleID)
	err := userService.UserUpdateMobile(web.UserUpdateRequestMobile{

		Username: "updated_user",
		Email:    "updated_user@example.com",
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'UserUpdateRequestMobile.ID")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}
func TestUserUpdatePasswordMobile_Success(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with an existing user
	expectedUser := &domain.User{ID: 1, Password: "hashedOldPassword"}
	mockRepository.On("UserAvailableByID", mock.Anything).Return(expectedUser, nil)

	// Mock the Password.ComparePassword method with success
	mockPasswordComparer.On("ComparePassword", expectedUser.Password, mock.Anything).Return(nil)

	// Mock the UserUpdate method with success
	mockRepository.On("UserUpdate", mock.Anything).Return(nil)

	// Mock the HashPassword method with success
	mockPasswordComparer.On("HashPassword", mock.Anything).Return("hashedNewPassword", nil)

	// Call the function you want to test
	err := userService.UserUpdatePasswordMobile(web.UserUpdatePasswordRequestMobile{
		ID:          1,
		OldPassword: "oldPassword",
		NewPassword:    "newPassword",
	})

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

// TestUserUpdatePasswordMobile_UserNotFound tests the case when the user is not found.
func TestUserUpdatePasswordMobile_UserNotFound(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with no existing user
	mockRepository.On("UserAvailableByID", mock.Anything).Return(nil, nil)

	// Call the function you want to test
	err := userService.UserUpdatePasswordMobile(web.UserUpdatePasswordRequestMobile{
		ID:          1,
		OldPassword: "oldPassword",
		NewPassword:    "newPassword",
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

// TestUserUpdatePasswordMobile_InvalidOldPassword tests the case when the old password is invalid.
func TestUserUpdatePasswordMobile_InvalidOldPassword(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with an existing user
	expectedUser := &domain.User{ID: 1, Password: "hashedOldPassword"}
	mockRepository.On("UserAvailableByID", mock.Anything).Return(expectedUser, nil)

	// Mock the Password.ComparePassword method with an error (invalid old password)
	mockPasswordComparer.On("ComparePassword", expectedUser.Password, mock.Anything).Return(fmt.Errorf("invalid password"))

	// Call the function you want to test
	err := userService.UserUpdatePasswordMobile(web.UserUpdatePasswordRequestMobile{
		ID:          1,
		OldPassword: "oldPassword",
		NewPassword:    "newPassword",
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid password")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

// TestUserUpdatePasswordMobile_UpdateError tests the case when updating user encounters an error.
func TestUserUpdatePasswordMobile_UpdateError(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with an existing user
	// Mock the UserAvailableByID method with an existing user
	expectedUser := &domain.User{ID: 1, Password: "hashedOldPassword"}
	mockRepository.On("UserAvailableByID", mock.Anything).Return(expectedUser, nil)

	// Mock the Password.ComparePassword method with success
	mockPasswordComparer.On("ComparePassword", expectedUser.Password, mock.Anything).Return(nil)

	// Mock the HashPassword method with success
	mockPasswordComparer.On("HashPassword", mock.Anything).Return("hashedNewPassword", nil)

	// Mock the UserUpdate method with an error
	mockRepository.On("UserUpdate", mock.Anything).Return(fmt.Errorf("update error"))

	// Call the function you want to test
	err := userService.UserUpdatePasswordMobile(web.UserUpdatePasswordRequestMobile{
		ID:          1,
		OldPassword: "oldPassword",
		NewPassword:    "newPassword",
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "update error")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

// TestUserUpdatePasswordMobile_PasswordHashError tests the case when hashing the password encounters an error.
func TestUserUpdatePasswordMobile_PasswordHashError(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Mock the UserAvailableByID method with an existing user
	expectedUser := &domain.User{ID: 1, Password: "hashedOldPassword"}
	mockRepository.On("UserAvailableByID", mock.Anything).Return(expectedUser, nil)

	// Mock the Password.ComparePassword method with success
	mockPasswordComparer.On("ComparePassword", expectedUser.Password, mock.Anything).Return(nil)


	// Mock the HashPassword method with an error
	mockPasswordComparer.On("HashPassword", mock.Anything).Return("", fmt.Errorf("hash error"))

	// Call the function you want to test
	err := userService.UserUpdatePasswordMobile(web.UserUpdatePasswordRequestMobile{
		ID:          1,
		OldPassword: "oldPassword",
		NewPassword:    "newPassword",
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "hash error")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}

// TestUserUpdatePasswordMobile_ValidationError tests the case when there are validation errors in the request.
func TestUserUpdatePasswordMobile_ValidationError(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()
	mockPasswordComparer := new(mocks.PasswordComparer)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, mockPasswordComparer)

	// Call the function you want to test with an invalid request (missing OldPassword)
	err := userService.UserUpdatePasswordMobile(web.UserUpdatePasswordRequestMobile{
		ID:       1,
		NewPassword: "newPassword",
	})

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'UserUpdatePasswordRequestMobile.OldPassword")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordComparer.AssertExpectations(t)
}