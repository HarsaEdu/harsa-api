package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (courseRepository *CourseRepositoryImpl) Create(course *domain.Course) error {
	result := courseRepository.DB.Create(&course)
	if result.Error != nil {
		return result.Error
	}

	return nil
}