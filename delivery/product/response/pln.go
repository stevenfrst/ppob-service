package response

import "ppob-service/usecase/product"

type Product struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Tax         int    `json:"tax"`
	Stocks      int    `json:"stocks"`
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
