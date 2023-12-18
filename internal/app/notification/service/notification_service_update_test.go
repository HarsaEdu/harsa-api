package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestReadNotification_Success tests the case when updating the notification as read is successful.
func TestReadNotification_Success(t *testing.T) {
	mockRepository := new(mocks.NotificationRepository)

	// Create a NotificationServiceImpl with the mock repositories
	notificationService := NewNotificationService(mockRepository, nil)

	// Mock the NotificationRepository FindById method with valid results
	expectedResult := &domain.Notification{
		ID: uint(1),
	} // Replace with your actual struct
	mockRepository.On("FindById", mock.Anything).Return(expectedResult, nil)

	// Mock the NotificationRepository ReadNotification method with success
	mockRepository.On("ReadNotification", mock.Anything).Return(nil)

	// Call the function you want to test
	err := notificationService.ReadNotification(1) // Replace 1 with an actual ID

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestReadNotification_ErrorNotificationNotFound tests the case when the notification is not found.
func TestReadNotification_ErrorNotificationNotFound(t *testing.T) {
	mockRepository := new(mocks.NotificationRepository)

	// Create a NotificationServiceImpl with the mock repositories
	notificationService := NewNotificationService(mockRepository, nil)

	// Mock the NotificationRepository FindById method with no notification found
	mockRepository.On("FindById", mock.Anything).Return(nil, fmt.Errorf("error"))

	// Call the function you want to test
	err := notificationService.ReadNotification(1) // Replace 1 with an actual ID

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "notification not found")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestReadNotification_ErrorUpdateRead tests the case when there's an error updating the read status.
func TestReadNotification_ErrorUpdateRead(t *testing.T) {
	mockRepository := new(mocks.NotificationRepository)

	// Create a NotificationServiceImpl with the mock repositories
	notificationService := NewNotificationService(mockRepository, nil)

	// Mock the NotificationRepository FindById method with valid results
	expectedResult := &domain.Notification{
		ID: uint(1),
	} // Replace with your actual struct
	mockRepository.On("FindById", mock.Anything).Return(expectedResult, nil)

	// Mock the NotificationRepository ReadNotification method with an error
	mockRepository.On("ReadNotification", mock.Anything).Return(fmt.Errorf("error"))

	// Call the function you want to test
	err := notificationService.ReadNotification(1) // Replace 1 with an actual ID

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when updating read")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestArsipNotification_Success tests the case when archiving the notification is successful.
func TestArsipNotification_Success(t *testing.T) {
	mockRepository := new(mocks.NotificationRepository)

	// Create a NotificationServiceImpl with the mock repositories
	notificationService := NewNotificationService(mockRepository, nil)

	// Mock the NotificationRepository FindById method with valid results
	expectedResult := &domain.Notification{
		ID: uint(1),
	} // Replace with your actual struct
	mockRepository.On("FindById", mock.Anything).Return(expectedResult, nil)

	// Mock the NotificationRepository ArsipNotification method with success
	mockRepository.On("ArsipNotification", mock.Anything).Return(nil)

	// Call the function you want to test
	err := notificationService.ArsipNotification(web.ArsipNotification{ID: 1}) // Replace 1 with an actual ID

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestArsipNotification_ErrorNotificationNotFound tests the case when the notification is not found.
func TestArsipNotification_ErrorNotificationNotFound(t *testing.T) {
	mockRepository := new(mocks.NotificationRepository)

	// Create a NotificationServiceImpl with the mock repositories
	notificationService := NewNotificationService(mockRepository, nil)

	// Mock the NotificationRepository FindById method with no notification found
	mockRepository.On("FindById", mock.Anything).Return(nil, fmt.Errorf("error"))

	// Call the function you want to test
	err := notificationService.ArsipNotification(web.ArsipNotification{ID: 1}) // Replace 1 with an actual ID

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "notification not found")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestArsipNotification_ErrorArchiveNotification tests the case when there's an error archiving the notification.
func TestArsipNotification_ErrorArchiveNotification(t *testing.T) {
	mockRepository := new(mocks.NotificationRepository)

	// Create a NotificationServiceImpl with the mock repositories
	notificationService := NewNotificationService(mockRepository, nil)

	// Mock the NotificationRepository FindById method with valid results
	expectedResult := &domain.Notification{
		ID: uint(1),
	} // Replace with your actual struct
	mockRepository.On("FindById", mock.Anything).Return(expectedResult, nil)

	// Mock the NotificationRepository ArsipNotification method with an error
	mockRepository.On("ArsipNotification", mock.Anything).Return(fmt.Errorf("error"))

	// Call the function you want to test
	err := notificationService.ArsipNotification(web.ArsipNotification{ID: 1}) // Replace 1 with an actual ID

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when updating read")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// ... (You can continue adding more test cases as needed)

