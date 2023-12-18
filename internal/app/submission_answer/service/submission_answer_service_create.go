package service

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (submissionAnswerService *SubmissionAnswerServiceImpl) Create(ctx echo.Context, request *web.SubmissionAnswerRequest, idSubmission, idUser int) error {
	
	ifExist, _ := submissionAnswerService.Repository.FindByuserIDAndSubmissionID(idSubmission,idUser)
	if ifExist != nil {
		return fmt.Errorf("already exist")
	}

	
	result := conversion.ConvertSubmissionAnswerRequestToSubmissionAnswerDomain(request)

	result.SubmissionID = uint(idSubmission)
	result.UserID = uint(idUser)
	result.Status = domain.StatusSubmissionAnswerSubmitted

	fileUrl, err := submissionAnswerService.CloudinaryUploader.Uploader(ctx, "file", "submission_answer", true)
	if !strings.HasSuffix(fileUrl, ".pdf") {
		return fmt.Errorf("invalid file format")
	}

	result.SubmittedUrl = fileUrl

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
