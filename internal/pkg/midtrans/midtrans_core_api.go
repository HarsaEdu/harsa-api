package midtrans

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func (coreApi *MidtransCoreApiImpl) CreateTransaction(subscription *domain.SubsPlan, customer *domain.UserDetail) (*coreapi.ChargeResponse, error) {
	orderId := fmt.Sprintf("%s", uuid.New().String())

	chargeRequest := coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: orderId,
			GrossAmt: int64(subscription.Price),
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
		},
		Items: &[]midtrans.ItemDetails{
			midtrans.ItemDetails{
				Name: subscription.Title,
				ID: fmt.Sprintf("SUBSPLAN-%d", subscription.ID),
				Qty: 1,
				Price: int64(subscription.Price),
				Brand: "HarsaEdu",
				Category: "Subscription Plan",
				MerchantName: "HarsaEdu",
			},
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: customer.FirstName,
			LName: customer.LastName,
			Email: customer.Email,
			Phone: customer.PhoneNumber,
		},
	}

	chargeResponse, err := coreApi.chargeTransaction(&chargeRequest)
	if err != nil {
		return nil, err
	}

	return chargeResponse, nil
}

func (coreApi *MidtransCoreApiImpl) chargeTransaction(request *coreapi.ChargeReq) (*coreapi.ChargeResponse, error) {
	response, err := coreApi.Client.ChargeTransaction(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}