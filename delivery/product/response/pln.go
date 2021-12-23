package response

import "ppob-service/usecase/product"

type PLN struct {
	ID          uint `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int `json:"price"`
	Tax         int `json:"tax"`
}

func FromDomain(domain product.Domain) PLN {
	return PLN{
		ID:          domain.ID,
		Name:        domain.Name,
		Description: domain.Description,
		Category:    domain.Category,
		Price:       domain.Price,
		Tax: domain.Tax,
	}
}

func FromDomainList(products []product.Domain) []PLN {
	var dummyProducts []PLN
	for x := range products {
		dummyDomain := FromDomain(products[x])
		dummyProducts = append(dummyProducts,dummyDomain)
	}
	return dummyProducts
}
