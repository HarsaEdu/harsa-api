package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (chatbotRepository *ChatbotRepositoryImpl) UpdateTopic(topicId string, request *web.CreateThreadRequest) error {

	result := chatbotRepository.DB.Where("id = ?",  topicId).Updates(&domain.UserChatTopic{Topic: request.Topic})
	if result.Error != nil {
		return result.Error
	}

	return nil
}