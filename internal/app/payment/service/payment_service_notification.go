package service

import "fmt"

func (paymentService *PaymentServiceImpl) NotificationPayment(notificationPayload map[string]interface{}) error {
	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		return fmt.Errorf("error when get order id : order id not found")
	}

	transactionStatus, transactionResult, err := paymentService.MidtransCoreApi.CheckTransactionStatus(orderId)
	if err != nil {
		return fmt.Errorf("error when checking transaction status : %s", err.Error())
	}

	err = paymentService.PaymentRepository.UpdateStatusPaymentHistory(transactionStatus, transactionResult)

	transaction, err := paymentService.PaymentRepository.GetPaymentHistoryById(transactionResult.OrderID)
	if err != nil {
		return fmt.Errorf("error when get payment history : %s", err.Error())
	}

	if transaction == nil {
		return fmt.Errorf("error when get payment history : payment history not found")
	}

	if transaction.Status == "success" {
		err = paymentService.SubscriptionService.SubscriptionAdd(transaction.UserId, uint(transaction.Item.Duration_days))
	}

	return nil
}