package response

import "ppob-service/usecase/product"

type PLN struct {
	ID          uint
	Name        string
	Description string
	CategoryID  uint
	Price       int
}

func FromDomain(domain product.Domain) PLN {
	return PLN{
		ID:          domain.ID,
		Name:        domain.Name,
		Description: domain.Description,
		CategoryID:  domain.CategoryID,
		Price:       domain.Price,
	}
}