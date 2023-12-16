package service

import (
	"errors"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"

)

func TestGetAllByQuizWeb_Success(t *testing.T) {
	mockRepo := new(mocks.HistoryQuizRepository)
	mockValidator := validator.New()

	// Create a HistoryQuizServiceImpl with the mock repository
	historyQuizService := NewHistoryQuizService(mockRepo, mockValidator)

	// Define test data
	quizID := uint(1)
	offset := 0
	limit := 10
	search := "example"

	// Set up mock expectations for GetAllHistoryQuizByQuizId with expected results
	expectedHistoryQuizzes := []domain.HistoryQuiz{
		{ID: 1, QuizID: quizID, /* Add other necessary fields */},
		{ID: 2, QuizID: quizID, /* Add other necessary fields */},
	}
	mockRepo.On("GetAllHistoryQuizByQuizId", offset, limit, quizID, search).Return(expectedHistoryQuizzes, int64(len(expectedHistoryQuizzes)), nil)

	// Call the function you want to test
	result, pagination, err := historyQuizService.GetAllByQuizWeb(quizID, offset, limit, search)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, pagination)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetAllByQuizWeb_NoHistoryQuizzesFound(t *testing.T) {
	mockRepo := new(mocks.HistoryQuizRepository)
	mockValidator := validator.New()

	// Create a HistoryQuizServiceImpl with the mock repository
	historyQuizService := NewHistoryQuizService(mockRepo, mockValidator)

	// Define test data
	quizID := uint(1)
	offset := 0
	limit := 10
	search := "example"

	// Set up mock expectations for GetAllHistoryQuizByQuizId with no results
	mockRepo.On("GetAllHistoryQuizByQuizId", offset, limit, quizID, search).Return(nil, int64(0), nil)

	// Call the function you want to test
	result, pagination, err := historyQuizService.GetAllByQuizWeb(quizID, offset, limit, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.EqualError(t, err, "history quiz not found")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetAllByQuizWeb_InternalServerError(t *testing.T) {
	mockRepo := new(mocks.HistoryQuizRepository)
	mockValidator := validator.New()

	// Create a HistoryQuizServiceImpl with the mock repository
	historyQuizService := NewHistoryQuizService(mockRepo, mockValidator)

	// Define test data
	quizID := uint(1)
	offset := 0
	limit := 10
	search := "example"

	// Set up mock expectations for GetAllHistoryQuizByQuizId with an error
	mockRepo.On("GetAllHistoryQuizByQuizId", offset, limit, quizID, search).Return(nil, int64(0), errors.New("internal Server Error"))

	// Call the function you want to test
	result, pagination, err := historyQuizService.GetAllByQuizWeb(quizID, offset, limit, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.EqualError(t, err, "history quiz not found")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetAllByQuizWeb_InternalServerError2(t *testing.T) {
	mockRepo := new(mocks.HistoryQuizRepository)
	mockValidator := validator.New()

	// Create a HistoryQuizServiceImpl with the mock repository
	historyQuizService := NewHistoryQuizService(mockRepo, mockValidator)

	// Define test data
	quizID := uint(1)
	offset := 0
	limit := 10
	search := "example"

	// Set up mock expectations for GetAllHistoryQuizByQuizId with an error
	mockRepo.On("GetAllHistoryQuizByQuizId", offset, limit, quizID, search).Return(nil, int64(1), errors.New("internal Server Error"))

	// Call the function you want to test
	result, pagination, err := historyQuizService.GetAllByQuizWeb(quizID, offset, limit, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.EqualError(t, err, "internal Server Error")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}
