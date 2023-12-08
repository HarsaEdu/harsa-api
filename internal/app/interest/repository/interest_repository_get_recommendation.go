package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (repository *InterestRepositoryImpl) GetInterestRecommendation(profileID uint) ([]domain.CourseEntity, int64, error) {
	course := []domain.CourseEntity{}
	result := repository.DB.Model(&domain.Course{}).Select("courses.id as id, title, enrolled, rating, courses.image_url as image_url, user_profiles.user_id as user_id, user_profiles.first_name as first_name, user_profiles.last_name as last_name, user_profiles.job as job, categories.id as category_id, categories.name as category_name").
		Joins("JOIN categories ON courses.category_id = categories.id").
		Joins("JOIN user_interests ON user_interests.category_id = categories.id").
		Joins("JOIN user_profiles ON user_profiles.user_id = courses.user_id").
		Where("user_interests.profile_id = ?", profileID).
		Find(&course)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	var total int64
	result.Count(&total)
	return course, total, nil
}
