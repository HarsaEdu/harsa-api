package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateByUserAndCourseId_Success(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)
	validate := validator.New()

	// Create a FeedbackServiceImpl with the mock repository
	feedbackService := NewFeedbackService(mockFeedbackRepo, validate)

	// Define test data
	request := web.FeedbackUpdateRequest{
		Rating:  5,
		Content: "Great course!",
	}
	userId := uint(1)
	courseId := uint(2)

	// Set up mock expectations to simulate successful results
	mockFeedbackRepo.On("GetByIdUserAndCourseId", userId, courseId).
		Return(&domain.Feedback{}, nil) // Simulate existing feedback

	mockFeedbackRepo.On("UpdateByUserAndCourseId", userId, courseId, mock.AnythingOfType("*domain.Feedback")).
		Return(nil) // Simulate successful update

	mockFeedbackRepo.On("AutoUpdateRating", mock.AnythingOfType("uint")).
		Return(nil) // Simulate successful auto-update rating

	// Call the function you want to test
	err := feedbackService.UpdateByUserAndCourseId(request, userId, courseId)

	// Assert the result
	assert.NoError(t, err) // Assert that no error occurred

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

func TestUpdateByUserAndCourseId_FeedbackNotFound(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)
	validate := validator.New()

	// Create a FeedbackServiceImpl with the mock repository
	feedbackService := NewFeedbackService(mockFeedbackRepo, validate)

	// Define test data
	request := web.FeedbackUpdateRequest{
		Rating:  5,
		Content: "Great course!",
	}
	userId := uint(1)
	courseId := uint(2)

	// Set up mock expectations for feedback not found
	mockFeedbackRepo.On("GetByIdUserAndCourseId", userId, courseId).
		Return(nil, fmt.Errorf("simulated error"))

	// Call the function you want to test
	err := feedbackService.UpdateByUserAndCourseId(request, userId, courseId)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "feedback not found")

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

func TestUpdateByUserAndCourseId_ErrorUpdatingFeedback(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)
	validate := validator.New()

	// Create a FeedbackServiceImpl with the mock repository
	feedbackService := NewFeedbackService(mockFeedbackRepo, validate)

	// Define test data
	request := web.FeedbackUpdateRequest{
		Rating:  5,
		Content: "Great course!",
	}
	userId := uint(1)
	courseId := uint(2)

	// Set up mock expectations for an error during feedback update
	mockFeedbackRepo.On("GetByIdUserAndCourseId", userId, courseId).
		Return(&domain.Feedback{}, nil)
	mockFeedbackRepo.On("UpdateByUserAndCourseId", userId, courseId, mock.AnythingOfType("*domain.Feedback")).
		Return(fmt.Errorf("simulated error"))

	// Call the function you want to test
	err := feedbackService.UpdateByUserAndCourseId(request, userId, courseId)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when updating : simulated error")

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}
