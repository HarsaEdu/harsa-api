package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (moduleRepository *ModuleRepositoryImpl) GetByTitleAndCourseId(title string, courseId uint) (*domain.Module, error) {
	module := domain.Module{}
	result := moduleRepository.DB.Where("title = ? AND course_id = ?", title, courseId).First(&module)
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