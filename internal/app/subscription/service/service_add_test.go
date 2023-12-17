package service

import (
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubscriptionAdd_AddSubscription(t *testing.T) {
	mockRepository := new(mocks.SubscriptionRepository)

	// Create a SubscriptionServiceImpl with the mock repository
	subscriptionService := NewSubscriptionService(mockRepository)

	userID := uint(1)
	days := uint(30)

	// Mock the FindSubscription method with no existing subscription
	mockRepository.On("FindSubscription", userID).Return(nil, nil)

	// Mock the AddSubscription method with no error
	mockRepository.On("AddSubscription", mock.Anything).Return(nil)

	// Call the function you want to test
	err := subscriptionService.SubscriptionAdd(userID, days)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

func TestSubscriptionAdd_UpdateSubscription(t *testing.T) {
	mockRepository := new(mocks.SubscriptionRepository)

	// Create a SubscriptionServiceImpl with the mock repository
	subscriptionService := NewSubscriptionService(mockRepository)

	userID := uint(1)
	days := uint(30)

	// Mock the FindSubscription method with an existing subscription
	mockRepository.On("FindSubscription", userID).Return(&domain.Subscription{UserID: userID}, nil)

	// Mock the UpdateSubscription method with no error
	mockRepository.On("UpdateSubscription", mock.Anything).Return(nil)

	// Call the function you want to test
	err := subscriptionService.SubscriptionAdd(userID, days)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// Additional test cases for error scenarios can be added similar to previous examples.
