package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"gorm.io/gorm"
)

type ChatbotRepository interface {
	CreateUserChatTopic(userChatTopic *domain.UserChatTopic) error
	GetTopicById(id string) (*domain.UserChatTopic, error)
	GetAllTopicByUserId(userId uint) ([]domain.UserChatTopic, error)
	Delete(topic *domain.UserChatTopic) error
	UpdateTopic(topicId string, request *web.CreateThreadRequest) error
}

type ChatbotRepositoryImpl struct {
	DB *gorm.DB
}

func NewChatbotRepository(db *gorm.DB) ChatbotRepository {
	return &ChatbotRepositoryImpl{DB: db}
}