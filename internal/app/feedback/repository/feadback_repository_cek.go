package repository

import (

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (feedbackRepository *FeedbackRepositoryImpl) Cek(userId uint, courseID uint) (*domain.Feedback, error) {

	feedback := domain.Feedback{}

	if err := feedbackRepository.DB.Model(&domain.Feedback{}).Where("user_id  = ? AND course_id = ?" ,userId,courseID).First(&feedback).
		Error; err != nil {
		return nil,  err
	}
	
	return &feedback,  nil
}