package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (quizzesService *QuizzesServiceImpl) GetAll(moduleId uint, offset int, limit int, search string) ([]web.GetAllQuizResponse, *web.Pagination, error) {
	var quizzes []web.GetAllQuizResponse
	result, total, err := quizzesService.QuizzesRepository.GetAll(moduleId, offset, limit, search)
	if err != nil {
		return nil, nil, err
	}

	if result == nil {
		return nil, nil, fmt.Errorf("quiz not found")
	}

	quizzes = conversion.ConvertAllQuiz(result)

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)
	
	return quizzes , pagination,nil
}