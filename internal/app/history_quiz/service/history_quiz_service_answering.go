package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (historyQuizService *HistoryQuizServiceImpl) Create(request []web.AnswersRequest, userID uint, quizID uint) error {
	
	exist, err := historyQuizService.HistoryQuizRepository.Cek(userID, quizID)
	if exist != nil{
		return fmt.Errorf("You have completed the quiz")
	}
	for _, ans := range request {
		if err := historyQuizService.Validator.Struct(ans); err != nil {
		   return err
	    }
	}

	err = historyQuizService.HistoryQuizRepository.Answering(userID, request, quizID)
	if err != nil {
		return fmt.Errorf("error when creating history quiz %s", err.Error())
	}

	return nil

}