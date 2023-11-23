package service

import (
	"fmt"
)

func (optionsService *OptionsServiceImpl) Delete(userId uint, optionId uint, role string) error {

	questionExist, err := optionsService.OptionsRepository.CekIdFromOption(userId, optionId, role)
	if err != nil { 
		return fmt.Errorf("error when cek id user in optiondelete :%s", err.Error())
	}

	err = optionsService.OptionsRepository.Delete(questionExist)
	if err != nil {
		return fmt.Errorf("error when delete option : %s", err.Error())
	} 

	return nil

}