package midtrans

import (
	"github.com/HarsaEdu/harsa-api/configs"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type MidtransCoreApi interface {
	ChargeTransaction(request *coreapi.ChargeReq) (*coreapi.ChargeResponse, error)
	CheckTransactionStatus(orderId string) (string, *domain.PaymentTransactionStatus, error)
}

type MidtransCoreApiImpl struct {
	Client coreapi.Client
}

func NewMidtransCoreApi(config *configs.MidtransConfig) MidtransCoreApi {
	client := coreapi.Client{}
	client.New(config.ServerKey, midtrans.Sandbox)

	return &MidtransCoreApiImpl{
		Client: client,
	}
}