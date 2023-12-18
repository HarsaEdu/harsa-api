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

// TestUserGetAll_Success tests the case when getting users is successful.
func TestUserGetAll_Success(t *testing.T) {
	mockRepository := new(mocks.UserRepository)

	// Create a UserServiceImpl with the mock repository
	userService := NewUserService(mockRepository, nil, nil)

	// Mock the UserRepository UserGetAll method with valid results
	expectedUsers := []domain.UserEntity{{ID: 1, Username: "user1"}, {ID: 2, Username: "user2"}}
	mockRepository.On("UserGetAll", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(expectedUsers, int64(2), nil)

	// Mock the RecordToPaginationResponse method with valid results
	expectedPagination := &web.Pagination{Offset: 0, Limit: 10, Total: 2}
	// Call the function you want to test
	users, pagination, err := userService.UserGetAll(0, 10, "", 1)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
	assert.Equal(t, expectedPagination, pagination)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestUserGetAll_NoUsersFound tests the case when no users are found.
func TestUserGetAll_NoUsersFound(t *testing.T) {
	mockRepository := new(mocks.UserRepository)

	// Create a UserServiceImpl with the mock repository
	userService := NewUserService(mockRepository, nil, nil)

	// Mock the UserRepository UserGetAll method with no results
	mockRepository.On("UserGetAll", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(0), nil)

	// Call the function you want to test
	users, pagination, err := userService.UserGetAll(0, 10, "aku", 1)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "users not found")
	assert.Nil(t, users)
	assert.Nil(t, pagination)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestUserGetAll_Error tests the case when an error occurs while getting users.
func TestUserGetAll_Error(t *testing.T) {
	mockRepository := new(mocks.UserRepository)

	// Create a UserServiceImpl with the mock repository
	userService := NewUserService(mockRepository, nil, nil)

	// Mock the UserRepository UserGetAll method with an error
	mockRepository.On("UserGetAll", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(1), fmt.Errorf("database error"))

	// Call the function you want to test
	users, pagination, err := userService.UserGetAll(0, 10, "", 1)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "internal Server Error")
	assert.Nil(t, users)
	assert.Nil(t, pagination)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}
func TestUserGetAllStudentSubscribe_Success(t *testing.T) {
	mockRepository := new(mocks.UserRepository)

	// Create a UserServiceImpl with the mock repository
	userService := NewUserService(mockRepository, nil, nil)

	// Mock the UserRepository UserGetAllStudentSubscribe method with valid results
	expectedUsers := []domain.UserEntity{{ID: 1, Username: "student1"}, {ID: 2, Username: "student2"}}
	mockRepository.On("UserGetAllStudentSubscribe", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(expectedUsers, int64(2), nil)

	// Mock the RecordToPaginationResponse method with valid results
	expectedPagination := &web.Pagination{Offset: 0, Limit: 10, Total: 2}

	// Call the function you want to test
	users, pagination, err := userService.UserGetAllStudentSubscribe(0, 10, "", 1)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
	assert.Equal(t, expectedPagination, pagination)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestUserGetAllStudentSubscribe_NoStudentsFound tests the case when no student subscribers are found.
func TestUserGetAllStudentSubscribe_NoStudentsFound(t *testing.T) {
	mockRepository := new(mocks.UserRepository)

	// Create a UserServiceImpl with the mock repository
	userService := NewUserService(mockRepository, nil, nil)

	// Mock the UserRepository UserGetAllStudentSubscribe method with no results
	mockRepository.On("UserGetAllStudentSubscribe", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(0), nil)

	// Call the function you want to test
	users, pagination, err := userService.UserGetAllStudentSubscribe(0, 10, "aku", 1)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "users not found")
	assert.Nil(t, users)
	assert.Nil(t, pagination)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestUserGetAllStudentSubscribe_Error tests the case when an error occurs while getting student subscribers.
func TestUserGetAllStudentSubscribe_Error(t *testing.T) {
	mockRepository := new(mocks.UserRepository)

	// Create a UserServiceImpl with the mock repository
	userService := NewUserService(mockRepository, nil, nil)

	// Mock the UserRepository UserGetAllStudentSubscribe method with an error
	mockRepository.On("UserGetAllStudentSubscribe", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(1), fmt.Errorf("database error"))

	// Call the function you want to test
	users, pagination, err := userService.UserGetAllStudentSubscribe(0, 10, "", 1)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "internal Server Error")
	assert.Nil(t, users)
	assert.Nil(t, pagination)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestGetUserDetail_Success tests the case when getting user detail is successful.
func TestGetUserDetail_Success(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, nil)

	// Mock the UserRepository GetUserByID method with valid result
	expectedUser := &domain.UserDetail{UserID: 1, Username: "user1"}
	mockRepository.On("GetUserByID", mock.Anything).Return(expectedUser, nil)

	// Call the function you want to test
	userDetail, err := userService.GetUserDetail(web.UserGetByIDRequest{ID: 1})

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, userDetail)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestGetUserDetail_UserNotFound tests the case when the user is not found.
func TestGetUserDetail_UserNotFound(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, nil)

	// Mock the UserRepository GetUserByID method with no result
	mockRepository.On("GetUserByID", mock.Anything).Return(nil, fmt.Errorf("user not found"))

	// Call the function you want to test
	userDetail, err := userService.GetUserDetail(web.UserGetByIDRequest{ID: 1})

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, userDetail)
	assert.Contains(t, err.Error(), "user not found")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestGetUserDetail_Error tests the case when an error occurs while getting user detail.
func TestGetUserDetail_ValidationError(t *testing.T) {
	mockRepository := new(mocks.UserRepository)
	mockValidator := validator.New()

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, mockValidator, nil)

	// Call the function you want to test
	userDetail, err := userService.GetUserDetail(web.UserGetByIDRequest{})

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, userDetail)
	assert.Contains(t, err.Error(), "Key: 'UserGetByIDRequest.ID'")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

func TestGetUserAccount_Success(t *testing.T) {
	mockRepository := new(mocks.UserRepository)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, nil, nil)

	// Mock the UserRepository GetUserAccountByID method with valid result
	expectedUser := &domain.User{ID: 1, Username: "user1"}
	mockRepository.On("GetUserAccountByID", mock.Anything).Return(expectedUser, nil)

	// Call the function you want to test
	userAccount, err := userService.GetUserAccount(1)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, userAccount)
	assert.Equal(t, expectedUser.ID, userAccount.ID)
	assert.Equal(t, expectedUser.Username, userAccount.Username)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestGetUserAccount_UserNotFound tests the case when the user is not found.
func TestGetUserAccount_UserNotFound(t *testing.T) {
	mockRepository := new(mocks.UserRepository)

	// Create a UserServiceImpl with the mock repositories
	userService := NewUserService(mockRepository, nil, nil)

	// Mock the UserRepository GetUserAccountByID method with no result
	mockRepository.On("GetUserAccountByID", mock.Anything).Return(nil, fmt.Errorf("user not found"))

	// Call the function you want to test
	userAccount, err := userService.GetUserAccount(1)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, userAccount)
	assert.Contains(t, err.Error(), "user not found")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}
