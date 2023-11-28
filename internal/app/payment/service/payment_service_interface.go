package service

import (
	SubsPlanRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/subs_plan/repository"
	UserRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/midtrans"
	"github.com/go-playground/validator"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PaymentService interface {
	CreatePaymentSubscription(request *web.CreatePaymentSubscriptionRequest, userId uint) (*coreapi.ChargeResponse, error)
}

type PaymentServiceImpl struct {
	SubsPlanRepository SubsPlanRepositoryPkg.SubsPlanRepository
	UserRepository     UserRepositoryPkg.UserRepository
	MidtransCoreApi    midtrans.MidtransCoreApi
	Validate           *validator.Validate
}

func NewPaymentService(subsPlanRepository SubsPlanRepositoryPkg.SubsPlanRepository, userRepository UserRepositoryPkg.UserRepository, midtransCoreApi midtrans.MidtransCoreApi, validate *validator.Validate) PaymentService {
	return &PaymentServiceImpl{
		SubsPlanRepository: subsPlanRepository,
		UserRepository:     userRepository,
		MidtransCoreApi:    midtransCoreApi,
		Validate:           validate,
	}
}