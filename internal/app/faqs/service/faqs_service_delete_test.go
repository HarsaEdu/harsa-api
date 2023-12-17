package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDeleteFaqs_Success(t *testing.T) {
	mockRepo := new(mocks.FaqRepository)

	// Create a FaqsServiceImpl with the mock repository
	faqsService := NewFaqsService(mockRepo, nil)

	// Define test data
	faqsID := 1

	// Mock the repository FindById method with no error
	mockRepo.On("FindById", faqsID).Return(&domain.Faqs{}, nil)

	// Mock the repository Delete method with no error
	mockRepo.On("Delete", faqsID).Return(nil)

	// Call the function you want to test
	err := faqsService.Delete(faqsID)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestDeleteFaqs_FaqsNotFound(t *testing.T) {
	mockRepo := new(mocks.FaqRepository)

	// Create a FaqsServiceImpl with the mock repository
	faqsService := NewFaqsService(mockRepo, nil)

	// Define test data
	faqsID := 1

	// Mock the repository FindById method with faqs not found
	mockRepo.On("FindById", faqsID).Return(nil, nil)

	// Call the function you want to test
	err := faqsService.Delete(faqsID)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "faqs not found")

	// Assert that the expected calls were NOT made
	mockRepo.AssertExpectations(t)
}

func TestDeleteFaqs_DeleteError(t *testing.T) {
	mockRepo := new(mocks.FaqRepository)

	// Create a FaqsServiceImpl with the mock repository
	faqsService := NewFaqsService(mockRepo, nil)

	// Define test data
	faqsID := 1

	// Mock the repository FindById method with no error
	mockRepo.On("FindById", faqsID).Return(&domain.Faqs{}, nil)

	// Mock the repository Delete method with an error
	mockRepo.On("Delete", faqsID).Return(fmt.Errorf("error deleting faqs"))

	// Call the function you want to test
	err := faqsService.Delete(faqsID)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when deleting : error deleting faqs")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}
