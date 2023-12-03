package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/payment/repository"
	SubsPlanRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/subs_plan/repository"
	subscriptionServicePkg "github.com/HarsaEdu/harsa-api/internal/app/subscription/service"
	UserRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/midtrans"
	"github.com/go-playground/validator"
)

type PaymentService interface {
	CreatePaymentSubscription(request *web.CreatePaymentSubscriptionRequest, userId uint) (*web.GetPaymentResponse, error)
	NotificationPayment(notificationPayload map[string]interface{}) error
	GetPaymentHistoryById(orderId string) (*web.GetPaymentResponse, error)
	GetAllPaymentHistory(offset, limit int, search string, status string) ([]web.GetPaymentResponse, *web.Pagination, error)
	GetAllPaymentHistoryByUserId(userId uint, offset, limit int, search string, status string) ([]web.GetPaymentResponse, *web.Pagination, error)
	GetPaymentHistoryByUserIdAndPaymentId(userId uint, paymentId string) (*web.GetPaymentResponse, error)
}

type PaymentServiceImpl struct {
	PaymentRepository  repository.PaymentRepository
	SubsPlanRepository SubsPlanRepositoryPkg.SubsPlanRepository
	UserRepository     UserRepositoryPkg.UserRepository
	SubscriptionService subscriptionServicePkg.SubscriptionService
	MidtransCoreApi    midtrans.MidtransCoreApi
	Validate           *validator.Validate
}

func NewPaymentService(paymentRepository repository.PaymentRepository, subsPlanRepository SubsPlanRepositoryPkg.SubsPlanRepository, userRepository UserRepositoryPkg.UserRepository, subscriptionService subscriptionServicePkg.SubscriptionService, midtransCoreApi midtrans.MidtransCoreApi, validate *validator.Validate) PaymentService {
	return &PaymentServiceImpl{
		PaymentRepository: paymentRepository,
		SubsPlanRepository: subsPlanRepository,
		UserRepository:     userRepository,
		SubscriptionService: subscriptionService,
		MidtransCoreApi:    midtransCoreApi,
		Validate:           validate,
	}
}