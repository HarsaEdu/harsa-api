package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (repository *InterestRepositoryImpl) GetTopInterests(limit int) ([]domain.UserInterest, error) {
	topInterest := []domain.UserInterest{}

	result := repository.DB.Select("category_id, count(*) as interest_count").Group("category_id").Order("interest_count desc").Limit(limit).Preload("Category").Preload("Profile").Find(&topInterest)
	if result.Error != nil {
		return nil, result.Error
	}

	return topInterest, nil
}