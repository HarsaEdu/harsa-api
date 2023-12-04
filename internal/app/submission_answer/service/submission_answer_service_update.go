package service

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (submissionAnswerService *SubmissionAnswerServiceImpl) Update(ctx echo.Context, request *web.SubmissionAnswerUpdateRequest, id, idUser int) error {
	result := conversion.ConvertSubmissionAnswerUpdateToSubmissionAnswerDomain(request)

	ifExist, _ := submissionAnswerService.Repository.FindById(id)
	if ifExist == nil {
		return fmt.Errorf("submission answer not found")
	}

	fileUrl, err := submissionAnswerService.CloudinaryUploader.Uploader(ctx, "file", "submission_answer", true)
	if !strings.HasSuffix(fileUrl, ".pdf") {
		return fmt.Errorf("invalid file format")
	}

	result.SubmittedUrl = fileUrl

	if err != nil {
		return fmt.Errorf("error when updating submission answer : %s", err.Error())
	}

	err = submissionAnswerService.Repository.Update(*result, id, idUser)

	if err != nil {
		return fmt.Errorf("error when updating submission answer %s", err.Error())
	}

	err = submissionAnswerService.Validator.Struct(result)
	if err != nil {
		return err
	}

	return nil

}
