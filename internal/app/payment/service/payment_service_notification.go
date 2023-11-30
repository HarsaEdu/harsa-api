package service

import "fmt"

func (paymentService *PaymentServiceImpl) NotificationPayment(notificationPayload map[string]interface{}) error {
	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		return fmt.Errorf("error when get order id : order id not found")
	}

	response, err := paymentService.MidtransCoreApi.CheckTransactionStatus(orderId)
	if err != nil {
		return fmt.Errorf("error when checking transaction status : %s", err.Error())
	}

	fmt.Println(response)

	return nil
}