package service

import (
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/assert"
)

func TestGetAllMessagesInThread_Success(t *testing.T) {
	mockChatbotRepo := new(mocks.ChatbotRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockOpenAi := new(mocks.OpenAi)
	validate := validator.New()

	// Create a ChatbotServiceImpl with the mock repositories and dependencies
	chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

	threadID := "123"
	limit := 10
	after := ""
	before := ""

	// Set up mock expectations for GetTopicById
	mockChatbotRepo.On("GetTopicById", threadID).Return(&domain.UserChatTopic{
		ID: "123",
		// Set other fields as needed
	}, nil)

	// Set up mock expectations for GetAllMessagesInThread
	mockOpenAi.On("GetAllMessagesInThread", "123", limit, after, before).Return(openai.MessagesList{}, nil)

	// Call the function you want to test
	response, err := chatbotService.GetAllMessagesInThread(threadID, limit, after, before)

	// Assert the result
	assert.NoError(t, err)
	assert.Nil(t, response)

	// Assert that the expected calls were made
	mockChatbotRepo.AssertExpectations(t)
	mockOpenAi.AssertExpectations(t)
}

func TestGetAllThreadByUserId_Success(t *testing.T) {
	mockChatbotRepo := new(mocks.ChatbotRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockOpenAi := new(mocks.OpenAi)
	validate := validator.New()

	// Create a ChatbotServiceImpl with the mock repositories and dependencies
	chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

	userID := uint(1)

	// Set up mock expectations for GetUserByID
	mockUserRepo.On("GetUserByID", userID).Return(&domain.UserDetail{
		UserID:   userID,
		Username: "TestUser",
		// Set other fields as needed
	}, nil)

	// Set up mock expectations for GetAllTopicByUserId
	mockChatbotRepo.On("GetAllTopicByUserId", userID).Return([]domain.UserChatTopic{
		// Set mock response
	}, nil)

	// Call the function you want to test
	response, err := chatbotService.GetAllThreadByUserId(userID)

	// Assert the result
	assert.NoError(t, err)
	assert.Nil(t, response)

	// Assert that the expected calls were made
	mockUserRepo.AssertExpectations(t)
	mockChatbotRepo.AssertExpectations(t)
}

// Add more test cases for error scenarios, e.g., when a user is not found, when there is an error in repository methods, etc.
