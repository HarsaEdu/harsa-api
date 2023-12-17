package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateWeb_Success(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)
	mockValidator := validator.New()

	// Create a SubmissionAnswerServiceImpl with the mock repositories
	submissionAnswerService := NewSubmissionAnswer(mockRepository, mockValidator, nil, nil)

	id := 1
	request := &web.SubmissionAnswerUpdateWeb{}

	// Mock the FindById method with no error
	mockRepository.On("FindById", id).Return(&domain.SubmissionAnswer{ID: uint(id)}, nil)

	// Mock the UpdateWeb method with no error
	mockRepository.On("UpdateWeb", mock.Anything, id).Return(nil)

	// Call the function you want to test
	err := submissionAnswerService.UpdateWeb(nil, request, id)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

func TestUpdateWeb_SubmissionAnswerNotFound(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)
	mockValidator := validator.New()

	// Create a SubmissionAnswerServiceImpl with the mock repositories
	submissionAnswerService := NewSubmissionAnswer(mockRepository, mockValidator, nil, nil)

	id := 1
	request := &web.SubmissionAnswerUpdateWeb{}

	// Mock the FindById method with no error, but submission answer not found
	mockRepository.On("FindById", id).Return(nil, fmt.Errorf("find by id error"))

	// Call the function you want to test
	err := submissionAnswerService.UpdateWeb(nil, request, id)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "submission answer not found")

	// Assert that no further calls were made
	mockRepository.AssertExpectations(t)
}

func TestUpdateWeb_UpdateWebError(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)
	mockValidator := validator.New()

	// Create a SubmissionAnswerServiceImpl with the mock repositories
	submissionAnswerService := NewSubmissionAnswer(mockRepository, mockValidator, nil, nil)

	id := 1
	request := &web.SubmissionAnswerUpdateWeb{}

	// Mock the FindById method with no error
	mockRepository.On("FindById", id).Return(&domain.SubmissionAnswer{ID: uint(id)}, nil)

	// Mock the UpdateWeb method with an error
	mockRepository.On("UpdateWeb", mock.Anything, id).Return(fmt.Errorf("update web error"))

	// Call the function you want to test
	err := submissionAnswerService.UpdateWeb(nil, request, id)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when updating")

	// Assert that no further calls were made
	mockRepository.AssertExpectations(t)
}
