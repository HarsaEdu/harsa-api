package service

import (
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateFaqs_Success(t *testing.T) {
	mockRepo := new(mocks.FaqRepository)
	mockValidator := validator.New()

	// Create a FaqsServiceImpl with the mock repository
	faqsService := NewFaqsService(mockRepo, mockValidator)

	// Define test data
	faqsUpdateRequest := web.FaqsUpdateRequest{
		// Add necessary fields for FaqsUpdateRequest
	}
	faqsID := 1

	// Mock the repository FindById method with no error
	mockRepo.On("FindById", faqsID).Return(&domain.Faqs{}, nil)

	// Mock the repository Update method with no error
	mockRepo.On("Update", mock.Anything, faqsID).Return(nil)

	// Call the function you want to test
	err := faqsService.Update(faqsUpdateRequest, faqsID)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateFaqs_FaqsNotFound(t *testing.T) {
	mockRepo := new(mocks.FaqRepository)
	mockValidator := validator.New()

	// Create a FaqsServiceImpl with the mock repository
	faqsService := NewFaqsService(mockRepo, mockValidator)

	// Define test data
	faqsUpdateRequest := web.FaqsUpdateRequest{
		// Add necessary fields for FaqsUpdateRequest
	}
	faqsID := 1

	// Mock the repository FindById method with faqs not found
	mockRepo.On("FindById", faqsID).Return(nil, nil)

	// Call the function you want to test
	err := faqsService.Update(faqsUpdateRequest, faqsID)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "faqs not found")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

