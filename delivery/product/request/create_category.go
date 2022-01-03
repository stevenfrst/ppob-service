package request

import "ppob-service/usecase/product"

type Category struct {
	Name string `json:"name" validate:"required"`
}

func (r *Category) ToDomainCategory() product.Category {
	return product.Category{
		Name: r.Name,
	}
}
