package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (chatbotRepository *ChatbotRepositoryImpl) Delete(topic *domain.UserChatTopic) error {

	if err := chatbotRepository.DB.Delete(&topic).Error; err != nil {
		return  err
	}

	return nil
}