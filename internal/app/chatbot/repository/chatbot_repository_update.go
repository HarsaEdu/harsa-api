package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (chatbotRepository *ChatbotRepositoryImpl) UpdateTopic(id string, topic *web.CreateThreadRequest) error {

	result := chatbotRepository.DB.Where("id = ?",  "%" + id + "%").Updates(&domain.UserChatTopic{Topic: topic.Topic})
	if result.Error != nil {
		return result.Error
	}

	return nil
}