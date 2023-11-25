package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (courseRepository *CourseRepositoryImpl) GetAll(offset, limit int, search string, category string) ([]domain.CourseEntity, int64, error) {
	var courses []domain.CourseEntity
	var count int64

	query := courseRepository.DB.Model(&domain.Course{}).Select("courses.id as id, title, courses.description as description, enrolled, rating, courses.image_url as image_url,  courses.created_at as created_at, courses.updated_at as updated_at, roles.name as role_name, user_profiles.user_id as user_id, user_profiles.first_name as first_name, user_profiles.last_name as last_name, user_profiles.job as job, categories.id as category_id, categories.name as category_name").
		Joins("left join user_profiles on user_profiles.user_id = courses.user_id").
		Joins("left join users on users.id = user_profiles.user_id").
		Joins("left join roles on roles.id = users.role_id"). 
		Joins("left join categories on categories.id = courses.category_id")

	query.Where("categories.name LIKE ?", "%"+category+"%" )

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

	if offset >= int(count) {
		return nil, 0, nil
	}

	return courses, count, nil
}