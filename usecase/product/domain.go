package product

import (
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	Description string
	Category    string
	//Transaction []transaction.DetailTransaction
	Price     int
	Stocks    int
	Tax       int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type IProductUsecase interface {
	GetTagihanPLN() (Domain, error)
	GetProduct(id int) ([]Domain, error)
	EditProduct(item Domain) error
	Delete(id int) error
	GetBestSellerCategory(id int) ([]Domain, error)
}

type IProductRepository interface {
	GetTagihanPLN(id int) (Domain, error)
	CountItem(category int) (int, error)
	GetProduct(id int) ([]Domain, error)
	EditProduct(item Domain) error
	Delete(id int) error
	GetBestSellerCategory(id, item int) (Domain, error)
	GetBestSellerCategorySQL(id int) ([]Domain, error)
}
