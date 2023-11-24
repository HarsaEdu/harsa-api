package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (repository *InterestRepositoryImpl) FindByProfileID(profileID uint) (*domain.UserInterest, error) {
	interest := domain.UserInterest{}
	result := repository.DB.Where("profile_id=?", profileID).First(&interest)
	if result.Error != nil {
		return nil, result.Error
	}
	return &interest, nil
}
