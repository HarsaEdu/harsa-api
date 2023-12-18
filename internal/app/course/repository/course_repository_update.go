package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (courseRepository *CourseRepositoryImpl) Update(id uint, course *domain.Course) error {
	result := courseRepository.DB.Where("id = ?", id).Updates(&domain.Course{
		Title: course.Title,
		Description: course.Description,
		CategoryID: course.CategoryID,
		UserID: course.UserID,
		ImageUrl: course.ImageUrl,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}