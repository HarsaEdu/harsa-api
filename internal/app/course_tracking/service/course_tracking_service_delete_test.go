package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDelete_Success(t *testing.T) {
	mockTrackingRepo := new(mocks.CourseTrackingRepository)

	// Create a CourseTrackingServiceImpl with the mock repository
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, nil, nil, nil, nil)

	// Define test data
	courseTrackingID := uint(1)
	courseID := uint(2)
	userID := uint(3)
	role := "instructor"

	// Create a mock CourseTracking object for the trackingExist variable
	trackingExist := &domain.CourseTracking{
		// Add necessary fields for the CourseTracking object
	}

	// Set up mock expectations for CekIdFromCourse
	mockTrackingRepo.On("CekIdFromCourse", userID, courseTrackingID, role).Return(trackingExist, nil)

	// Set up mock expectations for Delete (CourseTrackingRepository)
	mockTrackingRepo.On("Delete", trackingExist).Return(nil)

	// Call the function you want to test
	err := courseTrackingService.Delete(courseTrackingID, courseID, userID, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockTrackingRepo.AssertExpectations(t)
}

// Add more test cases
func TestDelete_TrackingNotFound(t *testing.T) {
	mockTrackingRepo := new(mocks.CourseTrackingRepository)

	// Create a CourseTrackingServiceImpl with the mock repository
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, nil, nil, nil, nil)

	// Define test data
	courseTrackingID := uint(1)
	courseID := uint(2)
	userID := uint(3)
	role := "instructor"

	// Set up mock expectations for CekIdFromCourse (Tracking not found)
	mockTrackingRepo.On("CekIdFromCourse", userID, courseTrackingID, role).Return(nil, fmt.Errorf("tracking not found"))

	// Call the function you want to test
	err := courseTrackingService.Delete(courseTrackingID, courseID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "tracking not found")

	// Assert that the expected calls were made
	mockTrackingRepo.AssertExpectations(t)
}

func TestDelete_TrackingDeletionError(t *testing.T) {
	mockTrackingRepo := new(mocks.CourseTrackingRepository)

	// Create a CourseTrackingServiceImpl with the mock repository
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, nil, nil, nil, nil)

	// Define test data
	courseTrackingID := uint(1)
	courseID := uint(2)
	userID := uint(3)
	role := "instructor"

	// Create a mock CourseTracking object for the trackingExist variable
	trackingExist := &domain.CourseTracking{
		// Add necessary fields for the CourseTracking object
	}

	// Set up mock expectations for CekIdFromCourse
	mockTrackingRepo.On("CekIdFromCourse", userID, courseTrackingID, role).Return(trackingExist, nil)

	// Set up mock expectations for Delete (CourseTrackingRepository - Deletion Error)
	mockTrackingRepo.On("Delete", trackingExist).Return(fmt.Errorf("deletion error"))

	// Call the function you want to test
	err := courseTrackingService.Delete(courseTrackingID, courseID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "deletion error")

	// Assert that the expected calls were made
	mockTrackingRepo.AssertExpectations(t)
}
