package conversion

import (
	"fmt"
	"time"

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

func ChargeResponseToPaymentHistoryDomain(response *coreapi.ChargeResponse, customer *domain.UserDetail, itemId uint) *domain.PaymentHistory {
	parseTransactionTime, _ := time.Parse("2006-01-02 15:04:05", response.TransactionTime)
	parseExpireTime, _ := time.Parse("2006-01-02 15:04:05", response.ExpiryTime)
	return &domain.PaymentHistory{
		ID: response.OrderID,
		UserId: customer.UserID,
		ItemId: itemId,
		Status: response.TransactionStatus,
		Method: response.PaymentType,
		GrossAmount: response.GrossAmount,
		BankName: response.VaNumbers[0].Bank,
		VaNumber: response.VaNumbers[0].VANumber,
		CreatedAt: parseTransactionTime,
		UpdatedAt: parseTransactionTime,
		ExpiredAt: parseExpireTime,
	}
}

func PaymentHistoryDomainToPaymentHistoryResponse(paymentHistory *domain.PaymentHistory) *web.GetPaymentResponse {
	return &web.GetPaymentResponse{
		ID: paymentHistory.ID,
		UserId: paymentHistory.UserId,
		ItemId: paymentHistory.ItemId,
		Status: paymentHistory.Status,
		Method: paymentHistory.Method,
		GrossAmount: paymentHistory.GrossAmount,
		BankName: paymentHistory.BankName,
		VaNumber: paymentHistory.VaNumber,
		CreatedAt: paymentHistory.CreatedAt,
		UpdatedAt: paymentHistory.UpdatedAt,
		ExpiredAt: paymentHistory.ExpiredAt,
		Customer: web.PaymentCustomerResponse{
			ID: paymentHistory.User.ID,
			Name: paymentHistory.User.UserProfile.FirstName + " " + paymentHistory.User.UserProfile.LastName,
		},
		Item: web.PaymentItemResponse{
			ID: paymentHistory.Item.ID,
			Name: paymentHistory.Item.Title,
		},
	}
}