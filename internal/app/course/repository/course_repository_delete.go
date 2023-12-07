package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (courseRepository *CourseRepositoryImpl) Delete(course *domain.Course) error {
	
	if err := courseRepository.DB.Delete(&course).Error; err != nil {
		return  err
	}

	return nil
}
