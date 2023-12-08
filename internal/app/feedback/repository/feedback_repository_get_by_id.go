package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (feedbackRepository *FeedbackRepositoryImpl) GetByIdAndCourseId(courseId, id uint) (*domain.Feedback, error) {
	feedback := &domain.Feedback{}

	result := feedbackRepository.DB.Preload("User.UserProfile").Where("course_id = ? AND id = ?", courseId, id).First(&feedback, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return feedback, nil
}

func (feedbackRepository *FeedbackRepositoryImpl) GetByIdUserAndCourseId(userId, courseId uint) (*domain.Feedback, error) {
	feedback := &domain.Feedback{}

	result := feedbackRepository.DB.Preload("User.UserProfile").Where("user_id = ? AND course_id = ?", userId, courseId).First(&feedback)
	if result.Error != nil {
		return nil, result.Error
	}

	return feedback, nil
}