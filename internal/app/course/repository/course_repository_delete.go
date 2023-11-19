package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (courseRepository *CourseRepositoryImpl) Delete(id uint) error {
	result := courseRepository.DB.Where("id = ?", id).Delete(&domain.Course{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
