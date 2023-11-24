package service

import (
	"fmt"
)

func (questionsService *QuestionsServiceImpl) Delete(userId uint, questionId uint, role string) error {

	questionExist, err := questionsService.QuestionsRepository.CekIdFromQuestion(userId, questionId, role)
	if err != nil { 
		return fmt.Errorf("error when cek id user in question delete :%s", err.Error())
	}

	err = questionsService.QuestionsRepository.Delete(questionExist)
	if err != nil {
		return fmt.Errorf("error when delete question : %s", err.Error())
	} 

	return nil

}