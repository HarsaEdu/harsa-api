package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (submissionService *SubmissionServiceImpl) GetAll(moduleId int) ([]web.SubmissionsResponseWeb, error) {
	data, total, err := submissionService.SubmissionRepository.GetAllWeb(moduleId)
	if err != nil {
		return nil, fmt.Errorf("internal Server Error")
	}

	if total == 0 {
		return nil, fmt.Errorf("submission not found")
	}

	return data, nil
}
