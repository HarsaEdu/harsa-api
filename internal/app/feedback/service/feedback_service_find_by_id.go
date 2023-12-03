package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (feedbackService *FeedbackServiceImpl) FindById(id int) (*web.FeedBackResponseForTracking, error) {
	result, _ := feedbackService.FeedbackRepository.GetById(id)

	if result == nil {
		return nil, fmt.Errorf("feedback not found")
	}
	
	response := conversion.ConvertFeedbackForTracking(result)

	return response, nil
}

func (feedbackService *FeedbackServiceImpl) GetByIdUserAndCourseId(userId, courseId uint) (*web.FeedBackResponseForTracking, error) {
	result, _ := feedbackService.FeedbackRepository.GetByIdUserAndCourseId(userId, courseId)

	if result == nil {
		return nil, fmt.Errorf("feedback not found")
	}

	response := conversion.ConvertFeedbackForTracking(result)

	return response, nil
}