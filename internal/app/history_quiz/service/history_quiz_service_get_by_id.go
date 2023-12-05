package service

import (
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)


func (historyQuizService *HistoryQuizServiceImpl) GetById(id uint) (*web.HistoryQuizResponseWeb, error) {
	result,err :=historyQuizService.HistoryQuizRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	response := conversion.ConvertHistoryQuizResponseWeb(result)

	return response,  nil
}