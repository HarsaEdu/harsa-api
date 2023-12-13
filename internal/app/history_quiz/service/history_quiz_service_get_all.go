package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)


func (historyQuizService *HistoryQuizServiceImpl) GetAllByQuizWeb(quizId uint, offset, limit int, search string) ([]web.HistoryQuizResponseWeb, *web.Pagination, error) {
	result, total, err :=historyQuizService.HistoryQuizRepository.GetAllHistoryQuizByQuizId(offset, limit, quizId, search)

	if total == 0 {
		return nil, nil, fmt.Errorf("history quiz not found")
	}

	if err != nil {
		return nil, nil, fmt.Errorf("internal Server Error")
	}

	response := conversion.ConvertAllHistoryQuizResponseWeb(result)

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)

	return response, pagination, nil
}