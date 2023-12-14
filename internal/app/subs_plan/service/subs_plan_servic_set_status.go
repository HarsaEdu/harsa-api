package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (subsPlanService *SubsPlanServiceImpl) SetStatus(request *web.SubsPlanUpdateStatus, id uint) error {
	err := subsPlanService.Validator.Struct(request)
	if err != nil {
		return err
	}

	ifExist, _ := subsPlanService.SubsPlanRepository.FindById(int(id))
	if ifExist == nil {
		return fmt.Errorf("subs plan not found")
	}

	err = subsPlanService.SubsPlanRepository.UpdateStatus(request.IsActive, id)
	if err != nil {
		return err
	}
	return nil
}