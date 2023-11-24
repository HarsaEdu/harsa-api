package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (feedbackService *FeedbackServiceImpl) FindById(id int) (*domain.Feedback, error) {
	result, _ := feedbackService.FeedbackRepository.GetById(id)

	if result == nil {
		return nil, fmt.Errorf("feedback not found")
	}

	return result, nil
}
