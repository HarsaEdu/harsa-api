package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (submissionAnswerService *SubmissionAnswerServiceImpl) FindById(id int) (*web.SubmissionAnswerResponseMobile, error) {

	result, err := submissionAnswerService.Repository.FindById(id)

	if result == nil {
		return nil, fmt.Errorf(err.Error(), "submission answer not found")
	}

	if err != nil {
		return nil, fmt.Errorf("internal Server Error")
	}

	response := conversion.ConvertSubmissionAnswerResponseMobile(result)

	return response, nil
}

func (submissionAnswerService *SubmissionAnswerServiceImpl) FindByIdWeb(id int) (*web.SubmissionAnswerResponseWebById, error) {

	result, err := submissionAnswerService.Repository.FindByIdWeb(id)

	if result == nil {
		return nil, fmt.Errorf(err.Error(), "submission answer not found")
	}

	if err != nil {
		return nil, fmt.Errorf("internal Server Error")
	}

	return result, nil
}
