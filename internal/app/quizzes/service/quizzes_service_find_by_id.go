package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (quizzesService *QuizzesServiceImpl) FindById(quizId uint) (*web.QuizResponse, error) {
	
	quiz, err := quizzesService.QuizzesRepository.FindById(quizId)
	if err != nil { 
		return nil, fmt.Errorf("error when find quiz by id :%s", err.Error())
	}

	convertRes := conversion.ConvertQuizRes(quiz)

	return convertRes, nil

}