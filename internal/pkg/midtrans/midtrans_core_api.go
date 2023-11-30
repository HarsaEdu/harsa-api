package midtrans

import (
	"github.com/midtrans/midtrans-go/coreapi"
)

func (coreApi *MidtransCoreApiImpl) ChargeTransaction(request *coreapi.ChargeReq) (*coreapi.ChargeResponse, error) {
	response, err := coreApi.Client.ChargeTransaction(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (coreApi *MidtransCoreApiImpl) CheckTransactionStatus(orderId string) (string, error) {
	response, err := coreApi.Client.CheckTransaction(orderId)
	if err != nil {
		return "", err
	} else {
		if response != nil {
			if response.TransactionStatus == "capture" {
				if response.FraudStatus == "challenge" {
					return "challenge", nil
				} else if response.FraudStatus == "accept" {
					return "success", nil
				}
			} else if response.TransactionStatus == "settlement" {
				return "success", nil
			} else if response.TransactionStatus == "deny" {
				return "deny", nil
			} else if response.TransactionStatus == "cancel" || response.TransactionStatus == "expire" {
				return "failure", nil
			} else if response.TransactionStatus == "pending" {
				return "pending", nil
			}
		}
	}

	return "", nil
}