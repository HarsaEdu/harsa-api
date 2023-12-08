package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (submissionAnswerService *SubmissionAnswerServiceImpl) UpdateWeb(ctx echo.Context, request *web.SubmissionAnswerUpdateWeb, id int) error {
	result := conversion.ConvertSubmissionAnswerUpdateToSubmissionAnswerWebDomain(request)

	ifExist, _ := submissionAnswerService.Repository.FindById(id)
	if ifExist == nil {
		return fmt.Errorf("submission answer not found")
	}

	err := submissionAnswerService.Repository.UpdateWeb(*result, id)
	if err != nil {
		return fmt.Errorf("error when updating : %s", err.Error())
	}

	err = submissionAnswerService.Validator.Struct(result)
	if err != nil {
		return err
	}

	return nil
}
