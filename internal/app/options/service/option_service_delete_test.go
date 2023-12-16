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
	mockRepo := new(mocks.OptionsRepository)

	// Create an OptionsServiceImpl with the mock repository
	optionsService := NewOptionsService(mockRepo, nil)

	// Define test data
	userID := uint(1)
	optionID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromOption with no error
	mockRepo.On("CekIdFromOption", userID, optionID, role).Return(&domain.Options{ID: optionID}, nil)

	// Set up mock expectations for Delete with no error
	mockRepo.On("Delete", mock.Anything).Return(nil)

	// Call the function you want to test
	err := optionsService.Delete(userID, optionID, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestDelete_CekIdFromOptionError(t *testing.T) {
	mockRepo := new(mocks.OptionsRepository)

	// Create an OptionsServiceImpl with the mock repository
	optionsService := NewOptionsService(mockRepo, nil)

	// Define test data
	userID := uint(1)
	optionID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromOption with an error
	mockRepo.On("CekIdFromOption", userID, optionID, role).Return(nil, fmt.Errorf("error checking user from option"))

	// Call the function you want to test
	err := optionsService.Delete(userID, optionID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when cek id user in optiondelete :error checking user from option")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestDelete_DeleteError(t *testing.T) {
	mockRepo := new(mocks.OptionsRepository)

	// Create an OptionsServiceImpl with the mock repository
	optionsService := NewOptionsService(mockRepo, nil)

	// Define test data
	userID := uint(1)
	optionID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromOption with no error
	mockRepo.On("CekIdFromOption", userID, optionID, role).Return(&domain.Options{ID: optionID}, nil)

	// Set up mock expectations for Delete with an error
	mockRepo.On("Delete", mock.Anything).Return(fmt.Errorf("error deleting option"))

	// Call the function you want to test
	err := optionsService.Delete(userID, optionID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when delete option : error deleting option")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}
