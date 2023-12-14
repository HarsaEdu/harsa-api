package midtrans

import (
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/pkg/parse"
	"github.com/midtrans/midtrans-go/coreapi"
)

func (coreApi *MidtransCoreApiImpl) ChargeTransaction(request *coreapi.ChargeReq) (*coreapi.ChargeResponse, error) {
	response, err := coreApi.Client.ChargeTransaction(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (coreApi *MidtransCoreApiImpl) CheckTransactionStatus(orderId string) (string, *domain.PaymentTransactionStatus, error) {
	response, midtransErr := coreApi.Client.CheckTransaction(orderId)

	var parseSettlementTime *time.Time
	var err error

	if response.SettlementTime != "" {
		parseSettlementTime, err = parse.ParseToJakartaZone(response.SettlementTime)
		if err != nil {
			return "", nil, err
		}
	}
	
	if midtransErr != nil {
		return "", nil, midtransErr
	} else {
		if response != nil {
			transactionStatus := domain.PaymentTransactionStatus{
				OrderID: response.OrderID,
				TransactionStatus: response.TransactionStatus,
				FraudStatus: response.FraudStatus,
			}

			if transactionStatus.TransactionStatus == "capture" {
				if transactionStatus.FraudStatus == "challenge" {
					return "challenge", &transactionStatus, nil
				} else if transactionStatus.FraudStatus == "accept" {
					transactionStatus.SettlementTime = *parseSettlementTime
					return "success", &transactionStatus, nil
				}
			} else if transactionStatus.TransactionStatus == "settlement" {
				transactionStatus.SettlementTime = *parseSettlementTime
				return "success", &transactionStatus, nil
			} else if transactionStatus.TransactionStatus == "deny" {
				return "deny", &transactionStatus, nil
			} else if transactionStatus.TransactionStatus == "cancel" || transactionStatus.TransactionStatus == "expire" {
				return "failure", &transactionStatus, nil
			} else if transactionStatus.TransactionStatus == "pending" {
				return "pending", &transactionStatus, nil
			}
		}
	}

	return "", nil, nil
}