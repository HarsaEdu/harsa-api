package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (feedbackRepository *FeedbackRepositoryImpl) Update(id int, feedback *domain.Feedback) error {
	result := feedbackRepository.DB.Where("id = ?", id).Updates(&domain.Feedback{Rating: feedback.Rating, Content: feedback.Content, CourseID: feedback.CourseID})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
