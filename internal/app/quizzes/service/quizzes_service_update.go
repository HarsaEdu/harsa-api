package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (quizzesService *QuizzesServiceImpl) Update(request web.QuizRequest, quizId uint, role string) error {
	err := quizzesService.Validator.Struct(request)
	if err != nil {
		return err
	}

	quiz := conversionRequest.QuizCreateRequestToQuizzesModel(request)

	quizExist, err := quizzesService.QuizzesRepository.CekIdFromQuiz(request.UserId, quizId, role)
	if err != nil { 
		return fmt.Errorf("error when cek id user in quiz update :%s", err.Error())
	}

	err = quizzesService.QuizzesRepository.Update(quiz, quizExist)
	if err != nil {
		return fmt.Errorf("error when update quiz %s", err.Error())
	} 

	return nil

}