package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (chatbotService *ChatbotServiceImpl) ChatWithAssistant(threadId string, request *web.ChatWithAssistantRequest) (string, error) {
	err := chatbotService.Validate.Struct(request)
	if err != nil {
		return "", err
	}

	existingTopic, err := chatbotService.ChatbotRepository.GetTopicById(threadId)
	if err != nil {
		return "", fmt.Errorf("error when get topic by id : %s", err.Error())
	}

	if existingTopic.Topic == "" {
		return "", fmt.Errorf("topic not found")
	}

	chatResponse, err := chatbotService.OpenAi.ChatWithAssistant(threadId, request.Message)
	if err != nil {
		return "", fmt.Errorf("error when chat with assistant : %s", err.Error())
	}

	return chatResponse, nil
}