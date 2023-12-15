package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDelete_Success(t *testing.T) {
	mockChatbotRepo := new(mocks.ChatbotRepository)
		mockUserRepo := new(mocks.UserRepository)
		mockOpenAi := new(mocks.OpenAi)
		validate := validator.New()
	
		// Create a ChatbotServiceImpl with the mock repositories and dependencies
		chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

	threadID := "123"
	userID := uint(1)
	role := "admin"

	// Set up mock expectations for GetTopicById
	mockChatbotRepo.On("GetTopicById", threadID).Return(&domain.UserChatTopic{
		ID:     "123",
		Topic:  "TestTopic",
		UserID: userID,
	}, nil)

	// Set up mock expectations for Delete
	mockChatbotRepo.On("Delete", mock.Anything).Return(nil)

	err := chatbotService.Delete(threadID, userID, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockChatbotRepo.AssertExpectations(t)
}

func TestDelete_TopicNotFound(t *testing.T) {
	mockChatbotRepo := new(mocks.ChatbotRepository)
		mockUserRepo := new(mocks.UserRepository)
		mockOpenAi := new(mocks.OpenAi)
		validate := validator.New()
	
		// Create a ChatbotServiceImpl with the mock repositories and dependencies
		chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

	threadID := "123"
	userID := uint(1)
	role := "admin"

	// Set up mock expectations for GetTopicById
	mockChatbotRepo.On("GetTopicById", threadID).Return(&domain.UserChatTopic{}, nil)

	err := chatbotService.Delete(threadID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, "topic not found", err.Error())

	// Assert that the expected calls were made
	mockChatbotRepo.AssertExpectations(t)
}

func TestDelete_UnauthorizedUser(t *testing.T) {
	mockChatbotRepo := new(mocks.ChatbotRepository)
		mockUserRepo := new(mocks.UserRepository)
		mockOpenAi := new(mocks.OpenAi)
		validate := validator.New()
	
		// Create a ChatbotServiceImpl with the mock repositories and dependencies
		chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

	threadID := "123"
	userID := uint(2) // Different user ID than the one in the topic
	role := "user"

	// Set up mock expectations for GetTopicById
	mockChatbotRepo.On("GetTopicById", threadID).Return(&domain.UserChatTopic{
		ID:     "123",
		Topic:  "TestTopic",
		UserID: 1, // Assume that the topic belongs to user with ID 1
	}, nil)

	err := chatbotService.Delete(threadID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, "you are not allowed to delete this topic", err.Error())

	// Assert that the expected calls were made
	mockChatbotRepo.AssertExpectations(t)
}

func TestDelete_ErrorDeletingTopic(t *testing.T) {
	mockChatbotRepo := new(mocks.ChatbotRepository)
		mockUserRepo := new(mocks.UserRepository)
		mockOpenAi := new(mocks.OpenAi)
		validate := validator.New()
	
		// Create a ChatbotServiceImpl with the mock repositories and dependencies
		chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

	threadID := "123"
	userID := uint(1)
	role := "admin"

	// Set up mock expectations for GetTopicById
	mockChatbotRepo.On("GetTopicById", threadID).Return(&domain.UserChatTopic{
		ID:     "123",
		Topic:  "TestTopic",
		UserID: userID,
	}, nil)

	// Set up mock expectations for Delete
	mockChatbotRepo.On("Delete", mock.Anything).Return(fmt.Errorf("error deleting topic"))

	err := chatbotService.Delete(threadID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, "error when delete topic : error deleting topic", err.Error())

	// Assert that the expected calls were made
	mockChatbotRepo.AssertExpectations(t)
}


