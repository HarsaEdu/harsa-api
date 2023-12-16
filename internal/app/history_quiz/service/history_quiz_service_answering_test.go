package service

import (
	"errors"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"

)

func TestCreate_Success(t *testing.T) {
	mockRepo := new(mocks.HistoryQuizRepository)
	mockValidator := validator.New()

	// Create a HistoryQuizServiceImpl with the mock repository
	historyQuizService := NewHistoryQuizService(mockRepo, mockValidator)

	// Define test data
	userID := uint(1)
	quizID := uint(2)
	answersRequest := []web.AnswersRequest{
		{Question_id: uint(1), Option_id: uint(1)},
		{Question_id: uint(2), Option_id: uint(5)},
	}

	// Set up mock expectations for Cek
	mockRepo.On("Cek", userID, quizID).Return(nil, nil)

	// Set up mock expectations for Answering
	mockRepo.On("Answering", userID, answersRequest, quizID).Return(nil)

	// Call the function you want to test
	err := historyQuizService.Create(answersRequest, userID, quizID)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestCreate_AlreadyCompletedQuizError(t *testing.T) {
	mockRepo := new(mocks.HistoryQuizRepository)
	mockValidator := validator.New()

	// Create a HistoryQuizServiceImpl with the mock repository
	historyQuizService := NewHistoryQuizService(mockRepo, mockValidator)

	// Define test data
	userID := uint(1)
	quizID := uint(2)
	answersRequest := []web.AnswersRequest{
		{Question_id: uint(1), Option_id: uint(1)},
		{Question_id: uint(2), Option_id: uint(5)},
	}

	historyQuiz := &domain.HistoryQuiz{
		ID: uint(1),
	}

	// Set up mock expectations for Cek
	mockRepo.On("Cek", userID, quizID).Return(historyQuiz, nil)

	// Call the function you want to test
	err := historyQuizService.Create(answersRequest, userID, quizID)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "You have completed the quiz")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestCreate_AnsweringError(t *testing.T) {
	mockRepo := new(mocks.HistoryQuizRepository)
	mockValidator := validator.New()

	// Create a HistoryQuizServiceImpl with the mock repository
	historyQuizService := NewHistoryQuizService(mockRepo, mockValidator)

	// Define test data
	userID := uint(1)
	quizID := uint(2)
	answersRequest := []web.AnswersRequest{
		{Question_id: uint(1), Option_id: uint(1)},
		{Question_id: uint(2), Option_id: uint(5)},
	}

	// Set up mock expectations for Cek
	mockRepo.On("Cek", userID, quizID).Return(nil, nil)

	// Set up mock expectations for Answering with an error
	mockRepo.On("Answering", userID, answersRequest, quizID).Return(errors.New("error answering quiz"))

	// Call the function you want to test
	err := historyQuizService.Create(answersRequest, userID, quizID)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when creating history quiz error answering quiz")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}
