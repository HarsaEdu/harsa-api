package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (moduleRepository *ModuleRepositoryImpl) GetAllByCourseId(offset, limit int, search string, courseId uint) ([]domain.Module, int64, error) {
	var modules []domain.Module
	var total int64

	query := moduleRepository.DB.Preload("SubModules").Preload("Quizzes").Preload("Submissions").Where("course_id = ?", courseId)

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	query.Find(&modules).Count(&total)

	query = query.Offset(offset).Limit(limit)

	result := query.Find(&modules)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if total == 0 {
		return nil, 0, nil
	}

	return modules, total, nil
}

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