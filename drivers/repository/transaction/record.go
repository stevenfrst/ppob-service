package transaction

import (
	"gorm.io/gorm"
	"ppob-service/usecase/transaction"
	"time"
)

type Transaction struct {
	ID                  uint `gorm:"primarykey"`
	UserID              uint
	DetailTransactionID uint
	DetailTransaction   DetailTransaction
	Total               int
	Link                string
	TransactionStatus   string
	FraudStatus         string
	PaymentType         string
	Provider            string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

type DetailTransaction struct {
	ID        uint `gorm:"primarykey"`
	ProductID uint
	Discount  int
	Subtotal  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomainTransaction(domain transaction.Domain) Transaction {
	return Transaction{
		ID:                domain.ID,
		UserID:            domain.UserID,
		Total:             domain.Total,
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
		Total:             t.Total,
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
