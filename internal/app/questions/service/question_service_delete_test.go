package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDelete_Success(t *testing.T) {
	mockRepo := new(mocks.QuestionsRepository)

	// Create a QuestionsServiceImpl with the mock repository
	questionsService := NewQuestionsService(mockRepo, nil)

	// Define test data
	userID := uint(1)
	questionID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromQuestion with no error
	mockRepo.On("CekIdFromQuestion", userID, questionID, role).Return(&domain.Questions{ID: questionID}, nil)

	// Set up mock expectations for Delete with no error
	mockRepo.On("Delete", mock.Anything).Return(nil)

	// Call the function you want to test
	err := questionsService.Delete(userID, questionID, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestDelete_CekIdFromQuestionError(t *testing.T) {
	mockRepo := new(mocks.QuestionsRepository)

	// Create a QuestionsServiceImpl with the mock repository
	questionsService := NewQuestionsService(mockRepo, nil)

	// Define test data
	userID := uint(1)
	questionID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromQuestion with an error
	mockRepo.On("CekIdFromQuestion", userID, questionID, role).Return(nil, fmt.Errorf("error checking user from question"))

	// Call the function you want to test
	err := questionsService.Delete(userID, questionID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when cek id user in question delete :error checking user from question")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestDelete_DeleteError(t *testing.T) {
	mockRepo := new(mocks.QuestionsRepository)

	// Create a QuestionsServiceImpl with the mock repository
	questionsService := NewQuestionsService(mockRepo, nil)

	// Define test data
	userID := uint(1)
	questionID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromQuestion with no error
	mockRepo.On("CekIdFromQuestion", userID, questionID, role).Return(&domain.Questions{ID: questionID}, nil)

	// Set up mock expectations for Delete with an error
	mockRepo.On("Delete", mock.Anything).Return(fmt.Errorf("error deleting question"))

	// Call the function you want to test
	err := questionsService.Delete(userID, questionID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when delete question : error deleting question")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}
