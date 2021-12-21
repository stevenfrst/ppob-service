package transaction

import "time"

type Domain struct {
	ID                uint
	UserID            uint
	ProductID         uint
	Nominal           int
	Link              string
	TransactionStatus string
	FraudStatus       string
	PaymentType       string
	Provider          string
	CreatedAt         time.Time
}

type Notification struct {
	TransactionStatus string
	OrderID           string
	PaymentType       string
	FraudStatus       string
}

type ITransactionUsecase interface {
}

type ITransactionRepository interface {
	Create(input Domain) error
	GetByID(ID int) (Domain, error)
	Update(input Domain) error
	GetByUserID(id int) ([]Domain, error)
}
