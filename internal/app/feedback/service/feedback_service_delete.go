package service

import (
	"fmt"
)

func (feedbackService *FeedbackServiceImpl) DeleteByUserAndCourseId(userId, courseId uint) error {

	IfExist, _ := feedbackService.FeedbackRepository.GetByIdUserAndCourseId(userId, courseId)
	if IfExist == nil {
		return fmt.Errorf("feedback not found")
	}
	err := feedbackService.FeedbackRepository.DeleteByUserAndCourseId(userId, courseId)
	if err != nil {
		return fmt.Errorf("error when deleting : %s", err)
	}

	err = feedbackService.FeedbackRepository.AutoUpdateRating(IfExist.CourseID)
	if err != nil {
		return fmt.Errorf("error when update rating %s", err.Error())
	}

	return nil
}

func (feedbackService *FeedbackServiceImpl) DeleteById(id uint) error {

	IfExist, _ := feedbackService.FeedbackRepository.GetById(id)
	if IfExist == nil {
		return fmt.Errorf("feedback not found")
	}
	err := feedbackService.FeedbackRepository.DeleteById(id)
	if err != nil {
		return fmt.Errorf("error when deleting : %s", err)
	}

	err = feedbackService.FeedbackRepository.AutoUpdateRating(IfExist.CourseID)
	if err != nil {
		return fmt.Errorf("error when update rating %s", err.Error())
	}

	return nil
}
