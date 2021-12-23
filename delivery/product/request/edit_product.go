package request

import "ppob-service/usecase/product"

type EditProduct struct {
	ID          uint   `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Stocks      int    `json:"stocks" validate:"required"`
	Tax         int    `json:"tax"`
}

func (edit EditProduct) ToDomain() product.Domain {
	return product.Domain{
		ID:          edit.ID,
		Name:        edit.Name,
		Description: edit.Description,
		Price:       edit.Price,
		Stocks:      edit.Stocks,
		Tax:         edit.Tax,
	}
}
