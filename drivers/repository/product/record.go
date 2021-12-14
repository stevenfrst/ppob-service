package product

import (
	"gorm.io/gorm"
	"ppob-service/drivers/repository/transaction"
	"time"
)

type Product struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	Description string
	CategoryID  uint
	Category    Category
	Transaction []transaction.Transaction
	Price       int
	Stocks      int
	Discount    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Category struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
