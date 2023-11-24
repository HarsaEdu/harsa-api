package service

import "fmt"

func (subsPlanService *SubsPlanServiceImpl) Delete(id int) error {
	IfExist, _ := subsPlanService.SubsPlanRepository.FindById(id)
	if IfExist == nil {
		return fmt.Errorf("subs plan not found")
	}
	err := subsPlanService.SubsPlanRepository.Delete(id)

	if err != nil {
		return fmt.Errorf("error when deleting : %s", err)
	}

	return nil
}
