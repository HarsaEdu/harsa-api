package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (chatbotService *ChatbotServiceImpl) GetAllMessagesInThread(threadId string, limit int, after string, before string) ([]web.GetMessageInThreadResponse, error) {
	existingThread, err := chatbotService.ChatbotRepository.GetTopicById(threadId)
	if err != nil {
		return nil, fmt.Errorf("error when get thread : %s", err.Error())
	}

	openAiResponse, err := chatbotService.OpenAi.GetAllMessagesInThread(existingThread.ID, limit, after, before)
	if err != nil {
		return nil, fmt.Errorf("error when get all message : %s", err.Error())
	}

	response := conversion.OpenAiThreadMessagesResponseToGetAllMessageInThreadResponse(openAiResponse)

	return response, nil
}

func (chatbotService *ChatbotServiceImpl) GetAllThreadByUserId(userId uint) ([]web.GetAllThreadByUserIdResponse, error) {
	existinUser, err := chatbotService.UserRepository.GetUserByID(userId)
	if err != nil {
		return nil, fmt.Errorf("error when get user : %s", err.Error())
	}

	if existinUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	result, err := chatbotService.ChatbotRepository.GetAllTopicByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("error when get all thread : %s", err.Error())
	}

	response := conversion.UserChatDomainToGetAllThreadByUserIdResponse(result)

	return response, nil
}