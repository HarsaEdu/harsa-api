package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (repository *InterestRepositoryImpl) GetInterestRecommendation(profileID uint) ([]domain.Course, int64, error) {
	course := []domain.Course{}
	result := repository.DB.Joins("JOIN categories ON courses.category_id = categories.id").
		Joins("JOIN user_interests ON user_interests.category_id = categories.id").
		Where("user_interests.profile_id = ?", profileID).
		Find(&course)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var total int64
	result.Count(&total)
	return course, total, nil
}
