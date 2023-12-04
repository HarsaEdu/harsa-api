package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (submissionAnswerService *SubmissionAnswerServiceImpl) Get(offset, limit, submissionID int, search string) (*web.SubmissionAnswerResponseWeb, *web.Pagination, error) {
	answer, total, err := submissionAnswerService.Repository.Get(offset, limit, search, uint(submissionID))
	if total == 0 {
		return nil, nil, fmt.Errorf("not found")
	}
	if err != nil {
		return nil, nil, err
	}

	submission, err := submissionAnswerService.SubmissionRepo.FindById(submissionID)
	if err != nil {
		return nil, nil,  err
	}

	submissionRes := conversion.ConvertSubmissionResponseModule(submission)
	answerRes := conversion.ConverstSubmissionAnswerToResponseWeb(answer,submissionRes)
	pagination := conversion.RecordToPaginationResponse(offset, limit, total)

	return answerRes, pagination, nil

}
