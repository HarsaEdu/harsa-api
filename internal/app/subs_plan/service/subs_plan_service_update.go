package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (subsPlanService *SubsPlanServiceImpl) Update(subsPlan *web.SubsPlanUpdateRequest, id int) error {
	err := subsPlanService.Validator.Struct(subsPlan)
	if err != nil {
		return err
	}
	ifExist, _ := subsPlanService.SubsPlanRepository.FindById(id)
	if ifExist == nil {
		return fmt.Errorf("subs plan not found")
	}

	result := conversion.SubsPlanUpdateRequestToSubsPlanDomain(subsPlan)

	err = subsPlanService.SubsPlanRepository.Update(result, id)
	if err != nil {
		return fmt.Errorf("error when updating subs plan %s", err.Error())
	}

	return nil
}
