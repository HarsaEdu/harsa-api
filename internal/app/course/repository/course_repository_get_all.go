package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (courseRepository *CourseRepositoryImpl) GetAll() ([]domain.Course, error) {
	var courses []domain.Course
	var count int64
	result := courseRepository.DB.Preload("User.Role").Preload("Category").Find(&courses).Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	if count == 0 {
		return nil, nil
	}

	return courses, nil
}