package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAllFeedbacks_Success(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)

	// Create a FeedbackServiceImpl with the mock repository
	feedbackService := NewFeedbackService(mockFeedbackRepo, nil)

	// Define test data
	courseID := uint(1) // Use uint type for courseId
	offset := 0
	limit := 10
	search := "test"

	// Set up mock expectations
	mockFeedbackRepo.On("GetAllByCourseId", courseID, offset, limit, search).Return([]domain.Feedback{{ID: 1, Rating: 1}}, int64(1), nil)

	// Call the function you want to test
	result, pagination, err := feedbackService.GetAllByCourseId(courseID, offset, limit, search)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, pagination)
	assert.Equal(t, 1, len(result))

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

func TestGetAllFeedbacks_NoFeedbacksFound(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)

	// Create a FeedbackServiceImpl with the mock repository
	feedbackService := NewFeedbackService(mockFeedbackRepo, nil)

	// Define test data
	courseID := uint(1)
	offset := 0
	limit := 10
	search := "test"

	// Set up mock expectations for no feedbacks found
	mockFeedbackRepo.On("GetAllByCourseId", courseID, offset, limit, search).
		Return(nil, int64(0), nil)

	// Call the function you want to test
	result, pagination, err := feedbackService.GetAllByCourseId(courseID, offset, limit, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.EqualError(t, err, "feedback not found")

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

func TestGetAllFeedbacks_ErrorGettingFeedbacks(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)

	// Create a FeedbackServiceImpl with the mock repository
	feedbackService := NewFeedbackService(mockFeedbackRepo, nil)

	// Define test data
	courseId := uint(1)
	offset := 0
	limit := 10
	search := "test"

	// Set up mock expectations for an error during Feedback retrieval
	mockFeedbackRepo.On("GetAllByCourseId", courseId, offset, limit, search).
		Return(nil, int64(1), fmt.Errorf("simulated retrieval error"))

	// Call the function you want to test
	result, pagination, err := feedbackService.GetAllByCourseId(courseId, offset, limit, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.EqualError(t, err, "internal Server Error")

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

// Add more test cases as needed for different scenarios
