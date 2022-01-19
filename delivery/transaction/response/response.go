package response

import (
	"ppob-service/usecase/transaction"
)

type History struct {
	ID                uint   `json:"id"`
	ProductName       string `json:"product_name"`
	Discount          int    `json:"discount"`
	Tax               int    `json:"tax"`
	Total             int    `json:"total"`
	Link              string `json:"link"`
	TransactionStatus string `json:"transaction_status"`
	FraudStatus       string `json:"fraud_status"`
	PaymentType       string `json:"payment_type"`
	Provider          string `json:"provider"`
	CreatedAt         string `json:"created_at"`
}

func FromDomain(hd transaction.HistoryDomain) History {
	hd.CreatedAt.Format("2006-01-02 15:04:05")
	return History{
		ID:                hd.ID,
		ProductName:       hd.ProductName,
		Discount:          hd.Discount,
		Tax:               hd.Tax,
		Total:             hd.Total,
		Link:              hd.Link,
		TransactionStatus: hd.TransactionStatus,
		FraudStatus:       hd.FraudStatus,
		PaymentType:       hd.PaymentType,
		Provider:          hd.Provider,
		CreatedAt:         hd.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func FromDomainList(tx []transaction.HistoryDomain) []History {
	var dummyProducts []History
	for x := range tx {
		dummyDomain := FromDomain(tx[x])
		dummyProducts = append(dummyProducts, dummyDomain)
	}
	return dummyProducts
}
