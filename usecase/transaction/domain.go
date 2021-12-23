package transaction

import (
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	ID                uint `gorm:"primarykey"`
	UserID            uint
	Total             int
	Link              string
	TransactionStatus string
	FraudStatus       string
	PaymentType       string
	Provider          string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
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
