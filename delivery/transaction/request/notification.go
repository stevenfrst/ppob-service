package request

import "ppob-service/usecase/transaction"

type NotificationReq struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}

func (c NotificationReq) ToDomainNotification() transaction.Notification {
	return transaction.Notification{
		TransactionStatus: c.TransactionStatus,
		OrderID:           c.OrderID,
		PaymentType:       c.PaymentType,
		FraudStatus:       c.FraudStatus,
	}
}
