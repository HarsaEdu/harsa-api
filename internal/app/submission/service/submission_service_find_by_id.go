package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (submissionService *SubmissionServiceImpl) FindById(id int) (*web.SubmissionsResponseModule, error) {

	data, _ := submissionService.SubmissionRepository.FindById(id)
	if data == nil {
		return nil, fmt.Errorf("submission not found")
	}

	result := conversion.ConvertSubmissionResponseModule(data)
	return result, nil
}
