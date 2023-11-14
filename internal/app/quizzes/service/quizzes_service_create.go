package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (quizzesService *QuizzesServiceImpl) Create(ctx echo.Context, request web.QuizRequest) error {
	err := quizzesService.Validator.Struct(request)
	if err != nil {
		return err
	}

	quiz := conversionRequest.QuizCreateRequestToQuizzesModel(request)

	err = quizzesService.QuizzesRepository.Create(quiz)
	if err != nil {
		return fmt.Errorf("error when creating quiz %s", err.Error())
	}

	return nil

}

