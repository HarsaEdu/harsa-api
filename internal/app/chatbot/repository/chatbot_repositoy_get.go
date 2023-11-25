package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (chatbotRepository *ChatbotRepositoryImpl) GetTopicById(id string) (*domain.UserChatTopic, error) {
	userChatTopic := domain.UserChatTopic{}

	result := chatbotRepository.DB.Where("id = ?", id).First(&userChatTopic)
	if result.Error != nil {
		return nil, result.Error
	}

	return &userChatTopic, nil
}

func (chatbotRepository *ChatbotRepositoryImpl) GetAllTopicByUserId(userId uint) ([]domain.UserChatTopic, error) {
	userChatTopics := []domain.UserChatTopic{}

	result := chatbotRepository.DB.
			Joins("JOIN users ON user_chat_topics.user_id = users.id").
			Where("user_chat_topics.user_id = ?", userId).
			Order("user_chat_topics.updated_at DESC").
			Find(&userChatTopics)

	if result.Error != nil {
			return nil, result.Error
	}

	return userChatTopics, nil
}
