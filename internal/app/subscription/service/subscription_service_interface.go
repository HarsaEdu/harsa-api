package service

import (
	SubsPlanRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/subs_plan/repository"
	UserRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/midtrans"
	"github.com/go-playground/validator"
	"github.com/midtrans/midtrans-go/coreapi"
)

type SubscriptionService interface {
	Subscribe(request *web.CreateSubscribeRequest, userId uint) (*coreapi.ChargeResponse, error)
}

type SubscriptionServiceImpl struct {
	SubsPlanRepository SubsPlanRepositoryPkg.SubsPlanRepository
	UserRepository     UserRepositoryPkg.UserRepository
	MidtransCoreApi    midtrans.MidtransCoreApi
	Validate           *validator.Validate
}

func NewSubscriptionService(subsPlanRepository SubsPlanRepositoryPkg.SubsPlanRepository, userRepository UserRepositoryPkg.UserRepository, midtransCoreApi midtrans.MidtransCoreApi, validate *validator.Validate) SubscriptionService {
	return &SubscriptionServiceImpl{
		SubsPlanRepository: subsPlanRepository,
		UserRepository:     userRepository,
		MidtransCoreApi:    midtransCoreApi,
		Validate:           validate,
	}
}