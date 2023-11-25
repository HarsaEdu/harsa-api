package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (feedbackService *FeedbackServiceImpl) Update(request web.FeedbackUpdateRequest, id int) error {
	// Check if the request is valid
	err := feedbackService.Validator.Struct(request)
	if err != nil {
		return err
	}

	IfExist, _ := feedbackService.FeedbackRepository.GetById(id)
	if IfExist == nil {
		return fmt.Errorf("feedback not found")
	}

	feedback := conversionRequest.FeedbackUpdateRequestToCategoriesModel(request)

	err = feedbackService.FeedbackRepository.Update(id, feedback)
	if err != nil {
		return fmt.Errorf("error when updating : %s", err.Error())
	}

	if IfExist.Rating != feedback.Rating {
		err = feedbackService.FeedbackRepository.AutoUpdateRating(IfExist.CourseID)
		if err != nil {
			return fmt.Errorf("error when update rating %s", err.Error())
		}
	}

	return nil
}
