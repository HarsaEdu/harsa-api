package service

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
)

func TestCreateByUserAndCourseId(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)
	validate := validator.New()

	// Create a FeedbackServiceImpl with the mock repositories
	feedbackService := &FeedbackServiceImpl{
		FeedbackRepository: mockFeedbackRepo,
		Validator:          validate,
	}

	// Define test data
	request := web.FeedbackCreateRequest{
		// Set request fields as needed
		Rating: 5, // Example rating value
	}
	userId := uint(1)
	courseId := uint(2)

	// Set up mock expectations
	mockFeedbackRepo.On("Cek", userId, courseId).Return(nil, nil)
	mockFeedbackRepo.On("Create", mock.AnythingOfType("*domain.Feedback")).Return(nil)
	mockFeedbackRepo.On("AutoUpdateRating", courseId).Return(nil)

	// Call the function you want to test
	err := feedbackService.CreateByUserAndCourseId(request, userId, courseId)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

func TestCreateByUserAndCourseId_FeedbackExists(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)
	validate := validator.New()

	// Create a FeedbackServiceImpl with the mock repositories
	feedbackService := &FeedbackServiceImpl{
		FeedbackRepository: mockFeedbackRepo,
		Validator:          validate,
	}

	// Define test data
	request := web.FeedbackCreateRequest{
		// Set request fields as needed
		Rating: 5, // Example rating value
	}
	userId := uint(1)
	courseId := uint(2)

	// Set up mock expectations to simulate existing feedback
	mockFeedbackRepo.On("Cek", userId, courseId).Return(&domain.Feedback{}, nil)

	// Call the function you want to test
	err := feedbackService.CreateByUserAndCourseId(request, userId, courseId)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "feedback already exist")

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

func TestCreateByUserAndCourseId_ValidationFailure(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)
	validate := validator.New()

	// Create a FeedbackServiceImpl with the mock repositories
	feedbackService := &FeedbackServiceImpl{
		FeedbackRepository: mockFeedbackRepo,
		Validator:          validate,
	}

	// Define test data
	request := web.FeedbackCreateRequest{
		// Set request fields as needed
		// Rating is intentionally omitted for validation failure
	}
	userId := uint(1)
	courseId := uint(2)

	// Set up mock expectations
	mockFeedbackRepo.On("Cek", userId, courseId).Return(nil, nil).Once() // Set expectation for Cek method

	// Call the function you want to test
	err := feedbackService.CreateByUserAndCourseId(request, userId, courseId)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Rating' Error:Field validation for 'Rating' failed on the 'required' tag")

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}

func TestCreateByUserAndCourseId_AutoUpdateRatingFailure(t *testing.T) {
	// Create a mock FeedbackRepository
	mockFeedbackRepo := new(mocks.FeedbackRepository)
	validate := validator.New()

	// Create a FeedbackServiceImpl with the mock repositories
	feedbackService := &FeedbackServiceImpl{
		FeedbackRepository: mockFeedbackRepo,
		Validator:          validate,
	}

	// Define test data
	request := web.FeedbackCreateRequest{
		// Set request fields as needed
		Rating: 5, // Example rating value
	}
	userId := uint(1)
	courseId := uint(2)

	// Set up mock expectations for auto-update rating failure
	mockFeedbackRepo.On("Cek", userId, courseId).Return(nil, nil)
	mockFeedbackRepo.On("Create", mock.AnythingOfType("*domain.Feedback")).Return(nil)
	mockFeedbackRepo.On("AutoUpdateRating", courseId).Return(fmt.Errorf("update rating error"))

	// Call the function you want to test
	err := feedbackService.CreateByUserAndCourseId(request, userId, courseId)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when update rating")

	// Assert that the expected calls were made
	mockFeedbackRepo.AssertExpectations(t)
}
