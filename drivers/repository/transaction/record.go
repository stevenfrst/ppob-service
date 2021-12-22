package transaction

import (
	"gorm.io/gorm"
	"ppob-service/usecase/transaction"
	"time"
)

type Transaction struct {
	ID                uint `gorm:"primarykey"`
	UserID            uint
	ProductID         uint
	Nominal           int
	Link              string
	TransactionStatus string
	FraudStatus       string
	PaymentType       string
	Provider          string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

func FromDomainTransaction(domain transaction.Domain) Transaction {
	return Transaction{
		ID:                domain.ID,
		UserID:            domain.UserID,
		ProductID:         domain.ProductID,
		Nominal:           domain.Nominal,
		Link:              domain.Link,
		TransactionStatus: domain.TransactionStatus,
		FraudStatus:       domain.FraudStatus,
		PaymentType:       domain.PaymentType,
		Provider:          domain.Provider,
	}
}

func (t *Transaction) ToDomain() transaction.Domain {
	return transaction.Domain{
		ID:                t.ID,
		UserID:            t.UserID,
		ProductID:         t.ProductID,
		Nominal:           t.Nominal,
		Link:              t.Link,
		TransactionStatus: t.TransactionStatus,
		FraudStatus:       t.FraudStatus,
		PaymentType:       t.PaymentType,
		Provider:          t.Provider,
		CreatedAt:         t.CreatedAt,
	}
}

func ToDomainList(ts []Transaction) []transaction.Domain {
	var dummyDomain []transaction.Domain
	for x := range ts {
		dummyProducts := ts[x].ToDomain()
		dummyDomain = append(dummyDomain, dummyProducts)
	}
	return dummyDomain
}
