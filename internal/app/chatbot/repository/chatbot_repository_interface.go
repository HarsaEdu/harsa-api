package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type ChatbotRepository interface {
	CreateUserChatTopic(userChatTopic *domain.UserChatTopic) error
}

type ChatbotRepositoryImpl struct {
	DB *gorm.DB
}

func NewChatbotRepository(db *gorm.DB) ChatbotRepository {
	return &ChatbotRepositoryImpl{DB: db}
}