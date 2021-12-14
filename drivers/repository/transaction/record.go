package transaction

import (
	"gorm.io/gorm"
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
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
