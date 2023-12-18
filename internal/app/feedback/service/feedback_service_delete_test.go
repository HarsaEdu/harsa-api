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

func TestDeleteByUserAndCourseId_AutoUpdateRatingError(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)

	// Create a FeedbackServiceImpl with the mock repositories
	feedbackService := &FeedbackServiceImpl{
		FeedbackRepository: mockFeedbackRepo,
	}

	// Define test data
	userId := uint(1)
	courseId := uint(2)

	// Set up mock expectations for the AutoUpdateRating error case
	mockFeedbackRepo.On("GetByIdUserAndCourseId", userId, courseId).Return(&domain.Feedback{CourseID: courseId}, nil)
	mockFeedbackRepo.On("DeleteByUserAndCourseId", userId, courseId).Return(nil)
	mockFeedbackRepo.On("AutoUpdateRating", courseId).Return(fmt.Errorf("auto update rating error"))

	// Call the function you want to test
	err := feedbackService.DeleteByUserAndCourseId(userId, courseId)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when update rating auto update rating error")

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

func TestDeleteById(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)

	// Create a FeedbackServiceImpl with the mock repositories
	feedbackService := &FeedbackServiceImpl{
		FeedbackRepository: mockFeedbackRepo,
	}

	// Define test data
	feedbackId := uint(1)

	// Set up mock expectations
	mockFeedbackRepo.On("GetById", feedbackId).Return(&domain.Feedback{CourseID: 2}, nil)
	mockFeedbackRepo.On("DeleteById", feedbackId).Return(nil)
	mockFeedbackRepo.On("AutoUpdateRating", uint(2)).Return(nil)

	// Call the function you want to test
	err := feedbackService.DeleteById(feedbackId)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

func TestDeleteById_FeedbackNotFound(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)

	// Create a FeedbackServiceImpl with the mock repositories
	feedbackService := &FeedbackServiceImpl{
		FeedbackRepository: mockFeedbackRepo,
	}

	// Define test data
	feedbackId := uint(1)

	// Set up mock expectations for the case where feedback is not found
	mockFeedbackRepo.On("GetById", feedbackId).Return(nil, fmt.Errorf("not found"))

	// Call the function you want to test
	err := feedbackService.DeleteById(feedbackId)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "feedback not found")

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

func TestDeleteById_DeleteError(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)

	// Create a FeedbackServiceImpl with the mock repositories
	feedbackService := &FeedbackServiceImpl{
		FeedbackRepository: mockFeedbackRepo,
	}

	// Define test data
	feedbackId := uint(1)

	// Set up mock expectations for the delete error case
	mockFeedbackRepo.On("GetById", feedbackId).Return(&domain.Feedback{CourseID: 2}, nil)
	mockFeedbackRepo.On("DeleteById", feedbackId).Return(fmt.Errorf("delete error"))

	// Call the function you want to test
	err := feedbackService.DeleteById(feedbackId)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when deleting : delete error")

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

func TestDeleteById_AutoUpdateRatingError(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)

	// Create a FeedbackServiceImpl with the mock repositories
	feedbackService := &FeedbackServiceImpl{
		FeedbackRepository: mockFeedbackRepo,
	}

	// Define test data
	feedbackId := uint(1)

	// Set up mock expectations for the AutoUpdateRating error case
	mockFeedbackRepo.On("GetById", feedbackId).Return(&domain.Feedback{CourseID: 2}, nil)
	mockFeedbackRepo.On("DeleteById", feedbackId).Return(nil)
	mockFeedbackRepo.On("AutoUpdateRating", uint(2)).Return(fmt.Errorf("auto update rating error"))

	// Call the function you want to test
	err := feedbackService.DeleteById(feedbackId)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when update rating auto update rating error")

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}
