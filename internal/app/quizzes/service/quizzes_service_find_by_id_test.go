package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestFindByIdQuiz_Success(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	quizId := uint(1)

	quiz := &domain.Quizzes{
		// Populate with necessary fields
	}

	// Set up mock expectations for FindById with no error
	mockRepo.On("FindById", quizId).Return(quiz, nil)

	// Call the function you want to test
	result, err := quizzesService.FindById(quizId)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
	
}

func TestFindByIdQuiz_Error(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	quizId := uint(1)

	// Set up mock expectations for FindById with an error
	mockRepo.On("FindById", quizId).Return(nil, fmt.Errorf("error finding quiz"))

	// Call the function you want to test
	result, err := quizzesService.FindById(quizId)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "error when find quiz by id :error finding quiz")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
	
}

func TestFindByIdMobile_Success(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	quizId := uint(1)

	quiz := &domain.Quizzes{
		// Populate with necessary fields
	}

	// Set up mock expectations for FindById with no error
	mockRepo.On("FindById", quizId).Return(quiz, nil)

	// Call the function you want to test
	result, err := quizzesService.FindByIdMobile(quizId)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)

	// Add assertions specific to the FindByIdMobile conversion if needed
}

func TestFindByIdMobile_Error(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	quizId := uint(1)

	// Set up mock expectations for FindById with an error
	mockRepo.On("FindById", quizId).Return(nil, fmt.Errorf("error finding quiz"))

	// Call the function you want to test
	result, err := quizzesService.FindByIdMobile(quizId)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "error when find quiz by id :error finding quiz")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}