package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (submissionService *SubmissionServiceImpl) Update(ctx echo.Context, request *web.SubmissionUpdateRequest, submissionId int) error {

	err := submissionService.Validator.Struct(request)
	if err != nil {
		return err
	}

	ifExist, _ := submissionService.SubmissionRepository.FindById(submissionId)
	if ifExist == nil {
		return fmt.Errorf("submission not found")
	}

	result := conversion.SubmissionUpdateRequestToSubmissionDomain(request)

	err = submissionService.SubmissionRepository.Update(result, submissionId)

	if err != nil {
		return fmt.Errorf("error when updating submission %s", err.Error())
	}

	return nil
}
