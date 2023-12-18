package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestDelete_Success tests the case when deleting a notification is successful.
func TestDelete_Success(t *testing.T) {
	mockRepository := new(mocks.NotificationRepository)

	// Create a NotificationServiceImpl with the mock repositories
	notificationService := NewNotificationService(mockRepository, nil)

	// Mock the NotificationRepository FindById method with a valid result
	expectedNotification := &domain.Notification{
		ID: uint(1),
	} // Replace with your actual struct
	mockRepository.On("FindById", mock.Anything).Return(expectedNotification, nil)

	// Mock the NotificationRepository Delete method with no error
	mockRepository.On("Delete", 1).Return(nil)

	// Call the function you want to test
	err := notificationService.Delete(1)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestDelete_ErrorNotificationNotFound tests the case when the notification to delete is not found.
func TestDelete_ErrorNotificationNotFound(t *testing.T) {
	mockRepository := new(mocks.NotificationRepository)

	// Create a NotificationServiceImpl with the mock repositories
	notificationService := NewNotificationService(mockRepository, nil)

	// Mock the NotificationRepository FindById method with no notification found
	mockRepository.On("FindById", mock.Anything).Return(nil, fmt.Errorf("error"))

	// Call the function you want to test
	err := notificationService.Delete(1)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "notification not found")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestDelete_ErrorDeleteNotification tests the case when there's an error deleting the notification.
func TestDelete_ErrorDeleteNotification(t *testing.T) {
	mockRepository := new(mocks.NotificationRepository)

	// Create a NotificationServiceImpl with the mock repositories
	notificationService := NewNotificationService(mockRepository, nil)

	// Mock the NotificationRepository FindById method with a valid result
	expectedNotification := &domain.Notification{
		ID: uint(1),
	} // Replace with your actual struct
	mockRepository.On("FindById", mock.Anything).Return(expectedNotification, nil)

	// Mock the NotificationRepository Delete method with an error
	mockRepository.On("Delete", 1).Return(fmt.Errorf("delete error"))

	// Call the function you want to test
	err := notificationService.Delete(1)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when deleting")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// ... (You can continue adding more test cases as needed)
