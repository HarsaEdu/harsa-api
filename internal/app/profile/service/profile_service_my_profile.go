package service

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (service *ProfileServiceImpl) MyProfile(request *web.UserGetByIDRequest) (*domain.ProfileDetailMobile, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return nil, err
	}

	result, err := service.ProfileRepository.FindByUserID(request.ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
