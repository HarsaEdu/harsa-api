package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (submissionService *SubmissionServiceImpl) Create(ctx echo.Context, request web.SubmissionRequest) error {

	err := submissionService.Validator.Struct(request)
	if err != nil {
		return err
	}

	result := conversion.SubmissionRequestToSubmissionDomain(&request)

	err = submissionService.SubmissionRepository.Create(result)
	if err != nil {
		return fmt.Errorf("error when creating submission %s", err.Error())
	}

	return nil
}
