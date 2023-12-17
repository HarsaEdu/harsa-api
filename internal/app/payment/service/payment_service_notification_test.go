package service

import (
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNotificationPayment_Success(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockMidtransCoreApi := new(mocks.MidtransCoreApi)
	mockSubscriptionService := new(mocks.SubscriptionService)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repositories
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, mockSubscriptionService, mockMidtransCoreApi, mockValidator)

	// Define test data
	mockNotificationPayload := map[string]interface{}{
		"order_id": "order123",
	}

	mockTransactionStatus := "success"
	mockTransactionResult := &domain.PaymentTransactionStatus{
		OrderID: "order123",
	}

	mockTransaction := &domain.PaymentHistory{
		ID:       "payment123",
		UserId:   uint(1),
		Item:     domain.SubsPlan{Duration_days: 30},
		Status:   "success",
	}

	// Mock the CheckTransactionStatus method with no error
	mockMidtransCoreApi.On("CheckTransactionStatus", "order123").Return(mockTransactionStatus, mockTransactionResult, nil)

	// Mock the UpdateStatusPaymentHistory method with no error
	mockPaymentRepository.On("UpdateStatusPaymentHistory", mockTransactionStatus, mockTransactionResult).Return(nil)

	// Mock the GetPaymentHistoryById method with no error
	mockPaymentRepository.On("GetPaymentHistoryById", "order123").Return(mockTransaction, nil)

	// Mock the SubscriptionAdd method with no error
	mockSubscriptionService.On("SubscriptionAdd", mock.Anything, uint(mockTransaction.Item.Duration_days)).Return(nil)

	// Call the function you want to test
	err := paymentService.NotificationPayment(mockNotificationPayload)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockMidtransCoreApi.AssertExpectations(t)
	mockPaymentRepository.AssertExpectations(t)
	mockSubscriptionService.AssertExpectations(t)
}

func TestNotificationPayment_OrderIDNotFound(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockMidtransCoreApi := new(mocks.MidtransCoreApi)
	mockSubscriptionService := new(mocks.SubscriptionService)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repositories
	paymentService := NewPaymentService(mockPaymentRepository, nil, nil, mockSubscriptionService, mockMidtransCoreApi, mockValidator)

	// Define test data with missing "order_id"
	mockNotificationPayload := map[string]interface{}{
		// "order_id": "order123", // Missing "order_id"
	}

	// Call the function you want to test
	err := paymentService.NotificationPayment(mockNotificationPayload)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "order id not found")

	// Assert that no calls were made
	mockMidtransCoreApi.AssertExpectations(t)
	mockPaymentRepository.AssertExpectations(t)
	mockSubscriptionService.AssertExpectations(t)
}

// Add more test cases as needed
