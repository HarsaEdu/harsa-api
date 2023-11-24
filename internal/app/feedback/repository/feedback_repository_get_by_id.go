package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (feedbackRepository *FeedbackRepositoryImpl) GetById(id int) (*domain.Feedback, error) {
	feedback := &domain.Feedback{}

	result := feedbackRepository.DB.First(&feedback, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return feedback, nil
}
