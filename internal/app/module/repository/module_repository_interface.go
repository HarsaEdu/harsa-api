package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type ModuleRepository interface {
	CreateModule(module *domain.Module) error
	GetAllSectionByCourseId(offset, limit int, search string, courseId uint) ([]domain.Section, int64, error)
	GetAllModuleBySecsionId( sectionId uint) (*domain.Section, error)
	CreateSection(section *domain.Section) error
	UpdateOrderModule(order int, moduleExist *domain.Module) error
	UpdateOrderSection(order int, sectionExist *domain.Section) error
	DeleteSection(section *domain.Section) error
	GetByTitleAndCourseId(title string, courseId uint) (*domain.Module, error)
	GetByOrderAndCourseId(order int, courseId uint) (*domain.Module, error)
	GetByTypeAndId(id uint, modulType string) (*domain.Module, error)
	CekIdFromCourse(userId uint, courseId uint, role string) error
	CekIdFromModule(userId uint, moduleId uint, role string) (*domain.Module, error)
	UpdateModule(updateModul *domain.Module, moduleExist *domain.Module) error
	DeleteModule(module *domain.Module) error
	GetByTitleSectionAndCourseId(title string, courseId uint) (*domain.Section, error)
	GetByOrderSectionAndCourseId(order int, courseId uint) (*domain.Section, error)
	UpdateSection(UpdateSection *domain.Section, sectionExist *domain.Section) error
	CekIdFromSection(userId uint, sectionId uint, role string) (*domain.Section, error)
	GetModuleById(id uint) (*domain.Module, error)
}

type ModuleRepositoryImpl struct {
	DB *gorm.DB
}

func NewModuleRepository(db *gorm.DB) ModuleRepository {
	return &ModuleRepositoryImpl{
		DB: db,
	}
}