package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateQuiz_Success(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	quizId := uint(1)
	request := web.QuizRequest{
		UserID:      uint(1),
		Title:       "Updated Title",
		Description: "Updated Description",
	}
	role := "admin"

	// Mock the repository CekIdFromQuiz method with no error
	mockRepo.On("CekIdFromQuiz", request.UserID, quizId, role).Return(nil, nil)

	// Mock the repository Update method with no error
	mockRepo.On("Update", mock.Anything, mock.Anything).Return(nil)

	// Call the function you want to test
	err := quizzesService.Update(request, quizId, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateQuiz_CekIdFromQuizError(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	quizId := uint(1)
	request := web.QuizRequest{
		UserID:      uint(1),
		Title:       "Updated Title",
		Description: "Updated Description",
	}
	role := "admin"

	// Mock the repository CekIdFromQuiz method with an error
	mockRepo.On("CekIdFromQuiz", request.UserID, quizId, role).Return(nil, fmt.Errorf("error checking quiz"))

	// Call the function you want to test
	err := quizzesService.Update(request, quizId, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when cek id user in quiz update :error checking quiz")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateQuiz_ValidationFailed(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	quizId := uint(1)
	request := web.QuizRequest{
		// Add necessary fields for QuizRequest
	}
	role := "admin"

	// Call the function you want to test
	err := quizzesService.Update(request, quizId, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'QuizRequest.UserID'")

	// Assert that the expected calls were NOT made
	mockRepo.AssertExpectations(t)
}

func TestUpdateQuiz_UpdateError(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	quizId := uint(1)
	request := web.QuizRequest{
		UserID:      uint(1),
		Title:       "Updated Title",
		Description: "Updated Description",
	}
	role := "admin"

	// Mock the repository CekIdFromQuiz method with no error
	mockRepo.On("CekIdFromQuiz", request.UserID, quizId, role).Return(nil, nil)

	// Mock the repository Update method with an error
	mockRepo.On("Update", mock.Anything, mock.Anything).Return(fmt.Errorf("error updating quiz"))

	// Call the function you want to test
	err := quizzesService.Update(request, quizId, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when update quiz error updating quiz")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}
