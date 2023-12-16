package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteQuiz_Success(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	userId := uint(1)
	quizId := uint(2)
	role := "admin"

	quizExist := &domain.Quizzes{
		// Populate with necessary fields
	}

	// Set up mock expectations for CekIdFromQuiz with no error
	mockRepo.On("CekIdFromQuiz", userId, quizId, role).Return(quizExist, nil)

	// Set up mock expectations for Delete with no error
	mockRepo.On("Delete", mock.Anything).Return(nil)

	// Call the function you want to test
	err := quizzesService.Delete(userId, quizId, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestDeleteQuiz_CekIdFromQuizError(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	userId := uint(1)
	quizId := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromQuiz with an error
	mockRepo.On("CekIdFromQuiz", userId, quizId, role).Return(nil, fmt.Errorf("error checking user from quiz"))

	// Call the function you want to test
	err := quizzesService.Delete(userId, quizId, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when cek id user in quiz delete :error checking user from quiz")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
	
}

func TestDeleteQuiz_DeleteError(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	userId := uint(1)
	quizId := uint(2)
	role := "admin"

	quizExist := &domain.Quizzes{
		// Populate with necessary fields
	}

	// Set up mock expectations for CekIdFromQuiz with no error
	mockRepo.On("CekIdFromQuiz", userId, quizId, role).Return(quizExist, nil)

	// Set up mock expectations for Delete with an error
	mockRepo.On("Delete", mock.Anything).Return(fmt.Errorf("error deleting quiz"))

	// Call the function you want to test
	err := quizzesService.Delete(userId, quizId, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when delete quiz : error deleting quiz")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
	
}
