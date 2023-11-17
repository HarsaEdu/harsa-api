package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (courseRepository *CourseRepositoryImpl) GetAll(offset, limit int, search string) ([]domain.Course, int64, error) {
	var courses []domain.Course
	var count int64

	query := courseRepository.DB.Preload("User.Role").Preload("Category")

	if search != "" {
		searchQuery := "%" + search + "%"
		query = query.Where("title LIKE ?", searchQuery)
	}

	query.Find(&courses).Count(&count)

	query = query.Offset(offset).Limit(limit)

	result := query.Find(&courses)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if count == 0 {
		return nil, 0, nil
	}

	return courses, count, nil
}