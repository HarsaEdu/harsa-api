package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (feedbackRepository *FeedbackRepositoryImpl) Create(feedback *domain.Feedback) error {
	result := feedbackRepository.DB.Create(&feedback)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
