package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (chatbotService *ChatbotServiceImpl) Update(threadId string, userID uint, update *web.CreateThreadRequest , role string) error {
	err := chatbotService.Validate.Struct(update)
	if err != nil {
		return err
	}
	existingTopic, err := chatbotService.ChatbotRepository.GetTopicById(threadId)
	if err != nil {
		return fmt.Errorf("error when cek topic by id : %s", err.Error())
	}

	if existingTopic.Topic == "" {
		return fmt.Errorf("topic not found")
	}

	if existingTopic.UserID != userID && role != "admin"{
		return  fmt.Errorf("you are not allowed to delete this topic") 
	}

	err = chatbotService.ChatbotRepository.UpdateTopic(threadId, update)
	if err != nil {
		return fmt.Errorf("error when update topic : %s", err.Error())
	}

	return nil
}