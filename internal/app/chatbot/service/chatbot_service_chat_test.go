package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatWithAssistant(t *testing.T) {
	// Create mock repositories and dependencies
	mockChatbotRepo := new(mocks.ChatbotRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockOpenAi := new(mocks.OpenAi)
	validate := validator.New()

	// Create a ChatbotServiceImpl with the mock repositories and dependencies
	chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

	// Define test data
	threadId := "123"
	request := &web.ChatWithAssistantRequest{
		Message: "Hello, assistant!",
	}

	// Set up mock expectations for GetTopicById
	mockChatbotRepo.On("GetTopicById", threadId).Return(&domain.UserChatTopic{
		ID: "sadasdasds",
		Topic: "SADADAD",
	}, nil)

	// Set up mock expectations for ChatWithAssistant
	mockOpenAi.On("ChatWithAssistant", threadId, request.Message).Return("Assistant's response", nil)

	// Call the function you want to test
	chatResponse, err := chatbotService.ChatWithAssistant(threadId, request)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, "Assistant's response", chatResponse)

	// Assert that the expected calls were made
	mockChatbotRepo.AssertExpectations(t)
	mockOpenAi.AssertExpectations(t)
}

func TestChatWithAssistantNotfound(t *testing.T) {
	// Create mock repositories and dependencies
	mockChatbotRepo := new(mocks.ChatbotRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockOpenAi := new(mocks.OpenAi)
	validate := validator.New()

	// Create a ChatbotServiceImpl with the mock repositories and dependencies
	chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

	// Define test data
	threadId := "123"
	request := &web.ChatWithAssistantRequest{
		Message: "Hello, assistant!",
	}

	// Set up mock expectations for GetTopicById
	mockChatbotRepo.On("GetTopicById", threadId).Return(&domain.UserChatTopic{
		ID: "sadasdasds",
	}, nil)

	// Set up mock expectations for ChatWithAssistant
	mockOpenAi.On("ChatWithAssistant", threadId, request.Message).Return("Assistant's response", nil)

	// Call the function you want to test
	chatResponse, err := chatbotService.ChatWithAssistant(threadId, request)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, "", chatResponse)

	// Assert that the expected calls were made
	mockChatbotRepo.AssertExpectations(t)
}

func TestChatWithAssistantvalidation(t *testing.T) {
	// Create mock repositories and dependencies
	mockChatbotRepo := new(mocks.ChatbotRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockOpenAi := new(mocks.OpenAi)
	validate := validator.New()

	// Create a ChatbotServiceImpl with the mock repositories and dependencies
	chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

	// Define test data
	threadId := "123"
	request := &web.ChatWithAssistantRequest{
		
	}

	// Call the function you want to test
	chatResponse, err := chatbotService.ChatWithAssistant(threadId, request)

	// Assert the result
	require.Error(t, err)
	assert.Equal(t, "", chatResponse)

	// Assert that the expected calls were made
	mockChatbotRepo.AssertExpectations(t)
	mockOpenAi.AssertExpectations(t)
}

func TestChatWithAssistant_ErrorCases(t *testing.T) {
    // Create mock repositories and dependencies
    mockChatbotRepo := new(mocks.ChatbotRepository)
    mockUserRepo := new(mocks.UserRepository)
    mockOpenAi := new(mocks.OpenAi)
    validate := validator.New()

    // Create a ChatbotServiceImpl with the mock repositories and dependencies
    chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

    // Define test data
    threadId := "123"
    request := &web.ChatWithAssistantRequest{
        Message: "Hello, assistant!",
    }

    // Test case 1: Error in GetTopicById
    mockChatbotRepo.On("GetTopicById", threadId).Return(nil, fmt.Errorf("topic not found"))

    chatResponse, err := chatbotService.ChatWithAssistant(threadId, request)

    // Assert the result
    assert.Error(t, err)
    assert.Equal(t, "", chatResponse) // Check that the response is empty on error

    // Test case 2: Topic not found
    mockChatbotRepo.On("GetTopicById", threadId).Return(&domain.UserChatTopic{}, nil)

    chatResponse, err = chatbotService.ChatWithAssistant(threadId, request)

    // Assert the result
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "topic not found")
    assert.Equal(t, "", chatResponse)

    // Test case 3: Error in ChatWithAssistant
    mockChatbotRepo.On("GetTopicById", threadId).Return(&domain.UserChatTopic{
        ID:    "sadasdasds",
        Topic: "SADADAD",
    }, nil)
    mockOpenAi.On("ChatWithAssistant", threadId, request.Message).Return("", fmt.Errorf("some error"))

    chatResponse, err = chatbotService.ChatWithAssistant(threadId, request)

    // Assert the result
    assert.Error(t, err)
    assert.Equal(t, "", chatResponse)
}
