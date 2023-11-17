package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (courseRepository *CourseRepositoryImpl) UpdateImage(course *domain.Course) error {
	result := courseRepository.DB.Model(&domain.Course{}).Where("id = ?", course.ID).Update("image_url", course.ImageUrl)
	if result.Error != nil {
		return result.Error
	}

	return nil
}