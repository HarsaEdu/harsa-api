package service

import (
	"fmt"
)

func (quizzesService *QuizzesServiceImpl) Delete(userId uint, quizId uint, role string) error {

	quizExist, err := quizzesService.QuizzesRepository.CekIdFromQuiz(userId, quizId, role)
	if err != nil { 
		return fmt.Errorf("error when cek id user in quiz delete :%s", err.Error())
	}

	err = quizzesService.QuizzesRepository.Delete(quizExist)
	if err != nil {
		return fmt.Errorf("error when delete quiz : %s", err.Error())
	} 

	return nil

}