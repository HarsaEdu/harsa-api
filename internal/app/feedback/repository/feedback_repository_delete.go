package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (feedbackRepository *FeedbackRepositoryImpl) Delete(id int) error {
	result := feedbackRepository.DB.Where("id = ?", id).Delete(&domain.Feedback{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
