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

func TestCreateFaqs_Success(t *testing.T) {
	mockRepo := new(mocks.FaqRepository)
	mockValidator := validator.New()

	// Create a FaqsServiceImpl with the mock repository
	faqsService := NewFaqsService(mockRepo, mockValidator)

	// Define test data
	faqsRequest := web.FaqsRequest{
		Question: "asDSAD",
		Answer:   "sadsad",
	}

	// Mock the repository Create method with no error
	mockRepo.On("Create", mock.Anything).Return(nil)

	// Call the function you want to test
	err := faqsService.Create(faqsRequest)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestCreateFaqs_ValidationFailed(t *testing.T) {
	mockRepo := new(mocks.FaqRepository)
	mockValidator := validator.New()

	// Create a FaqsServiceImpl with the mock repository
	faqsService := NewFaqsService(mockRepo, mockValidator)

	// Define test data
	faqsRequest := web.FaqsRequest{
		// Add necessary fields for FaqsRequest
	}

	// Call the function you want to test
	err := faqsService.Create(faqsRequest)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'FaqsRequest.Question'")

	// Assert that the expected calls were NOT made
	mockRepo.AssertExpectations(t)
}

func TestCreateFaqs_CreateError(t *testing.T) {
	mockRepo := new(mocks.FaqRepository)
	mockValidator := validator.New()

	// Create a FaqsServiceImpl with the mock repository
	faqsService := NewFaqsService(mockRepo, mockValidator)

	// Define test data
	faqsRequest := web.FaqsRequest{
		Question: "asDSAD",
		Answer:   "sadsad",
	}

	// Mock the repository Create method with an error
	mockRepo.On("Create", mock.Anything).Return(fmt.Errorf("error creating faqs"))

	// Call the function you want to test
	err := faqsService.Create(faqsRequest)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when creating faqs error creating faqs")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}
