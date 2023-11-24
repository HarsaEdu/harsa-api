package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (interestRepository *InterestRepositoryImpl) CreateInterest(interest *domain.UserInterest) error {
	result := interestRepository.DB.Create(&interest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
