package request

import "ppob-service/usecase/transaction"

type CreatePayment struct {
	ProductID uint `json:"product_id"`
	Discount  int `json:"discount"`
	Tax       int `json:"tax"`
	Subtotal  int `json:"subtotal"`
	Bank      string `json:"bank"`
}

func (c CreatePayment) ToDomainPayment() transaction.CreateVA {
	return transaction.CreateVA{
		ProductID: c.ProductID,
		Discount:  c.Discount,
		Tax:       c.Tax,
		Subtotal:  c.Subtotal,
		Bank:      c.Bank,
	}
}
