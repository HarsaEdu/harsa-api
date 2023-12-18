package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestGetAll_Success tests the case when fetching all notifications is successful.
func TestGetAll_Success(t *testing.T) {
	mockRepository := new(mocks.NotificationRepository)

	// Create a NotificationServiceImpl with the mock repositories
	notificationService := NewNotificationService(mockRepository, nil)

	// Mock the NotificationRepository GetAll method with valid results
	expectedResult := domain.Notification{
		ID: uint(1),
	} // Replace with your actual struct

	notifications := []domain.Notification{
		expectedResult,
	}
	mockRepository.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(notifications, int64(10), nil)

	// Call the function you want to test
	results, pagination, err := notificationService.GetAll(0, 10, 123)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, pagination)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestGetAll_ErrorNotificationNotFound tests the case when no notifications are found.
func TestGetAll_ErrorNotificationNotFound(t *testing.T) {
	mockRepository := new(mocks.NotificationRepository)

	// Create a NotificationServiceImpl with the mock repositories
	notificationService := NewNotificationService(mockRepository, nil)

	// Mock the NotificationRepository GetAll method with no notifications found
	mockRepository.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(0), fmt.Errorf("error"))

	// Call the function you want to test
	results, pagination, err := notificationService.GetAll(0, 10, 123)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "notification not found")
	assert.Nil(t, results)
	assert.Nil(t, pagination)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestGetAll_ErrorInternalServerError tests the case when there's an internal server error.
func TestGetAll_ErrorInternalServerError(t *testing.T) {
	mockRepository := new(mocks.NotificationRepository)

	// Create a NotificationServiceImpl with the mock repositories
	notificationService := NewNotificationService(mockRepository, nil)

	// Mock the NotificationRepository GetAll method with an internal server error
	mockRepository.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(1), fmt.Errorf("error"))

	// Call the function you want to test
	results, pagination, err := notificationService.GetAll(0, 10, 123)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "internal Server Error")
	assert.Nil(t, results)
	assert.Nil(t, pagination)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

func TestGetById_Success(t *testing.T) {
	mockRepository := new(mocks.NotificationRepository)

	// Create a NotificationServiceImpl with the mock repositories
	notificationService := NewNotificationService(mockRepository, nil)

	// Mock the NotificationRepository FindById method with valid results
	expectedResult := &domain.Notification{
		ID: uint(1),
	} // Replace with your actual struct
	mockRepository.On("FindById", mock.Anything).Return(expectedResult, nil)

	// Call the function you want to test
	result, err := notificationService.GetById(1) // Replace 1 with an actual ID

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestGetById_ErrorNotificationNotFound tests the case when the notification is not found.
func TestGetById_ErrorNotificationNotFound(t *testing.T) {
	mockRepository := new(mocks.NotificationRepository)

	// Create a NotificationServiceImpl with the mock repositories
	notificationService := NewNotificationService(mockRepository, nil)

	// Mock the NotificationRepository FindById method with no notification found
	mockRepository.On("FindById", mock.Anything).Return(nil, fmt.Errorf("error"))

	// Call the function you want to test
	result, err := notificationService.GetById(1) // Replace 1 with an actual ID

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "notification not found")
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestGetById_ErrorInternalServerError tests the case when there's an internal server error.
func TestGetById_ErrorInternalServerError(t *testing.T) {
	mockRepository := new(mocks.NotificationRepository)

	// Create a NotificationServiceImpl with the mock repositories
	notificationService := NewNotificationService(mockRepository, nil)

	// Mock the NotificationRepository FindById method with an internal server error
	mockRepository.On("FindById", mock.Anything).Return(&domain.Notification{ID: uint(1)}, fmt.Errorf("error"))

	// Call the function you want to test
	result, err := notificationService.GetById(1) // Replace 1 with an actual ID

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "internal Server Error")
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}
