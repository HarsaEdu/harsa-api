package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

//	type HistorySubModuleRepository interface {
//		GetHistorySubModuleByUserID(request *web.GetHistorySubModuleRequest) (*domain.HistorySubModule, error)
//		CreateHistorySubModule(request *web.CreateHistorySubModuleRequest, userID uint) error
//		UpdateHistorySubModule(request *web.UpdateHistorySubModuleRequest) error
//	}
type HistorySubModuleRepository interface {
	GetHistorySubModuleByUserID(userID uint) ([]domain.HistorySubModule, int64, error)
	GetHistorySubModuleByID(id uint) (*domain.HistorySubModule, error)
	CreateHistorySubModule(request *domain.HistorySubModule) error
	UpdateHistorySubModule(request *domain.HistorySubModule, id uint) error
}

type HistorySubModuleRepositoryImpl struct {
	DB *gorm.DB
}

func NewHistorySubModuleRepository(db *gorm.DB) HistorySubModuleRepository {
	return &HistorySubModuleRepositoryImpl{
		DB: db,
	}
}
