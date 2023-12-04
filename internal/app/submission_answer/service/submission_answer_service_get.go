package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (submissionAnswerService *SubmissionAnswerServiceImpl) Get(offset, limit, moduleId int, search string) ([]*web.SubmissionAnswerResponseWeb, *web.SubmissionsResponseModule, *web.Pagination, int64, error) {
	answer, total, err := submissionAnswerService.Repository.Get(offset, limit, search)
	if total == 0 {
		return nil, nil, nil, 0, fmt.Errorf("not found")
	}
	if err != nil {
		return nil, nil, nil, 0, err
	}

	submission, err := submissionAnswerService.SubmissionRepo.FindById(moduleId)
	if err != nil {
		return nil, nil, nil, 0, err
	}

	answerRes := conversion.ConverstSubmissionAnswerToResponseWeb(answer)
	submissionRes := conversion.ConvertSubmissionResponseModule(submission)
	pagination := conversion.RecordToPaginationResponse(offset, limit, total)

	return answerRes, submissionRes, pagination, total, nil

}
