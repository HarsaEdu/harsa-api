package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/payment/repository"
	SubsPlanRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/subs_plan/repository"
	UserRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/midtrans"
	"github.com/go-playground/validator"
)

type PaymentService interface {
	CreatePaymentSubscription(request *web.CreatePaymentSubscriptionRequest, userId uint) (*web.GetPaymentResponse, error)
	NotificationPayment(notificationPayload map[string]interface{}) error
	GetPaymentHistoryById(orderId string) (*web.GetPaymentResponse, error)
}

type PaymentServiceImpl struct {
	PaymentRepository  repository.PaymentRepository
	SubsPlanRepository SubsPlanRepositoryPkg.SubsPlanRepository
	UserRepository     UserRepositoryPkg.UserRepository
	MidtransCoreApi    midtrans.MidtransCoreApi
	Validate           *validator.Validate
}

func NewPaymentService(paymentRepository repository.PaymentRepository, subsPlanRepository SubsPlanRepositoryPkg.SubsPlanRepository, userRepository UserRepositoryPkg.UserRepository, midtransCoreApi midtrans.MidtransCoreApi, validate *validator.Validate) PaymentService {
	return &PaymentServiceImpl{
		PaymentRepository: paymentRepository,
		SubsPlanRepository: subsPlanRepository,
		UserRepository:     userRepository,
		MidtransCoreApi:    midtransCoreApi,
		Validate:           validate,
	}
}