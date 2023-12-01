package service

import "fmt"

func (submissionService *SubmissionServiceImpl) Delete(id int) error {

	ifExist, _ := submissionService.SubmissionRepository.FindById(id)
	if ifExist == nil {
		return fmt.Errorf("submission not found")
	}

	err := submissionService.SubmissionRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting %s", err.Error())
	}

	return nil
}
