package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type ModuleRepository interface {
	Create(module *domain.Module) error
	GetAllByCourseId(offset, limit int, search string, courseId uint) ([]domain.Module, int64, error)
	// GetById(id uint) (*domain.Module, error)
	GetByTitleAndCourseId(title string, courseId uint) (*domain.Module, error)
	GetByOrderAndCourseId(order int, courseId uint) (*domain.Module, error)
	GetByTypeAndId(id uint, modulType string) (*domain.Module, error)
	CekIdFromCourse(userId uint, courseId uint, role string) error
	CekIdFromModule(userId uint, moduleId uint, role string) (*domain.Module, error)
	Update(updateModul *domain.Module, moduleExist *domain.Module) error
	Delete(module *domain.Module) error
	// Update(id uint, module *domain.Module) error
	// UpdateImage(module *domain.Module) error
	// Delete(id uint) error
}

type ModuleRepositoryImpl struct {
	DB *gorm.DB
}

func NewModuleRepository(db *gorm.DB) ModuleRepository {
	return &ModuleRepositoryImpl{
		DB: db,
	}
}