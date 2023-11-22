package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (chatbotRepository *ChatbotRepositoryImpl) CreateThread(userChatTopic *domain.UserChatTopic) error {
	result := chatbotRepository.DB.Create(userChatTopic)
	if result.Error != nil {
		return result.Error
	}

	return nil
}