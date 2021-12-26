package product

import (
	"gorm.io/gorm"
	"ppob-service/drivers/repository/transaction"
	"ppob-service/usecase/product"
	"time"
)

type Product struct {
	ID            uint `gorm:"primarykey"`
	Name          string
	Description   string
	CategoryID    uint
	Category      Category
	Transaction   []transaction.DetailTransaction
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
	ID       uint `gorm:"primarykey"`
	Name     string
	Tax      int
	ImageURL string
	//Product []Product
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Category struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain product.CreateDomain) Product {
	return Product{
		ID:            domain.ID,
		Name:          domain.Name,
		Description:   domain.Description,
		CategoryID:    domain.CategoryID,
		Price:         domain.Price,
		Stocks:        domain.Stocks,
		SubCategoryID: domain.SubCategoryID,
	}
}

func (p *Product) ToDomain() product.Domain {
	return product.Domain{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Category:    p.Category.Name,
		Price:       p.Price,
		Stocks:      p.Stocks,
		Tax:         p.SubCategory.Tax,
		SubCategory: p.SubCategory.Name,
		Link:        p.SubCategory.ImageURL,
	}
}

func (p *Category) ToDomain() product.Category {
	return product.Category{
		ID:   p.ID,
		Name: p.Name,
	}
}

func ToCategoryList(category []Category) []product.Category {
	var dummyDomain []product.Category
	for x := range category {
		dummyProducts := category[x].ToDomain()
		dummyDomain = append(dummyDomain, dummyProducts)
	}
	return dummyDomain
}

func (p *SubCategory) ToDomain() product.SubCategory {
	return product.SubCategory{
		ID:   p.ID,
		Name: p.Name,
		Tax:  p.Tax,
		ImageURL: p.ImageURL,
	}
}

func ToSubCategoryList(sub []SubCategory) []product.SubCategory {
	var dummyDomain []product.SubCategory
	for x := range sub {
		dummyProducts := sub[x].ToDomain()
		dummyDomain = append(dummyDomain, dummyProducts)
	}
	return dummyDomain
}

func ToDomainList(products []Product) []product.Domain {
	var dummyDomain []product.Domain
	for x := range products {
		dummyProducts := products[x].ToDomain()
		dummyDomain = append(dummyDomain, dummyProducts)
	}
	return dummyDomain
}
