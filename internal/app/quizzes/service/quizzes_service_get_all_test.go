package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestGetAllQuizzes_Success(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	moduleID := uint(1)
	offset := 0
	limit := 10
	search := ""

	quizzes := domain.Quizzes{
		ID: uint(1),
	}

	// Mock the repository GetAll method with sample data
	mockQuizzes := []domain.Quizzes{
		quizzes,
	}
	mockTotal := int64(len(mockQuizzes))
	mockTitleCourse := "Sample Title Course"

	mockRepo.On("GetAll", moduleID, offset, limit, search).Return(mockQuizzes, mockTotal, mockTitleCourse, nil)

	// Call the function you want to test
	result, pagination, err := quizzesService.GetAll(moduleID, offset, limit, search)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, pagination)
	assert.Len(t, result, len(mockQuizzes))
	assert.Equal(t, pagination.Total, mockTotal)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)

	// Add specific assertions based on your conversion logic if needed
}

func TestGetAllQuizzes_Error(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	moduleID := uint(1)
	offset := 0
	limit := 10
	search := ""

	// Mock the repository GetAll method with an error
	mockRepo.On("GetAll", moduleID, offset, limit, search).Return(nil, int64(0), "", fmt.Errorf("error getting quizzes"))

	// Call the function you want to test
	result, pagination, err := quizzesService.GetAll(moduleID, offset, limit, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Contains(t, err.Error(), "error getting quizzes")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetAllQuizzesNil_Error(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	moduleID := uint(1)
	offset := 0
	limit := 10
	search := ""

	// Mock the repository GetAll method with an error
	mockRepo.On("GetAll", moduleID, offset, limit, search).Return(nil, int64(0), "", nil)

	// Call the function you want to test
	result, pagination, err := quizzesService.GetAll(moduleID, offset, limit, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Contains(t, err.Error(), "quiz not found")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}
