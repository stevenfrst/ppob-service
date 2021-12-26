package response

import (
	"ppob-service/usecase/product"
)

type Product struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Tax         int    `json:"tax"`
	Stocks      int    `json:"stocks"`
	SubCategory string `json:"sub_category"`
	Link        string `json:"link"`
}

type Category struct {
	ID   uint `json:"id"`
	Name string `json:"name"`
}

func FromDomainCategory(domain product.Category) Category {
	return Category{
		ID:   domain.ID,
		Name: domain.Name,
	}
}

type SubCategory struct {
	ID        uint `json:"id"`
	Name      string `json:"name"`
	Tax       int `json:"tax"`
	ImageURL  string `json:"image_url"`
}

func FromSubDomainCategory(domain product.SubCategory) SubCategory {
	return SubCategory{
		ID:       domain.ID,
		Name:     domain.Name,
		Tax:      domain.Tax,
		ImageURL: domain.ImageURL,
	}
}


func FromDomain(domain product.Domain) Product {
	return Product{
		ID:          domain.ID,
		Name:        domain.Name,
		Description: domain.Description,
		Category:    domain.Category,
		Price:       domain.Price,
		Tax:         domain.Tax,
		Stocks:      domain.Stocks,
		SubCategory: domain.SubCategory,
		Link:        domain.Link,
	}
}

func FromDomainList(products []product.Domain) []Product {
	var dummyProducts []Product
	for x := range products {
		dummyDomain := FromDomain(products[x])
		dummyProducts = append(dummyProducts, dummyDomain)
	}
	return dummyProducts
}


func FromDomainCategoryList(category []product.Category) []Category {
	var dummyProducts []Category
	for x := range category {
		dummyDomain := FromDomainCategory(category[x])
		dummyProducts = append(dummyProducts, dummyDomain)
	}
	return dummyProducts
}

func FromDomainSubCategoryList(sub []product.SubCategory) []SubCategory {
	var dummyProducts []SubCategory
	for x := range sub {
		dummyDomain := FromSubDomainCategory(sub[x])
		dummyProducts = append(dummyProducts, dummyDomain)
	}
	return dummyProducts
}
