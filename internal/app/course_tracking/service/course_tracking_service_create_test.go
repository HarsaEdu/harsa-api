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

func TestCreate_Success(t *testing.T) {
	mockSubscription := new(mocks.SubscriptionService)
	mockCourseRepo := new(mocks.CourseRepository)
	mockTrackingRepo := new(mocks.CourseTrackingRepository)
	mockQuizzService := new(mocks.QuizzesService)
	mockValidator := validator.New()
	mockFirebase := new(mocks.Firebase)


	// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, mockValidator, mockCourseRepo, mockQuizzService, mockSubscription, mockFirebase)

	// Define test data
	request := web.CourseTrackingRequest{
		UserID:   uint(1),
		CourseID: uint(2),
		// Add other fields as needed
	}

	// Set up mock expectations for IsSubscription
	mockSubscription.On("IsSubscription", mock.Anything, request.UserID).Return(true, nil)

	// Set up mock expectations for Cek
	mockTrackingRepo.On("Cek", request.UserID, request.CourseID).Return(nil, nil)

	// Set up mock expectations for Create (CourseTrackingRepository)
	mockTrackingRepo.On("Create", mock.Anything).Return(nil)

	// Set up mock expectations for NotifEnrolled
	mockTrackingRepo.On("NotifEnrolled", request.UserID, request.CourseID).Return(&web.NotificationPersonal{}, nil)

	// Set up mock expectations for SendNotificationPersonal (Firebase)
	mockFirebase.On("SendNotificationPersonal", mock.Anything).Return(nil)

	// Call the function you want to test
	err := courseTrackingService.Create(nil, request)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockSubscription.AssertExpectations(t)
	mockTrackingRepo.AssertExpectations(t)
	mockFirebase.AssertExpectations(t)
}

func TestCreate_SubscriptionError(t *testing.T) {
	mockSubscription := new(mocks.SubscriptionService)
	mockCourseRepo := new(mocks.CourseRepository)
	mockTrackingRepo := new(mocks.CourseTrackingRepository)
	mockQuizzService := new(mocks.QuizzesService)
	mockValidator := validator.New()
	mockFirebase := new(mocks.Firebase)

	// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, mockValidator, mockCourseRepo, mockQuizzService, mockSubscription, mockFirebase)

	// Define test data
	request := web.CourseTrackingRequest{
		UserID:   uint(1),
		CourseID: uint(2),
		// Add other fields as needed
	}

	// Set up mock expectations for IsSubscription (Error case)
	mockSubscription.On("IsSubscription", mock.Anything, request.UserID).Return(false, fmt.Errorf("subscription error"))

	// Call the function you want to test
	err := courseTrackingService.Create(nil, request)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "subscription error")

	// Assert that the expected calls were made
	mockSubscription.AssertExpectations(t)
}

func TestCreate_TrackingExistsError(t *testing.T) {
	mockSubscription := new(mocks.SubscriptionService)
	mockCourseRepo := new(mocks.CourseRepository)
	mockTrackingRepo := new(mocks.CourseTrackingRepository)
	mockQuizzService := new(mocks.QuizzesService)
	mockValidator := validator.New()
	mockFirebase := new(mocks.Firebase)

	// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, mockValidator, mockCourseRepo, mockQuizzService, mockSubscription, mockFirebase)

	// Define test data
	request := web.CourseTrackingRequest{
		UserID:   uint(1),
		CourseID: uint(2),
		// Add other fields as needed
	}

	// Set up mock expectations for IsSubscription
	mockSubscription.On("IsSubscription", mock.Anything, request.UserID).Return(true, nil)

	// Set up mock expectations for Cek (Tracking exists case)
	mockTrackingRepo.On("Cek", request.UserID, request.CourseID).Return(&domain.CourseTracking{}, nil)

	// Call the function you want to test
	err := courseTrackingService.Create(nil, request)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "course tracking already exist")

	// Assert that the expected calls were made
	mockSubscription.AssertExpectations(t)
	mockTrackingRepo.AssertExpectations(t)
}

func TestCreateWeb_Success(t *testing.T) {
	mockSubscription := new(mocks.SubscriptionService)
	mockCourseRepo := new(mocks.CourseRepository)
	mockTrackingRepo := new(mocks.CourseTrackingRepository)
	mockQuizzService := new(mocks.QuizzesService)
	mockValidator := validator.New()

	// Set the current time to be 24 hours ago to simulate a user who subscribed more than 24 hours ago
	createdAt := time.Now().Add(-24 * time.Hour).Unix()

	// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, mockValidator, mockCourseRepo, mockQuizzService, mockSubscription, nil)

	// Define test data
	request := web.CourseTrackingRequest{
		UserID:   uint(1),
		CourseID: uint(2),
		// Add other fields as needed
	}

	// Set up mock expectations for GetCreatedAt
	mockTrackingRepo.On("GetCreatedAt", request.UserID).Return(createdAt, nil)

	// Set up mock expectations for IsSubscriptionWeb
	mockSubscription.On("IsSubscriptionWeb", createdAt, request.UserID).Return(true, nil)

	// Set up mock expectations for Cek
	mockTrackingRepo.On("Cek", request.UserID, request.CourseID).Return(nil, nil)

	// Set up mock expectations for Create (CourseTrackingRepository)
	mockTrackingRepo.On("Create", mock.Anything).Return(nil)

	// Call the function you want to test
	err := courseTrackingService.CreateWeb(request)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockSubscription.AssertExpectations(t)
	mockTrackingRepo.AssertExpectations(t)
}