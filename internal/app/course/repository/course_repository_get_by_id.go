package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (courseRepository *CourseRepositoryImpl) GetById(id uint) (*domain.Course, error) {
	course := &domain.Course{}

	result := courseRepository.DB.Preload("User.Role").Preload("Category").First(&course, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return course, nil
}