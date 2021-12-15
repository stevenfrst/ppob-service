package product

import (
	"gorm.io/gorm"
	"ppob-service/drivers/repository/transaction"
	"ppob-service/usecase/product"
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

func (p *Product) ToDomain() product.Domain {
	return product.Domain{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		CategoryID:  p.CategoryID,
		Price:       p.Price,
		Stocks:      p.Stocks,
		Discount:    p.Discount,
	}
}