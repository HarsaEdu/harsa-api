package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindByIdMobileByUserIdAndCourseId_Success(t *testing.T) {
	mockSubscription := new(mocks.SubscriptionService)
	mockCourseTrackingRepo := new(mocks.CourseTrackingRepository)
	mockCourseRepo := new(mocks.CourseRepository)
	mockValidator := validator.New()

	// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
	courseTrackingService := NewCourseTrackingService(mockCourseTrackingRepo, mockValidator, mockCourseRepo, nil, mockSubscription, nil)

	// Define test data
	userID := uint(1)
	courseID := uint(2)

	// Set up mock expectations for IsSubscription
	mockSubscription.On("IsSubscription", mock.Anything, userID).Return(true, nil)

	// Set up mock expectations for FindByUserIdAndCourseID
	mockCourseTrackingRepo.On("FindByUserIdAndCourseID", courseID, userID).Return(&domain.CourseTracking{
		ID: uint(2),
	}, nil)

	// Set up mock expectations for GetByIdMobile (CourseRepository)
	mockCourseRepo.On("GetByIdMobile", courseID).Return(&domain.Course{
		ID: uint(2),
	}, int64(2), int64(3), nil)

	// Set up mock expectations for FindAllModuleTrackingWithProgress
	mockCourseTrackingRepo.On("FindAllModuleTrackingWithProgress", mock.Anything, userID, courseID).Return(nil, float32(0), nil)

	// Call the function you want to test
	result, err := courseTrackingService.FindByIdMobileByUserIdAndCourseId(nil, userID, courseID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockSubscription.AssertExpectations(t)
	mockCourseTrackingRepo.AssertExpectations(t)
	mockCourseRepo.AssertExpectations(t)
}

func TestFindByIdMobileByUserIdAndCourseId_CourseTrackingNotFound(t *testing.T) {
    mockSubscription := new(mocks.SubscriptionService)
    mockCourseTrackingRepo := new(mocks.CourseTrackingRepository)
    mockCourseRepo := new(mocks.CourseRepository)
    mockValidator := validator.New()

    // Create a CourseTrackingServiceImpl with the mock repositories and dependencies
    courseTrackingService := NewCourseTrackingService(mockCourseTrackingRepo, mockValidator, mockCourseRepo, nil, mockSubscription, nil)

    // Define test data
    userID := uint(1)
    courseID := uint(2)

    // Set up mock expectations for IsSubscription
    mockSubscription.On("IsSubscription", mock.Anything, userID).Return(true, nil)

    // Set up mock expectations for FindByUserIdAndCourseID (CourseTracking not found)
    mockCourseTrackingRepo.On("FindByUserIdAndCourseID", courseID, userID).Return(nil, fmt.Errorf("course tracking not found"))

    // Call the function you want to test
    result, err := courseTrackingService.FindByIdMobileByUserIdAndCourseId(nil, userID, courseID)

    // Assert the result
    assert.Error(t, err)
    assert.Nil(t, result)

    // Assert that the expected calls were made
    mockSubscription.AssertExpectations(t)
    mockCourseTrackingRepo.AssertExpectations(t)
    mockCourseRepo.AssertExpectations(t)
}



func TestFindByIdMobileByUserIdAndCourseId_CourseNotFound(t *testing.T) {
    mockSubscription := new(mocks.SubscriptionService)
    mockCourseTrackingRepo := new(mocks.CourseTrackingRepository)
    mockCourseRepo := new(mocks.CourseRepository)
    mockValidator := validator.New()

    // Create a CourseTrackingServiceImpl with the mock repositories and dependencies
    courseTrackingService := NewCourseTrackingService(mockCourseTrackingRepo, mockValidator, mockCourseRepo, nil, mockSubscription, nil)

    // Define test data
    userID := uint(1)
    courseID := uint(2)

    // Set up mock expectations for IsSubscription
    mockSubscription.On("IsSubscription", mock.Anything, userID).Return(true, nil)

    // Set up mock expectations for FindByUserIdAndCourseID (CourseTracking found)
    mockCourseTrackingRepo.On("FindByUserIdAndCourseID", courseID, userID).Return(&domain.CourseTracking{
        ID: uint(2),
    }, nil)

    // Set up mock expectations for GetByIdMobile (Course not found)
    mockCourseRepo.On("GetByIdMobile", courseID).Return(nil, int64(0), int64(0), fmt.Errorf("course not found"))

    // Call the function you want to test
    result, err := courseTrackingService.FindByIdMobileByUserIdAndCourseId(nil, userID, courseID)

    // Assert the result
    assert.Error(t, err)
    assert.Nil(t, result)

    // Assert that the expected calls were made
    mockSubscription.AssertExpectations(t)
    mockCourseTrackingRepo.AssertExpectations(t)
    mockCourseRepo.AssertExpectations(t)
}

func TestFindByIdMobileByUserIdAndCourseId_CourseNotFoundNil(t *testing.T) {
    mockSubscription := new(mocks.SubscriptionService)
    mockCourseTrackingRepo := new(mocks.CourseTrackingRepository)
    mockCourseRepo := new(mocks.CourseRepository)
    mockValidator := validator.New()

    // Create a CourseTrackingServiceImpl with the mock repositories and dependencies
    courseTrackingService := NewCourseTrackingService(mockCourseTrackingRepo, mockValidator, mockCourseRepo, nil, mockSubscription, nil)

    // Define test data
    userID := uint(1)
    courseID := uint(2)

    // Set up mock expectations for IsSubscription
    mockSubscription.On("IsSubscription", mock.Anything, userID).Return(true, nil)

    // Set up mock expectations for FindByUserIdAndCourseID (CourseTracking found)
    mockCourseTrackingRepo.On("FindByUserIdAndCourseID", courseID, userID).Return(&domain.CourseTracking{
        ID: uint(2),
    }, nil)
	

	course := &domain.Course{
		ID: uint(0),
	}
    // Set up mock expectations for GetByIdMobile (Course not found)
    mockCourseRepo.On("GetByIdMobile", courseID).Return(course, int64(0), int64(0), nil)

    // Call the function you want to test
    result, err := courseTrackingService.FindByIdMobileByUserIdAndCourseId(nil, userID, courseID)

    // Assert the result
    assert.Error(t, err)
    assert.Nil(t, result)

    // Assert that the expected calls were made
    mockSubscription.AssertExpectations(t)
    mockCourseTrackingRepo.AssertExpectations(t)
    mockCourseRepo.AssertExpectations(t)
}
