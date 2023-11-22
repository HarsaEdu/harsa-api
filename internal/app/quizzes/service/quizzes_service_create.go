package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (quizzesService *QuizzesServiceImpl) Create(request web.QuizRequest, role string) error {
	err := quizzesService.Validator.Struct(request)
	if err != nil {
		return err
	}

	quiz := conversionRequest.QuizCreateRequestToQuizzesModel(request)

	err = quizzesService.QuizzesRepository.CekIdFromModule(request.UserID, request.ModuleID, role)
	if err != nil { 
		return fmt.Errorf("error when cek id user in quiz update :%s", err.Error())
	}

	err = quizzesService.QuizzesRepository.Create(quiz)
	if err != nil {
		return fmt.Errorf("error when creating quiz %s", err.Error())
	}

	return nil

}

