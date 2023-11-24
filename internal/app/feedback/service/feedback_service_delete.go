package service

import (
	"fmt"
)

func (feedbackService *FeedbackServiceImpl) Delete(id int) error {

	IfExist, _ := feedbackService.FeedbackRepository.GetById(id)
	if IfExist == nil {
		return fmt.Errorf("feedback not found")
	}
	err := feedbackService.FeedbackRepository.Delete(id)

	if err != nil {
		return fmt.Errorf("error when deleting : %s", err)
	}

	return nil
}
