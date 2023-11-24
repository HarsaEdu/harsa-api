package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (feedbackService *FeedbackServiceImpl) Create(request web.FeedbackCreateRequest) error {
	err := feedbackService.Validator.Struct(request)
	if err != nil {
		return err
	}

	feedback := conversionRequest.FeedbackCreateRequestToCategoriesModel(request)

	err = feedbackService.FeedbackRepository.Create(feedback)
	if err != nil {
		return fmt.Errorf("error when creating feedback %s", err.Error())
	}

	return nil

}
