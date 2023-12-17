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

func TestGetRecommendations_Success(t *testing.T) {
	mockValidator := validator.New()
	mockRecommendationsApi := new(mocks.RecommendationsApi)
	mockUserRepository := new(mocks.UserRepository)

	// Create a RecommendationsServiceImpl with the mock repositories
	recommendationsService := NewRecommendationsService(mockValidator, mockRecommendationsApi, nil, nil, mockUserRepository, nil, nil)

	// Define test data
	request := &web.GetRecommendationsRequest{
		UserId: 1,
		Max:    3,
	}

	// Mock the UserAvailableByID method with no error
	mockUserRepository.On("UserAvailableByID", request.UserId).Return(&domain.User{}, nil)

	// Mock the GetRecommendations method with no error
	mockRecommendationsApi.On("GetRecommendations", request).Return(&web.GetRecommendationsResponse{}, nil)

	// Call the function you want to test
	response, err := recommendationsService.GetRecommendations(request)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, response)

	// Assert that the expected calls were made
	mockUserRepository.AssertExpectations(t)
	mockRecommendationsApi.AssertExpectations(t)
}

func TestGetRecommendations_UserNotFound(t *testing.T) {
	mockValidator := validator.New()
	mockRecommendationsApi := new(mocks.RecommendationsApi)
	mockUserRepository := new(mocks.UserRepository)

	// Create a RecommendationsServiceImpl with the mock repositories
	recommendationsService := NewRecommendationsService(mockValidator, mockRecommendationsApi, nil, nil, mockUserRepository, nil, nil)

	// Define test data
	request := &web.GetRecommendationsRequest{
		UserId: 1,
		Max:    3,
	}

	// Mock the UserAvailableByID method with user not found
	mockUserRepository.On("UserAvailableByID", request.UserId).Return(nil, nil)

	// Call the function you want to test
	response, err := recommendationsService.GetRecommendations(request)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "user not found")

	// Assert that the expected calls were made
	mockUserRepository.AssertExpectations(t)
	mockRecommendationsApi.AssertExpectations(t)
}

func TestGetRecommendations_ValidationError(t *testing.T) {
	mockValidator := validator.New()
	mockRecommendationsApi := new(mocks.RecommendationsApi)
	mockUserRepository := new(mocks.UserRepository)

	// Create a RecommendationsServiceImpl with the mock repositories
	recommendationsService := NewRecommendationsService(mockValidator, mockRecommendationsApi, nil, nil, mockUserRepository, nil, nil)

	// Define test data
	request := &web.GetRecommendationsRequest{
		// Add necessary fields for GetRecommendationsRequest
	}

	// Set up validation error
	mockValidator.RegisterValidation("your_validation_tag", func(fl validator.FieldLevel) bool {
		return false
	})

	// Call the function you want to test
	response, err := recommendationsService.GetRecommendations(request)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "Key: 'GetRecommendationsRequest.UserId'")

	// Assert that the expected calls were made
	mockUserRepository.AssertExpectations(t)
	mockRecommendationsApi.AssertExpectations(t)
}

func TestGetRecommendationsForInstructor_Success(t *testing.T) {
	mockInterestRepository := new(mocks.InterestRepository)
	mockUserRepository := new(mocks.UserRepository)
	mockOpenAi := new(mocks.OpenAi)
	mockNotificationRepository := new(mocks.NotificationRepository)
	mockFirebase := new(mocks.Firebase)

	// Create a RecommendationsServiceImpl with the mock repositories
	recommendationsService := NewRecommendationsService(nil, nil, mockOpenAi, mockFirebase, mockUserRepository, mockInterestRepository, mockNotificationRepository)

	// Define test data
	mockTopInterests := []domain.UserInterest{
		{ID: uint(1), ProfileID: uint(3), CategoryID: uint(4)},
		{ID: uint(2), ProfileID: uint(3), CategoryID: uint(4)},
		// Add more interests as needed
	}

	mockInstructors := []domain.UserRegistrationToken{
		{UserID: uint(1), RegistrationToken: "token1"},
		{UserID: uint(2), RegistrationToken: "token2"},
		// Add more instructors as needed
	}

	// Mock the GetTopInterests method with no error
	mockInterestRepository.On("GetTopInterests", 5).Return(mockTopInterests, nil)

	// Mock the GetUsersRegistrationToken method with no error
	mockUserRepository.On("GetUsersRegistrationToken", 2).Return(mockInstructors, nil)

	// Mock the GetChatCompletion method with no error
	mockOpenAi.On("GetChatCompletion", mock.Anything, mock.Anything).Return("ChatResponse", nil)

	// Mock the CreateMany method with no error
	mockNotificationRepository.On("CreateMany", mock.Anything).Return(nil)

	// Mock the SendNotificationMulticast method with no error
	mockFirebase.On("SendNotificationMulticast", mock.Anything).Return(nil)

	// Call the function you want to test
	err := recommendationsService.GetRecommendationsForInstructor()

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockInterestRepository.AssertExpectations(t)
	mockUserRepository.AssertExpectations(t)
	mockOpenAi.AssertExpectations(t)
	mockNotificationRepository.AssertExpectations(t)
	mockFirebase.AssertExpectations(t)
}
