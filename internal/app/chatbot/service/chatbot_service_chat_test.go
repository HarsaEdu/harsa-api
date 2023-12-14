package service

import (
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
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
	mockChatbotRepo.On("GetTopicById", threadId).Return(&domain.UserChatTopic{}, nil)

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

