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

func TestCreateThread_Success(t *testing.T) {
	// Create mock repositories and dependencies
		// Create mock repositories and dependencies
		mockChatbotRepo := new(mocks.ChatbotRepository)
		mockUserRepo := new(mocks.UserRepository)
		mockOpenAi := new(mocks.OpenAi)
		validate := validator.New()
	
		// Create a ChatbotServiceImpl with the mock repositories and dependencies
		chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

		// Define test data
		userID := uint(1)
		request := &web.CreateThreadRequest{
			Topic: "TestTopic",
			// Add other fields as needed
		}

	// Set up mock expectations for GetUserByID
	mockUserRepo.On("GetUserByID", userID).Return(&domain.UserDetail{
		UserID: uint(6),
		Username: "UserName",
	}, nil)

	// Set up mock expectations for CreateThread and CreateUserChatTopic
	mockOpenAi.On("CreateThread", "UserName", request.Topic).Return("ThreadId", nil)
	mockChatbotRepo.On("CreateUserChatTopic", mock.Anything).Return(nil)

	// Call the function you want to test
	err := chatbotService.CreateThread(request, userID)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockUserRepo.AssertExpectations(t)
	mockOpenAi.AssertExpectations(t)
	mockChatbotRepo.AssertExpectations(t)
}

func TestCreateThread_UserNotFound(t *testing.T) {
	mockChatbotRepo := new(mocks.ChatbotRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockOpenAi := new(mocks.OpenAi)
	validate := validator.New()

	chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

	userID := uint(1)
	request := &web.CreateThreadRequest{
		Topic: "TestTopic",
	}

	// Set up mock expectations for GetUserByID
	mockUserRepo.On("GetUserByID", userID).Return(nil, fmt.Errorf("user not found"))

	err := chatbotService.CreateThread(request, userID)

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())

	mockUserRepo.AssertExpectations(t)
	mockOpenAi.AssertExpectations(t)
	mockChatbotRepo.AssertExpectations(t)
}

func TestCreateThread_ErrorCreatingThread(t *testing.T) {
	mockChatbotRepo := new(mocks.ChatbotRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockOpenAi := new(mocks.OpenAi)
	validate := validator.New()

	chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

	userID := uint(1)
	request := &web.CreateThreadRequest{
		Topic: "TestTopic",
	}

	mockUserRepo.On("GetUserByID", userID).Return(&domain.UserDetail{
		UserID:   uint(6),
		Username: "UserName",
	}, nil)

	mockOpenAi.On("CreateThread", "UserName", request.Topic).Return("", fmt.Errorf("error creating thread"))

	err := chatbotService.CreateThread(request, userID)

	assert.Error(t, err)
	assert.Equal(t, "error when creating thread : error creating thread", err.Error())

	mockUserRepo.AssertExpectations(t)
	mockOpenAi.AssertExpectations(t)
	mockChatbotRepo.AssertExpectations(t)
}

func TestCreateThread_ErrorCreatingUserChatTopic(t *testing.T) {
	mockChatbotRepo := new(mocks.ChatbotRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockOpenAi := new(mocks.OpenAi)
	validate := validator.New()

	chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

	userID := uint(1)
	request := &web.CreateThreadRequest{
		Topic: "TestTopic",
	}

	mockUserRepo.On("GetUserByID", userID).Return(&domain.UserDetail{
		UserID:   uint(6),
		Username: "UserName",
	}, nil)

	mockOpenAi.On("CreateThread", "UserName", request.Topic).Return("ThreadId", nil)
	mockChatbotRepo.On("CreateUserChatTopic", mock.Anything).Return(fmt.Errorf("error creating user chat topic"))

	err := chatbotService.CreateThread(request, userID)

	assert.Error(t, err)
	assert.Equal(t, "error when creating thread : error creating user chat topic", err.Error())

	mockUserRepo.AssertExpectations(t)
	mockOpenAi.AssertExpectations(t)
	mockChatbotRepo.AssertExpectations(t)
}

