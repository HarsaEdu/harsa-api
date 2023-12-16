package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
)

func TestGetByIdAndCourseId(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)

	// Create a FeedbackServiceImpl with the mock repositories
	feedbackService := &FeedbackServiceImpl{
		FeedbackRepository: mockFeedbackRepo,
	}

	// Define test data
	courseId := uint(1)
	id := uint(2)

	// Set up mock expectations
	mockFeedbackRepo.On("GetByIdAndCourseId", courseId, id).Return(&domain.Feedback{}, nil) // Update return type to *domain.Feedback

	// Call the function you want to test
	result, err := feedbackService.GetByIdAndCourseId(courseId, id)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

func TestGetByIdAndCourseId_FeedbackNotFound(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)

	// Create a FeedbackServiceImpl with the mock repositories
	feedbackService := &FeedbackServiceImpl{
		FeedbackRepository: mockFeedbackRepo,
	}

	// Define test data
	courseId := uint(1)
	id := uint(2)

	// Set up mock expectations for the case where feedback is not found
	mockFeedbackRepo.On("GetByIdAndCourseId", courseId, id).Return(nil, fmt.Errorf("not found"))

	// Call the function you want to test
	result, err := feedbackService.GetByIdAndCourseId(courseId, id)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "feedback not found")
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

func TestGetByIdUserAndCourseId(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)

	// Create a FeedbackServiceImpl with the mock repositories
	feedbackService := &FeedbackServiceImpl{
		FeedbackRepository: mockFeedbackRepo,
	}

	// Define test data
	userId := uint(1)
	courseId := uint(2)

	// Set up mock expectations
	// Update return type to *domain.Feedback
	mockFeedbackRepo.On("GetByIdUserAndCourseId", userId, courseId).Return(&domain.Feedback{}, nil)

	// Call the function you want to test
	result, err := feedbackService.GetByIdUserAndCourseId(userId, courseId)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

func TestGetByIdUserAndCourseId_FeedbackNotFound(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)

	// Create a FeedbackServiceImpl with the mock repositories
	feedbackService := &FeedbackServiceImpl{
		FeedbackRepository: mockFeedbackRepo,
	}

	// Define test data
	userId := uint(1)
	courseId := uint(2)

	// Set up mock expectations for the case where feedback is not found
	mockFeedbackRepo.On("GetByIdUserAndCourseId", userId, courseId).Return(nil, fmt.Errorf("not found"))

	// Call the function you want to test
	result, err := feedbackService.GetByIdUserAndCourseId(userId, courseId)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "feedback not found")
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}
