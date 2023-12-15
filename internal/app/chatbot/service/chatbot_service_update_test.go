package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestUpdate_Success(t *testing.T) {
	mockChatbotRepo := new(mocks.ChatbotRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockOpenAi := new(mocks.OpenAi)
	validate := validator.New()

	// Create a ChatbotServiceImpl with the mock repositories and dependencies
	chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

	threadID := "123"
	userID := uint(1)
	role := "admin"
	updateRequest := &web.CreateThreadRequest{
		Topic: "UpdatedTopic",
		// Add other fields as needed
	}

	// Set up mock expectations for GetTopicById
	mockChatbotRepo.On("GetTopicById", threadID).Return(&domain.UserChatTopic{
		ID:     "123",
		Topic:  "TestTopic",
		UserID: userID,
		// Set other fields as needed
	}, nil)

	// Set up mock expectations for UpdateTopic
	mockChatbotRepo.On("UpdateTopic", threadID, updateRequest).Return(nil)

	// Call the function you want to test
	err := chatbotService.Update(threadID, userID, updateRequest, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockChatbotRepo.AssertExpectations(t)
}

func TestUpdate_TopicNotFound(t *testing.T) {
	mockChatbotRepo := new(mocks.ChatbotRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockOpenAi := new(mocks.OpenAi)
	validate := validator.New()

	// Create a ChatbotServiceImpl with the mock repositories and dependencies
	chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)
	threadID := "123"
	userID := uint(1)
	role := "admin"
	updateRequest := &web.CreateThreadRequest{
		Topic: "UpdatedTopic",
		// Add other fields as needed
	}

	// Set up mock expectations for GetTopicById
	mockChatbotRepo.On("GetTopicById", threadID).Return(&domain.UserChatTopic{}, nil)

	err := chatbotService.Update(threadID, userID, updateRequest, role)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, "topic not found", err.Error())

	// Assert that the expected calls were made
	mockChatbotRepo.AssertExpectations(t)
}

func TestUpdate_UnauthorizedUser(t *testing.T) {
	mockChatbotRepo := new(mocks.ChatbotRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockOpenAi := new(mocks.OpenAi)
	validate := validator.New()

	// Create a ChatbotServiceImpl with the mock repositories and dependencies
	chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

	threadID := "123"
	userID := uint(2) // Different user ID than the one in the topic
	role := "user"
	updateRequest := &web.CreateThreadRequest{
		Topic: "UpdatedTopic",
		// Add other fields as needed
	}

	// Set up mock expectations for GetTopicById
	mockChatbotRepo.On("GetTopicById", threadID).Return(&domain.UserChatTopic{
		ID:     "123",
		Topic:  "TestTopic",
		UserID: 1, // Assume that the topic belongs to user with ID 1
		// Set other fields as needed
	}, nil)

	err := chatbotService.Update(threadID, userID, updateRequest, role)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, "you are not allowed to delete this topic", err.Error())

	// Assert that the expected calls were made
	mockChatbotRepo.AssertExpectations(t)
}

func TestUpdate_ErrorUpdatingTopic(t *testing.T) {
	mockChatbotRepo := new(mocks.ChatbotRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockOpenAi := new(mocks.OpenAi)
	validate := validator.New()

	// Create a ChatbotServiceImpl with the mock repositories and dependencies
	chatbotService := NewChatbotService(mockChatbotRepo, mockUserRepo, validate, mockOpenAi)

	threadID := "123"
	userID := uint(1)
	role := "admin"
	updateRequest := &web.CreateThreadRequest{
		Topic: "UpdatedTopic",
		// Add other fields as needed
	}

	// Set up mock expectations for GetTopicById
	mockChatbotRepo.On("GetTopicById", threadID).Return(&domain.UserChatTopic{
		ID:     "123",
		Topic:  "TestTopic",
		UserID: userID,
		// Set other fields as needed
	}, nil)

	// Set up mock expectations for UpdateTopic
	mockChatbotRepo.On("UpdateTopic", threadID, updateRequest).Return(fmt.Errorf("error updating topic"))

	err := chatbotService.Update(threadID, userID, updateRequest, role)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, "error when update topic : error updating topic", err.Error())

	// Assert that the expected calls were made
	mockChatbotRepo.AssertExpectations(t)
}
