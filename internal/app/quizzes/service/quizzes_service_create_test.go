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

func TestCreateQuiz_Success(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	request := web.QuizRequest{
		UserID:      uint(1),
		Title:       "asdsad",
		Description: "ADSAD",
	}
	role := "admin"

	// Set up mock expectations for CekIdFromModule with no error
	mockRepo.On("CekIdFromModule", request.UserID, request.ModuleID, role).Return(nil)

	// Set up mock expectations for Create with no error
	mockRepo.On("Create", mock.Anything).Return(nil)

	// Call the function you want to test
	err := quizzesService.Create(request, role)

	// Assert the result
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestCreateQuiz_ErValidationor(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	request := web.QuizRequest{}
	role := "admin"

	// Set up mock expectations for Create with no error
	mockRepo.On("Create", mock.Anything).Return(nil)

	// Call the function you want to test
	err := quizzesService.Create(request, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'QuizRequest.UserID'")

}

func TestCreateQuiz_CekIdFromModuleError(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	request := web.QuizRequest{
		UserID:      uint(1),
		Title:       "asdsad",
		Description: "ADSAD",
	}
	role := "admin"

	// Set up mock expectations for CekIdFromModule with an error
	mockRepo.On("CekIdFromModule", request.UserID, request.ModuleID, role).Return(fmt.Errorf("error checking user from module"))

	// Call the function you want to test
	err := quizzesService.Create(request, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when cek id user in quiz update :error checking user from module")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)

}

func TestCreateQuiz_CreateError(t *testing.T) {
	mockRepo := new(mocks.QuizzesRepository)
	mockValidator := validator.New()

	// Create a QuizzesServiceImpl with the mock repository
	quizzesService := NewQuizzesService(mockRepo, mockValidator)

	// Define test data
	request := web.QuizRequest{
		UserID:      uint(1),
		Title:       "asdsad",
		Description: "ADSAD",
	}
	role := "admin"

	// Set up mock expectations for CekIdFromModule with no error
	mockRepo.On("CekIdFromModule", request.UserID, request.ModuleID, role).Return(nil)

	// Set up mock expectations for Create with an error
	mockRepo.On("Create", mock.Anything).Return(fmt.Errorf("error creating quiz"))

	// Call the function you want to test
	err := quizzesService.Create(request, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when creating quiz error creating quiz")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)

}
