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
	CreatedAt           *time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

type User struct {
	ID           uint `gorm:"primarykey"`
	Role         string
	Username     string
	Password     string
	Email        string `gorm:"unique"`
	PhoneNumber  string
	IsVerified   bool
	Transactions []Transaction
	CreatedAt    *time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
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

type Product struct {
	ID            uint   `gorm:"primarykey"`
	Name          string `gorm:"unique"`
	Description   string
	CategoryID    uint
	Category      Category
	Transaction   []DetailTransaction
	Price         int
	Stocks        int
	Sold          int
	SubCategoryID uint
	SubCategory   SubCategory
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type SubCategory struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `gorm:"unique"`
	Tax      int
	ImageURL string
	//Product []Product
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Category struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomainDetailTransaction(domain transaction.DetailDomain) DetailTransaction {
	return DetailTransaction{
		ID:        domain.ID,
		ProductID: domain.ProductID,
		Discount:  domain.Discount,
		Subtotal:  domain.Subtotal,
	}
}

func (d *DetailTransaction) ToDomainDetail() transaction.DetailDomain {
	return transaction.DetailDomain{
		ID:        d.ID,
		ProductID: d.ProductID,
		Discount:  d.Discount,
		Subtotal:  d.Subtotal,
	}
}

func FromDomainTransaction(domain transaction.Domain) Transaction {
	return Transaction{
		ID:                  domain.ID,
		UserID:              domain.UserID,
		DetailTransactionID: domain.DetailTransactionID,
		Total:               domain.Total,
		Link:                domain.Link,
		TransactionStatus:   domain.TransactionStatus,
		FraudStatus:         domain.FraudStatus,
		PaymentType:         domain.PaymentType,
		Provider:            domain.Provider,
	}
}

func (t *Transaction) ToDomain() transaction.Domain {
	return transaction.Domain{
		ID:                  t.ID,
		UserID:              t.UserID,
		DetailTransactionID: t.DetailTransactionID,
		Total:               t.Total,
		Link:                t.Link,
		TransactionStatus:   t.TransactionStatus,
		FraudStatus:         t.FraudStatus,
		PaymentType:         t.PaymentType,
		Provider:            t.Provider,
	}
}

func (t *Transaction) ToHistoryDomain() transaction.HistoryDomain {
	return transaction.HistoryDomain{
		ID:                t.ID,
		ProductID:         t.DetailTransaction.ProductID,
		ProductName:       "",
		Discount:          t.DetailTransaction.Discount,
		Tax:               0,
		Total:             t.Total,
		Link:              t.Link,
		TransactionStatus: t.TransactionStatus,
		FraudStatus:       t.FraudStatus,
		PaymentType:       t.PaymentType,
		Provider:          t.Provider,
		CreatedAt: t.DetailTransaction.CreatedAt,
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

func ToHistoryDomainList(ts []Transaction) []transaction.HistoryDomain {
	var dummyDomain []transaction.HistoryDomain
	for x := range ts {
		dummyProducts := ts[x].ToHistoryDomain()
		dummyDomain = append(dummyDomain, dummyProducts)
	}
	return dummyDomain
}
