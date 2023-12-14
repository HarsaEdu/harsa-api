package service

import (
	"fmt"
)

func (chatbotService *ChatbotServiceImpl) Delete(threadId string, userID uint, role string)  error {

	existingTopic, err := chatbotService.ChatbotRepository.GetTopicById(threadId)
	if err != nil {
		return fmt.Errorf("error when cek topic by id : %s", err.Error())
	}

	if existingTopic.Topic == "" {
		return  fmt.Errorf("topic not found")
	}

	if existingTopic.UserID != userID && role != "admin"{
		return  fmt.Errorf("you are not allowed to delete this topic") 
	}

	err = chatbotService.ChatbotRepository.Delete(existingTopic)
	if err != nil {
		return  fmt.Errorf("error when delete topic : %s", err.Error()) 
	}

	return  nil
}