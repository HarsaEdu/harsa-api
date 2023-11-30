package conversion

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func CreatePaymentSubscriptionRequestToMidtransChargeRequest(subscription *domain.SubsPlan, customer *domain.UserDetail, request *web.CreatePaymentSubscriptionRequest) *coreapi.ChargeReq {
	orderId := fmt.Sprintf("%s", uuid.New().String())

	item := midtrans.ItemDetails{
		ID: fmt.Sprintf("SUBSPLAN-%d", subscription.ID),
		Name: subscription.Title,
		Price: int64(subscription.Price),
		Qty: 1,
		Brand: "HarsaEdu",
		MerchantName: "HarsaEdu",
		Category: "Subscription Plan",
	}


	chargeRequest := coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: orderId,
			GrossAmt: int64(subscription.Price),
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: customer.FirstName,
			LName: customer.LastName,
			Email: customer.Email,
			Phone: customer.PhoneNumber,
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.Bank(request.BankName),
		},
		Items: &[]midtrans.ItemDetails{item},
	}

	return &chargeRequest
}