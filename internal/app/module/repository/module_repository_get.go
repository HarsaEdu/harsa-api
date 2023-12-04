package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

func (moduleRepository *ModuleRepositoryImpl) GetAllSectionByCourseId(offset, limit int, search string, courseId uint) ([]domain.Section, int64, error) {
	var section []domain.Section
	var total int64

	query := moduleRepository.DB.Preload("Modules", func(db *gorm.DB) *gorm.DB {
        return db.Order("ABS(order_by) ASC, ABS(id) ASC")
    }).
	Where("course_id = ?", courseId).Order("ABS(order_by) ASC, ABS(id) ASC")

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	query.Find(&section).Count(&total)

	query = query.Offset(offset).Limit(limit)

	result := query.Find(&section)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if total == 0 {
		return nil, 0, nil
	}

	return section, total, nil
}

func (moduleRepository *ModuleRepositoryImpl) GetAllModuleBySecsionId( sectionId uint) (*domain.Section, error) {
	var section domain.Section

	if err := moduleRepository.DB.Preload("Modules", func(db *gorm.DB) *gorm.DB {
        return db.Order("ABS(order_by) ASC, ABS(id) ASC")
    }).First(&section, sectionId).Error; err != nil {
		return nil, err
	}

	return &section, nil
}

func (moduleRepository *ModuleRepositoryImpl) GetByTitleAndSectionId(title string, sectionId uint) (*domain.Module, error) {
	module := domain.Module{}
	result := moduleRepository.DB.Where("title = ? AND section_id = ?", title, sectionId).First(&module)
	if result.Error != nil {
		return nil, result.Error
	}

	return &module, nil
}

func (moduleRepository *ModuleRepositoryImpl) GetByOrderAndCourseId(order int, courseId uint) (*domain.Module, error) {
	module := domain.Module{}
	result := moduleRepository.DB.Where("order = ? AND course_id = ?", order, courseId).First(&module)
	if result.Error != nil {
		return nil, result.Error
	}

	return &module, nil
}

func (moduleRepository *ModuleRepositoryImpl) GetByTypeAndId(id uint, modulType string) (*domain.Module, error) {
	module := domain.Module{}
	result := moduleRepository.DB.Where("id = ? AND type = ?", id, modulType).First(&module)
	if result.Error != nil {
		return nil, result.Error
	}

	return &module, nil
}

func (moduleRepository *ModuleRepositoryImpl) GetModuleById(id uint) (*domain.Module, error) {
	module := domain.Module{}
	result := moduleRepository.DB.Preload("SubModules").Preload("Quizzes").Preload("Submissions").First(&module, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &module, nil
}


func (moduleRepository *ModuleRepositoryImpl) GetByTitleSectionAndCourseId(title string, courseId uint) (*domain.Section, error) {
	section := domain.Section{}
	result := moduleRepository.DB.Where("title = ? AND course_id = ?", title, courseId).First(&section)
	if result.Error != nil {
		return nil, result.Error
	}

	return &section, nil
}

func (moduleRepository *ModuleRepositoryImpl) GetByOrderSectionAndCourseId(order int, courseId uint) (*domain.Section, error) {
	section := domain.Section{}
	result := moduleRepository.DB.Where("order = ? AND course_id = ?", order, courseId).First(&section)
	if result.Error != nil {
		return nil, result.Error
	}

	return &section, nil
}



