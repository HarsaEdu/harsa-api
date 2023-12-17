package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAllFaqs_Success(t *testing.T) {
	mockRepo := new(mocks.FaqRepository)

	// Create a FaqsServiceImpl with the mock repository
	faqsService := NewFaqsService(mockRepo, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"

	// Mock the repository GetAll method with no error
	mockRepo.On("GetAll", offset, limit, search).Return([]domain.Faqs{}, int64(1), nil)

	// Call the function you want to test
	result, pagination, err := faqsService.GetAll(offset, limit, search)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, pagination)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetAllFaqs_FaqsNotFound(t *testing.T) {
	mockRepo := new(mocks.FaqRepository)

	// Create a FaqsServiceImpl with the mock repository
	faqsService := NewFaqsService(mockRepo, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"

	// Mock the repository GetAll method with faqs not found
	mockRepo.On("GetAll", offset, limit, search).Return(nil, int64(0), nil)

	// Call the function you want to test
	result, pagination, err := faqsService.GetAll(offset, limit, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Contains(t, err.Error(), "faqs not found")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetAllFaqs_NotFoundError(t *testing.T) {
	mockRepo := new(mocks.FaqRepository)

	// Create a FaqsServiceImpl with the mock repository
	faqsService := NewFaqsService(mockRepo, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"

	// Mock the repository GetAll method with an internal server error
	mockRepo.On("GetAll", offset, limit, search).Return(nil, int64(0), fmt.Errorf("internal server error"))

	// Call the function you want to test
	result, pagination, err := faqsService.GetAll(offset, limit, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Contains(t, err.Error(), "faqs not found")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetAllFaqs_InternalServerError(t *testing.T) {
	mockRepo := new(mocks.FaqRepository)

	// Create a FaqsServiceImpl with the mock repository
	faqsService := NewFaqsService(mockRepo, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"

	// Mock the repository GetAll method with an internal server error
	mockRepo.On("GetAll", offset, limit, search).Return(nil, int64(1), fmt.Errorf("internal server error"))

	// Call the function you want to test
	result, pagination, err := faqsService.GetAll(offset, limit, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Contains(t, err.Error(), "internal Server Error")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}


func TestGetFaqsById_Success(t *testing.T) {
	mockRepo := new(mocks.FaqRepository)

	// Create a FaqsServiceImpl with the mock repository
	faqsService := NewFaqsService(mockRepo, nil)

	// Define test data
	faqsID := 1

	// Mock the repository FindById method with no error
	mockRepo.On("FindById", faqsID).Return(&domain.Faqs{}, nil)

	// Call the function you want to test
	result, err := faqsService.GetById(faqsID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetFaqsById_FaqsNotFound(t *testing.T) {
	mockRepo := new(mocks.FaqRepository)

	// Create a FaqsServiceImpl with the mock repository
	faqsService := NewFaqsService(mockRepo, nil)

	// Define test data
	faqsID := 1

	// Mock the repository FindById method with faqs not found
	mockRepo.On("FindById", faqsID).Return(nil, nil)

	// Call the function you want to test
	result, err := faqsService.GetById(faqsID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "faqs not found")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetFaqsById_NotFoundError(t *testing.T) {
	mockRepo := new(mocks.FaqRepository)

	// Create a FaqsServiceImpl with the mock repository
	faqsService := NewFaqsService(mockRepo, nil)

	// Define test data
	faqsID := 1

	// Mock the repository FindById method with an internal server error
	mockRepo.On("FindById", faqsID).Return(nil, fmt.Errorf("internal server error"))

	// Call the function you want to test
	result, err := faqsService.GetById(faqsID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "faqs not found")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetFaqsById_InternalServerError(t *testing.T) {
	mockRepo := new(mocks.FaqRepository)

	// Create a FaqsServiceImpl with the mock repository
	faqsService := NewFaqsService(mockRepo, nil)

	// Define test data
	faqsID := 1

	// Mock the repository FindById method with an internal server error
	mockRepo.On("FindById", faqsID).Return(&domain.Faqs{}, fmt.Errorf("internal server error"))

	// Call the function you want to test
	result, err := faqsService.GetById(faqsID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "internal Server Error")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}