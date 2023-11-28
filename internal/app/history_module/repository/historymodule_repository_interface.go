package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type HistoryModuleRepository interface {
	Create(historyModule *domain.HistoryModule) error
	GetAll(offset, limit int, search string) ([]domain.HistoryModule, int64, error)
	GetById(id int) (*domain.HistoryModule, error)
	Update(id int, historyModule *domain.HistoryModule) error
}

type HistoryModuleRepositoryImpl struct {
	DB *gorm.DB
}

func NewHistoryModuleRepository(db *gorm.DB) HistoryModuleRepository {
	return &HistoryModuleRepositoryImpl{
		DB: db,
	}
}
