package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (chatbotRepository *ChatbotRepositoryImpl) GetTopicById(id string) (*domain.UserChatTopic, error) {
	userChatTopic := domain.UserChatTopic{}

	result := chatbotRepository.DB.Where("id = ?", id).First(&userChatTopic)
	if result.Error != nil {
		return nil, result.Error
	}

	
	return &userChatTopic, nil
}

func (chatbotRepository *ChatbotRepositoryImpl) CreateUserChatTopic(userChatTopic *domain.UserChatTopic) error {
	result := chatbotRepository.DB.Create(userChatTopic)
	if result.Error != nil {
		return result.Error
	}

	return nil
}