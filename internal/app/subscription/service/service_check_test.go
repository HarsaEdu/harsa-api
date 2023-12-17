package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestIsSubscriptionWeb_Active(t *testing.T) {
	mockRepository := new(mocks.SubscriptionRepository)

	// Create a SubscriptionServiceImpl with the mock repository
	subscriptionService := NewSubscriptionService(mockRepository)

	createdAt := time.Now().Unix() - (7 * 24 * 60 * 60) // Within 7 days
	userID := uint(1)

	// Mock the FindSubscription method with an active subscription
	mockRepository.On("FindSubscription", userID).Return(&domain.Subscription{
		ID:       1,
		UserID:   userID,
		StartDate: time.Now().Add(-48 * time.Hour),
		EndDate:   time.Now().Add(48 * time.Hour),
	}, nil)

	// Call the function you want to test
	result, err := subscriptionService.IsSubscriptionWeb(createdAt, userID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

func TestIsSubscriptionWeb_Inactive(t *testing.T) {
	mockRepository := new(mocks.SubscriptionRepository)

	// Create a SubscriptionServiceImpl with the mock repository
	subscriptionService := NewSubscriptionService(mockRepository)

	createdAt := time.Now().Unix() - (8 * 24 * 60 * 60) // More than 7 days
	userID := uint(1)

	// Mock the FindSubscription method with no subscription
	mockRepository.On("FindSubscription", userID).Return(&domain.Subscription{ID : 0}, nil)

	// Call the function you want to test
	result, err := subscriptionService.IsSubscriptionWeb(createdAt, userID)

	// Assert the result
	assert.NoError(t, err)
	assert.False(t, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

func TestIsSubscriptionWeb_Error(t *testing.T) {
	mockRepository := new(mocks.SubscriptionRepository)

	// Create a SubscriptionServiceImpl with the mock repository
	subscriptionService := NewSubscriptionService(mockRepository)

	createdAt := time.Now().Unix() - (8 * 24 * 60 * 60) // Within 7 days
	userID := uint(1)

	// Mock the FindSubscription method with an error
	mockRepository.On("FindSubscription", userID).Return(nil, fmt.Errorf("find subscription error"))

	// Call the function you want to test
	result, err := subscriptionService.IsSubscriptionWeb(createdAt, userID)

	// Assert the result
	assert.Error(t, err)
	assert.False(t, result)
	assert.Contains(t, err.Error(), "find subscription error")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}
