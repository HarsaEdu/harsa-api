package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePaymentSubscription_Success(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockSubsPlanRepository := new(mocks.SubsPlanRepository)
	mockUserRepository := new(mocks.UserRepository)
	mockSubscriptionService := new(mocks.SubscriptionService)
	mockMidtransCoreApi := new(mocks.MidtransCoreApi)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repositories
	paymentService := NewPaymentService(mockPaymentRepository, mockSubsPlanRepository, mockUserRepository, mockSubscriptionService, mockMidtransCoreApi, mockValidator)

	// Define test data
	request := web.CreatePaymentSubscriptionRequest{
		PlanId:   1,
		BankName: "ASdsad",
	}

	userId := uint(1)

	// Mock the GetUserByID method with no error
	mockUserRepository.On("GetUserByID", userId).Return(&domain.UserDetail{UserID: userId}, nil)

	// Mock the FindById method with no error
	mockSubsPlanRepository.On("FindById", request.PlanId).Return(&domain.SubsPlan{ID: uint(request.PlanId)}, nil)

	// Mock the ChargeTransaction method with no error
	mockMidtransCoreApi.On("ChargeTransaction", mock.Anything).Return(&coreapi.ChargeResponse{
		TransactionTime: "2006-01-02 15:04:05", 
		ExpiryTime: "2006-01-02 15:04:05",
		VaNumbers: []coreapi.VANumber{
			{Bank: "adsa", VANumber: "aSDAD"},
		} ,
	}, nil)

	// Mock the CreatePaymentHistory method with no error
	mockPaymentRepository.On("CreatePaymentHistory", mock.Anything).Return(nil)

	// Mock the GetPaymentHistoryById method with no error
	mockPaymentRepository.On("GetPaymentHistoryById", mock.Anything).Return(&domain.PaymentHistory{ID: "paymentId"}, nil)

	// Call the function you want to test
	result, err := paymentService.CreatePaymentSubscription(&request, userId)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockUserRepository.AssertExpectations(t)
	mockSubsPlanRepository.AssertExpectations(t)
	mockMidtransCoreApi.AssertExpectations(t)
	mockPaymentRepository.AssertExpectations(t)
}

func TestCreatePaymentSubscription_ValidationEror(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockSubsPlanRepository := new(mocks.SubsPlanRepository)
	mockUserRepository := new(mocks.UserRepository)
	mockSubscriptionService := new(mocks.SubscriptionService)
	mockMidtransCoreApi := new(mocks.MidtransCoreApi)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repositories
	paymentService := NewPaymentService(mockPaymentRepository, mockSubsPlanRepository, mockUserRepository, mockSubscriptionService, mockMidtransCoreApi, mockValidator)

	// Define test data
	request := web.CreatePaymentSubscriptionRequest{
		// Populate with necessary fields
	}

	userId := uint(1)

	// Call the function you want to test
	result, err := paymentService.CreatePaymentSubscription(&request, userId)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Key: 'CreatePaymentSubscriptionRequest.PlanId'")

	mockPaymentRepository.AssertExpectations(t)
}

func TestCreatePaymentSubscription_PlanNotFoundError(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockSubsPlanRepository := new(mocks.SubsPlanRepository)
	mockUserRepository := new(mocks.UserRepository)
	mockSubscriptionService := new(mocks.SubscriptionService)
	mockMidtransCoreApi := new(mocks.MidtransCoreApi)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repositories
	paymentService := NewPaymentService(mockPaymentRepository, mockSubsPlanRepository, mockUserRepository, mockSubscriptionService, mockMidtransCoreApi, mockValidator)

	// Define test data
	request := web.CreatePaymentSubscriptionRequest{
		PlanId:   1,
		BankName: "ASdsad",
	}

	userId := uint(1)

	// Mock the GetUserByID method with no error
	mockUserRepository.On("GetUserByID", userId).Return(&domain.UserDetail{UserID: userId}, nil)

	// Mock the FindById method with an error
	mockSubsPlanRepository.On("FindById", request.PlanId).Return(nil, fmt.Errorf("plan not found"))

	// Call the function you want to test
	result, err := paymentService.CreatePaymentSubscription(&request, userId)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "error when get subscription plan")

	// Assert that the expected calls were made
	mockUserRepository.AssertExpectations(t)
	mockSubsPlanRepository.AssertExpectations(t)
	mockMidtransCoreApi.AssertExpectations(t)
	mockPaymentRepository.AssertExpectations(t)
}

func TestCreatePaymentSubscription_ChargeTransactionError(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockSubsPlanRepository := new(mocks.SubsPlanRepository)
	mockUserRepository := new(mocks.UserRepository)
	mockSubscriptionService := new(mocks.SubscriptionService)
	mockMidtransCoreApi := new(mocks.MidtransCoreApi)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repositories
	paymentService := NewPaymentService(mockPaymentRepository, mockSubsPlanRepository, mockUserRepository, mockSubscriptionService, mockMidtransCoreApi, mockValidator)

	// Define test data
	request := web.CreatePaymentSubscriptionRequest{
		PlanId:   1,
		BankName: "ASdsad",
	}

	userId := uint(1)

	// Mock the GetUserByID method with no error
	mockUserRepository.On("GetUserByID", userId).Return(&domain.UserDetail{UserID: userId}, nil)

	// Mock the FindById method with no error
	mockSubsPlanRepository.On("FindById", request.PlanId).Return(&domain.SubsPlan{ID: uint(request.PlanId)}, nil)

	// Mock the ChargeTransaction method with an error
	mockMidtransCoreApi.On("ChargeTransaction", mock.Anything).Return(nil, fmt.Errorf("charge transaction error"))

	// Call the function you want to test
	result, err := paymentService.CreatePaymentSubscription(&request, userId)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "error when create payment")

	// Assert that the expected calls were made
	mockUserRepository.AssertExpectations(t)
	mockSubsPlanRepository.AssertExpectations(t)
	mockMidtransCoreApi.AssertExpectations(t)
	mockPaymentRepository.AssertExpectations(t)
}

func TestCreatePaymentSubscription_ConvertPaymentHistoryError(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockSubsPlanRepository := new(mocks.SubsPlanRepository)
	mockUserRepository := new(mocks.UserRepository)
	mockSubscriptionService := new(mocks.SubscriptionService)
	mockMidtransCoreApi := new(mocks.MidtransCoreApi)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repositories
	paymentService := NewPaymentService(mockPaymentRepository, mockSubsPlanRepository, mockUserRepository, mockSubscriptionService, mockMidtransCoreApi, mockValidator)

	// Define test data
	request := web.CreatePaymentSubscriptionRequest{
		PlanId:   1,
		BankName: "ASdsad",
	}

	userId := uint(1)

	// Mock the GetUserByID method with no error
	mockUserRepository.On("GetUserByID", userId).Return(&domain.UserDetail{UserID: userId}, nil)

	// Mock the FindById method with no error
	mockSubsPlanRepository.On("FindById", request.PlanId).Return(&domain.SubsPlan{ID: uint(request.PlanId)}, nil)

	// Mock the ChargeTransaction method with no error
	mockMidtransCoreApi.On("ChargeTransaction", mock.Anything).Return(&coreapi.ChargeResponse{}, nil)


	// Call the function you want to test
	result, err := paymentService.CreatePaymentSubscription(&request, userId)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "error when convert payment history")

	// Assert that the expected calls were made
	mockUserRepository.AssertExpectations(t)
	mockSubsPlanRepository.AssertExpectations(t)
	mockMidtransCoreApi.AssertExpectations(t)
	mockPaymentRepository.AssertExpectations(t)
}

func TestCreatePaymentSubscription_CreatePaymentHistoryError(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockSubsPlanRepository := new(mocks.SubsPlanRepository)
	mockUserRepository := new(mocks.UserRepository)
	mockSubscriptionService := new(mocks.SubscriptionService)
	mockMidtransCoreApi := new(mocks.MidtransCoreApi)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repositories
	paymentService := NewPaymentService(mockPaymentRepository, mockSubsPlanRepository, mockUserRepository, mockSubscriptionService, mockMidtransCoreApi, mockValidator)

	// Define test data
	request := web.CreatePaymentSubscriptionRequest{
		PlanId:   1,
		BankName: "ASdsad",
	}

	userId := uint(1)

	// Mock the GetUserByID method with no error
	mockUserRepository.On("GetUserByID", userId).Return(&domain.UserDetail{UserID: userId}, nil)

	// Mock the FindById method with no error
	mockSubsPlanRepository.On("FindById", request.PlanId).Return(&domain.SubsPlan{ID: uint(request.PlanId)}, nil)

	// Mock the ChargeTransaction method with no error
	mockMidtransCoreApi.On("ChargeTransaction", mock.Anything).Return(&coreapi.ChargeResponse{
		TransactionTime: "2006-01-02 15:04:05", 
		ExpiryTime: "2006-01-02 15:04:05",
		VaNumbers: []coreapi.VANumber{
			{Bank: "adsa", VANumber: "aSDAD"},
		} ,
	}, nil)

	mockPaymentRepository.On("CreatePaymentHistory", mock.Anything).Return(fmt.Errorf("create payment history error"))

	// Call the function you want to test
	result, err := paymentService.CreatePaymentSubscription(&request, userId)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "error when create payment history")

	// Assert that the expected calls were made
	mockUserRepository.AssertExpectations(t)
	mockSubsPlanRepository.AssertExpectations(t)
	mockMidtransCoreApi.AssertExpectations(t)
	mockPaymentRepository.AssertExpectations(t)
}

func TestCreatePaymentSubscription_GetPaymentHistoryByIdError(t *testing.T) {
	mockPaymentRepository := new(mocks.PaymentRepository)
	mockSubsPlanRepository := new(mocks.SubsPlanRepository)
	mockUserRepository := new(mocks.UserRepository)
	mockSubscriptionService := new(mocks.SubscriptionService)
	mockMidtransCoreApi := new(mocks.MidtransCoreApi)
	mockValidator := validator.New()

	// Create a PaymentServiceImpl with the mock repositories
	paymentService := NewPaymentService(mockPaymentRepository, mockSubsPlanRepository, mockUserRepository, mockSubscriptionService, mockMidtransCoreApi, mockValidator)

	// Define test data
	request := web.CreatePaymentSubscriptionRequest{
		PlanId:   1,
		BankName: "ASdsad",
	}

	userId := uint(1)

	// Mock the GetUserByID method with no error
	mockUserRepository.On("GetUserByID", userId).Return(&domain.UserDetail{UserID: userId}, nil)

	// Mock the FindById method with no error
	mockSubsPlanRepository.On("FindById", request.PlanId).Return(&domain.SubsPlan{ID: uint(request.PlanId)}, nil)

	// Mock the ChargeTransaction method with no error
	mockMidtransCoreApi.On("ChargeTransaction", mock.Anything).Return(&coreapi.ChargeResponse{
		TransactionTime: "2006-01-02 15:04:05", 
		ExpiryTime: "2006-01-02 15:04:05",
		VaNumbers: []coreapi.VANumber{
			{Bank: "adsa", VANumber: "aSDAD"},
		} ,
	}, nil)

	mockPaymentRepository.On("CreatePaymentHistory", mock.Anything).Return(nil)
	mockPaymentRepository.On("GetPaymentHistoryById", mock.Anything).Return(nil, fmt.Errorf("get payment history error"))

	// Call the function you want to test
	result, err := paymentService.CreatePaymentSubscription(&request, userId)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "error when get payment history")

	// Assert that the expected calls were made
	mockUserRepository.AssertExpectations(t)
	mockSubsPlanRepository.AssertExpectations(t)
	mockMidtransCoreApi.AssertExpectations(t)
	mockPaymentRepository.AssertExpectations(t)
}
