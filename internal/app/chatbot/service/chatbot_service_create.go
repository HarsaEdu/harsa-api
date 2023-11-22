package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (chatbotService *ChatbotServiceImpl) CreateThread(request *web.CreateThreadRequest, userId uint) error {
	err := chatbotService.Validate.Struct(request)
	if err != nil {
		return err
	}

	exisitingUser, _ := chatbotService.UserRepository.GetUserByID(userId)
	if exisitingUser == nil {
		return fmt.Errorf("user not found")
	}

	var studentName string

	if exisitingUser.FirstName == "" {
		studentName = exisitingUser.Username
	} else {
		studentName = exisitingUser.FirstName
	}

	threadId, err := chatbotService.OpenAi.CreateThread(studentName, request.Topic)
	if err != nil {
		return fmt.Errorf("error when creating thread : %s", err.Error())
	}

	if threadId == "" {
		return fmt.Errorf("error when creating thread : thread is not exist")
	}

	topic := conversion.CreateThreadRequestToUserChatTopicDomain(request, threadId)

	err = chatbotService.ChatbotRepository.CreateThread(topic)
	if err != nil {
		return fmt.Errorf("error when creating thread : %s", err.Error())
	}

	return nil
}