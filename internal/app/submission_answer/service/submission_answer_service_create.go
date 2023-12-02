package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (submissionAnswerService *SubmissionAnswerServiceImpl) Create(ctx echo.Context, request *web.SubmissionAnswerRequest, idSubmission, idUser int) error {

	result := conversion.ConvertSubmissionAnswerRequestToSubmissionAnswerDomain(request)

	result.SubmissionID = uint(idSubmission)
	result.UserID = uint(idUser)

	fileUrl, err := submissionAnswerService.CloudinaryUploader.Uploader(ctx, "file", "submission_answer", true)
	if fileUrl != "" {
		result.SubmittedUrl = fileUrl
	}
	if err != nil {
		return fmt.Errorf("error when uploading submission answer : %s", err.Error())
	}

	err = submissionAnswerService.Repository.Create(*result)

	if err != nil {
		return fmt.Errorf("error when upload submission answer %s", err.Error())
	}

	err = submissionAnswerService.Validator.Struct(result)
	if err != nil {
		return err
	}

	return nil
}
