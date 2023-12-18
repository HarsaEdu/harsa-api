package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestIsSubscription_Active(t *testing.T) {
	mockSubscriptionRepository := new(mocks.SubscriptionRepository)

	// Create a SubscriptionServiceImpl with the mock repository
	subscriptionService := NewSubscriptionService(mockSubscriptionRepository)

	// Mock the FindSubscription method with an active subscription
	mockSubscriptionRepository.On("FindSubscription", mock.Anything).Return(&domain.Subscription{ID: 1, EndDate: time.Now().AddDate(0, 0, 7)}, nil)

	// Call the function you want to test
	result, err := subscriptionService.IsSubscription(createTestEchoContext(), 1)

	// Assert the result
	assert.NoError(t, err)
	assert.True(t, result)

	// Assert that the expected calls were made
	mockSubscriptionRepository.AssertExpectations(t)
}

// TestIsSubscription_NotActive tests the case when the user has no subscription or the subscription is expired.
func TestIsSubscription_NotActive(t *testing.T) {
	mockSubscriptionRepository := new(mocks.SubscriptionRepository)

	// Create a SubscriptionServiceImpl with the mock repository
	subscriptionService := NewSubscriptionService(mockSubscriptionRepository)

	// Mock the FindSubscription method with no subscription
	mockSubscriptionRepository.On("FindSubscription", mock.Anything).Return(&domain.Subscription{ID : 1}, nil)

	// Call the function you want to test
	result, err := subscriptionService.IsSubscription(createTestEchoContext(), 1)

	// Assert the result
	assert.NoError(t, err)
	assert.False(t, result)

	// Assert that the expected calls were made
	mockSubscriptionRepository.AssertExpectations(t)
}

// TestIsSubscription_Error tests the case when an error occurs while checking the subscription.
func TestIsSubscription_Error(t *testing.T) {
	mockSubscriptionRepository := new(mocks.SubscriptionRepository)

	// Create a SubscriptionServiceImpl with the mock repository
	subscriptionService := NewSubscriptionService(mockSubscriptionRepository)

	// Mock the FindSubscription method with an error
	mockSubscriptionRepository.On("FindSubscription", mock.Anything).Return(nil, fmt.Errorf("eror"))

	// Call the function you want to test
	result, err := subscriptionService.IsSubscription(createTestEchoContext(), 1)

	// Assert the result
	assert.Error(t, err)
	assert.False(t, result)

	// Assert that the expected calls were made
	mockSubscriptionRepository.AssertExpectations(t)
}

// Helper function to create a simple echo.Context for testing purposes
func createTestEchoContext() echo.Context {
	// Create a new echo.Context using net/http/httptest
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	// Set a mock value for user create date (replace it with the actual type and value)
	c.Set("create_user", int64(1639200000)) // Mocking user creation timestamp

	return c
}