package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPaymentHistoryById_Success(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repository
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, nil, nil, mockValidator)

	// Define test data
	orderId := "paymentId"
	mockPaymentHistory := &domain.PaymentHistory{
		ID: orderId,
		// Populate with other necessary fields
	}

	// Mock the GetPaymentHistoryById method with no error
	mockPaymentRepository.On("GetPaymentHistoryById", orderId).Return(mockPaymentHistory, nil)

	// Call the function you want to test
	result, err := paymentService.GetPaymentHistoryById(orderId)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockPaymentRepository.AssertExpectations(t)
}

func TestGetPaymentHistoryById_PaymentHistoryNotFound(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repository
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, nil, nil, mockValidator)

	// Define test data
	orderId := "nonexistentId"

	// Mock the GetPaymentHistoryById method with no error
	mockPaymentRepository.On("GetPaymentHistoryById", orderId).Return(nil, nil)

	// Call the function you want to test
	result, err := paymentService.GetPaymentHistoryById(orderId)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "payment history not found")

	// Assert that the expected calls were made
	mockPaymentRepository.AssertExpectations(t)
}

func TestGetPaymentHistoryById_Error(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repository
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, nil, nil, mockValidator)

	// Define test data
	orderId := "paymentId"

	// Mock the GetPaymentHistoryById method with an error
	mockPaymentRepository.On("GetPaymentHistoryById", orderId).Return(nil, fmt.Errorf("get payment history error"))

	// Call the function you want to test
	result, err := paymentService.GetPaymentHistoryById(orderId)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "error when get payment history by id")

	// Assert that the expected calls were made
	mockPaymentRepository.AssertExpectations(t)
}

func TestGetAllPaymentHistory_Success(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repository
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, nil, nil, mockValidator)

	// Define test data
	offset := 0
	limit := 10
	search := "some search"
	status := "some status"
	mockPaymentHistory := []domain.PaymentHistory{
		{ID : "SAdda"},
	}

	// Mock the GetAllPaymentHistory method with no error
	mockPaymentRepository.On("GetAllPaymentHistory", offset, limit, search, status).Return(mockPaymentHistory, int64(len(mockPaymentHistory)), nil)

	// Call the function you want to test
	result, pagination, err := paymentService.GetAllPaymentHistory(offset, limit, search, status)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, pagination)

	// Assert that the expected calls were made
	mockPaymentRepository.AssertExpectations(t)
}

func TestGetAllPaymentHistory_PaymentHistoryNotFound(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repository
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, nil, nil, mockValidator)

	// Define test data
	offset := 0
	limit := 10
	search := "some search"
	status := "some status"

	// Mock the GetAllPaymentHistory method with no error
	mockPaymentRepository.On("GetAllPaymentHistory", offset, limit, search, status).Return(nil, int64(0), nil)

	// Call the function you want to test
	result, pagination, err := paymentService.GetAllPaymentHistory(offset, limit, search, status)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Contains(t, err.Error(), "payment history not found")

	// Assert that the expected calls were made
	mockPaymentRepository.AssertExpectations(t)
}

func TestGetAllPaymentHistory_Error(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repository
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, nil, nil, mockValidator)

	// Define test data
	offset := 0
	limit := 10
	search := "some search"
	status := "some status"

	// Mock the GetAllPaymentHistory method with an error
	mockPaymentRepository.On("GetAllPaymentHistory", offset, limit, search, status).Return(nil, int64(1), fmt.Errorf("get payment history error"))

	// Call the function you want to test
	result, pagination, err := paymentService.GetAllPaymentHistory(offset, limit, search, status)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Contains(t, err.Error(), "error when get payment history")

	// Assert that the expected calls were made
	mockPaymentRepository.AssertExpectations(t)
}

func TestGetAllPaymentHistoryByUserId_Success(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repository
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, nil, nil, mockValidator)

	// Define test data
	userId := uint(1)
	offset := 0
	limit := 10
	search := "some search"
	status := "some status"
	mockPaymentHistory := []domain.PaymentHistory{
		{ID : "ASda"},
	}

	// Mock the GetAllPaymentHistoryByUserId method with no error
	mockPaymentRepository.On("GetAllPaymentHistoryByUserId", userId, offset, limit, search, status).Return(mockPaymentHistory, int64(len(mockPaymentHistory)), nil)

	// Call the function you want to test
	result, pagination, err := paymentService.GetAllPaymentHistoryByUserId(userId, offset, limit, search, status)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, pagination)

	// Assert that the expected calls were made
	mockPaymentRepository.AssertExpectations(t)
}

func TestGetAllPaymentHistoryByUserId_PaymentHistoryNotFound(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repository
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, nil, nil, mockValidator)

	// Define test data
	userId := uint(1)
	offset := 0
	limit := 10
	search := "some search"
	status := "some status"

	// Mock the GetAllPaymentHistoryByUserId method with no error
	mockPaymentRepository.On("GetAllPaymentHistoryByUserId", userId, offset, limit, search, status).Return(nil, int64(0), nil)

	// Call the function you want to test
	result, pagination, err := paymentService.GetAllPaymentHistoryByUserId(userId, offset, limit, search, status)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Contains(t, err.Error(), "payment history not found")

	// Assert that the expected calls were made
	mockPaymentRepository.AssertExpectations(t)
}

func TestGetAllPaymentHistoryByUserId_Error(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repository
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, nil, nil, mockValidator)

	// Define test data
	userId := uint(1)
	offset := 0
	limit := 10
	search := "some search"
	status := "some status"

	// Mock the GetAllPaymentHistoryByUserId method with an error
	mockPaymentRepository.On("GetAllPaymentHistoryByUserId", userId, offset, limit, search, status).Return(nil, int64(1), fmt.Errorf("get payment history by user id error"))

	// Call the function you want to test
	result, pagination, err := paymentService.GetAllPaymentHistoryByUserId(userId, offset, limit, search, status)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Contains(t, err.Error(), "error when get payment history by user id")

	// Assert that the expected calls were made
	mockPaymentRepository.AssertExpectations(t)
}

func TestGetPaymentHistoryByUserIdAndPaymentId_Success(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repository
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, nil, nil, mockValidator)

	// Define test data
	userId := uint(1)
	paymentId := "payment123"
	mockPaymentHistory := &domain.PaymentHistory{ID: paymentId}

	// Mock the GetPaymentHistoryByUserIdAndPaymentId method with no error
	mockPaymentRepository.On("GetPaymentHistoryByUserIdAndPaymentId", userId, paymentId).Return(mockPaymentHistory, nil)

	// Call the function you want to test
	result, err := paymentService.GetPaymentHistoryByUserIdAndPaymentId(userId, paymentId)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockPaymentRepository.AssertExpectations(t)
}

func TestGetPaymentHistoryByUserIdAndPaymentId_PaymentHistoryNotFound(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repository
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, nil, nil, mockValidator)

	// Define test data
	userId := uint(1)
	paymentId := "payment123"

	// Mock the GetPaymentHistoryByUserIdAndPaymentId method with no error
	mockPaymentRepository.On("GetPaymentHistoryByUserIdAndPaymentId", userId, paymentId).Return(nil, nil)

	// Call the function you want to test
	result, err := paymentService.GetPaymentHistoryByUserIdAndPaymentId(userId, paymentId)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "payment history not found")

	// Assert that the expected calls were made
	mockPaymentRepository.AssertExpectations(t)
}

func TestGetPaymentHistoryByUserIdAndPaymentId_Error(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repository
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, nil, nil, mockValidator)

	// Define test data
	userId := uint(1)
	paymentId := "payment123"

	// Mock the GetPaymentHistoryByUserIdAndPaymentId method with an error
	mockPaymentRepository.On("GetPaymentHistoryByUserIdAndPaymentId", userId, paymentId).Return(nil, fmt.Errorf("get payment history by user id and payment id error"))

	// Call the function you want to test
	result, err := paymentService.GetPaymentHistoryByUserIdAndPaymentId(userId, paymentId)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "error when get payment history by user id and payment id")

	// Assert that the expected calls were made
	mockPaymentRepository.AssertExpectations(t)
}

func TestGetLastYearPaymentHistory_Success(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repository
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, nil, nil, mockValidator)

	// Define test data
	mockPaymentHistories := []domain.PaymentHistory{
		{ID: "payment1"},
		{ID: "payment2"},
	}
	currentTime := time.Now()

	// Mock the GetLastYearPaymentHistory method with no error
	mockPaymentRepository.On("GetLastYearPaymentHistory", currentTime).Return(mockPaymentHistories, nil)

	// Call the function you want to test
	result, err := paymentService.GetLastYearPaymentHistory()

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockPaymentRepository.AssertExpectations(t)
}

func TestGetLastYearPaymentHistory_PaymentHistoryNotFound(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repository
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, nil, nil, mockValidator)

	// Mock the GetLastYearPaymentHistory method with no error
	mockPaymentRepository.On("GetLastYearPaymentHistory", mock.Anything).Return(nil, nil)

	// Call the function you want to test
	result, err := paymentService.GetLastYearPaymentHistory()

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "payment history not found")

	// Assert that the expected calls were made
	mockPaymentRepository.AssertExpectations(t)
}

func TestGetLastYearPaymentHistory_Error(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repository
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, nil, nil, mockValidator)

	// Mock the GetLastYearPaymentHistory method with an error
	mockPaymentRepository.On("GetLastYearPaymentHistory", mock.Anything).Return(nil, fmt.Errorf("get last year payment history error"))

	// Call the function you want to test
	result, err := paymentService.GetLastYearPaymentHistory()

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "error when get payment history")

	// Assert that the expected calls were made
	mockPaymentRepository.AssertExpectations(t)
}