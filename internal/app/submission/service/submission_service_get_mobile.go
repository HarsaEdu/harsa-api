package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (submissionService *SubmissionServiceImpl) GetAllMobile(moduleId int) ([]web.SubmissionsResponseModuleMobile, error) {
	data, total, err := submissionService.SubmissionRepository.GetAll(moduleId)
	if err != nil {
		return nil, fmt.Errorf("internal Server Error")
	}

	if total == 0 {
		return nil, fmt.Errorf("submission not found")
	}

	result := conversion.ConvertAllSubmissionModuleMobile(data)
	return result, nil
}
