package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
)

func TestDeleteByUserAndCourseId(t *testing.T) {
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
	mockFeedbackRepo.On("GetByIdUserAndCourseId", userId, courseId).Return(&domain.Feedback{CourseID: courseId}, nil)
	mockFeedbackRepo.On("DeleteByUserAndCourseId", userId, courseId).Return(nil)
	mockFeedbackRepo.On("AutoUpdateRating", courseId).Return(nil)

	// Call the function you want to test
	err := feedbackService.DeleteByUserAndCourseId(userId, courseId)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

func TestDeleteByUserAndCourseId_FeedbackNotFound(t *testing.T) {
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
	err := feedbackService.DeleteByUserAndCourseId(userId, courseId)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "feedback not found")

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

func TestDeleteByUserAndCourseId_DeleteError(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)

	// Create a FeedbackServiceImpl with the mock repositories
	feedbackService := &FeedbackServiceImpl{
		FeedbackRepository: mockFeedbackRepo,
	}

	// Define test data
	userId := uint(1)
	courseId := uint(2)

	// Set up mock expectations for the delete error case
	mockFeedbackRepo.On("GetByIdUserAndCourseId", userId, courseId).Return(&domain.Feedback{CourseID: courseId}, nil)
	mockFeedbackRepo.On("DeleteByUserAndCourseId", userId, courseId).Return(fmt.Errorf("delete error"))

	// Call the function you want to test
	err := feedbackService.DeleteByUserAndCourseId(userId, courseId)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when deleting : delete error")

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}
