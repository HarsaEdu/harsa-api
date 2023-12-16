package service

import (
	"errors"
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestGetById_Success(t *testing.T) {
	mockRepo := new(mocks.HistoryQuizRepository)
	mockValidator := validator.New()

	// Create a HistoryQuizServiceImpl with the mock repository
	historyQuizService := NewHistoryQuizService(mockRepo, mockValidator)

	// Define test data
	quizID := uint(1)

	// Set up mock expectations for GetById with expected results
	expectedHistoryQuiz := &domain.HistoryQuiz{ID: quizID, /* Add other necessary fields */}
	mockRepo.On("GetById", quizID).Return(expectedHistoryQuiz, nil)

	// Call the function you want to test
	result, err := historyQuizService.GetById(quizID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetById_HistoryQuizNotFound(t *testing.T) {
	mockRepo := new(mocks.HistoryQuizRepository)
	mockValidator := validator.New()

	// Create a HistoryQuizServiceImpl with the mock repository
	historyQuizService := NewHistoryQuizService(mockRepo, mockValidator)

	// Define test data
	quizID := uint(1)

	historyQuiz := &domain.HistoryQuiz{
		ID: uint(1),
	}

	// Set up mock expectations for GetById with no result
	mockRepo.On("GetById", quizID).Return(historyQuiz, fmt.Errorf("history quiz not found"))

	// Call the function you want to test
	result, err := historyQuizService.GetById(quizID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "history quiz not found")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetById_InternalServerError(t *testing.T) {
	mockRepo := new(mocks.HistoryQuizRepository)
	mockValidator := validator.New()

	// Create a HistoryQuizServiceImpl with the mock repository
	historyQuizService := NewHistoryQuizService(mockRepo, mockValidator)

	// Define test data
	quizID := uint(1)

	// Set up mock expectations for GetById with an error
	mockRepo.On("GetById", quizID).Return(nil, errors.New("internal Server Error"))

	// Call the function you want to test
	result, err := historyQuizService.GetById(quizID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "internal Server Error")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}
