package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (feedbackService *FeedbackServiceImpl) Create(request web.FeedbackCreateRequest, userId uint) error {
	err := feedbackService.Validator.Struct(request)
	if err != nil {
		return err
	}

	feedback := conversionRequest.FeedbackCreateRequestToCategoriesModel(request, userId)

	err = feedbackService.FeedbackRepository.Create(feedback)
	if err != nil {
		return fmt.Errorf("error when creating feedback %s", err.Error())
	}

	err = feedbackService.FeedbackRepository.AutoUpdateRating(feedback.CourseID)
	if err != nil {
		return fmt.Errorf("error when update rating %s", err.Error())
	}

	return nil

}
