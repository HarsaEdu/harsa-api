package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (service *InterestServiceImpl) CreateInterest(profileID uint, request *web.InterestRequest) error {
	err := service.Validator.Struct(request)
	if err != nil {
		return err
	}

	_, err = service.InterestRepository.FindByProfileID(profileID)
	if err == nil {
		return fmt.Errorf("interest exists")
	}

	for _, value := range request.CategoryID {
		interest := conversion.InterestRequestToModel(profileID, value)
		err = service.InterestRepository.CreateInterest(interest)
		if err != nil {
			return fmt.Errorf("failed to create interest, something happen")
		}
	}
	return nil
}
