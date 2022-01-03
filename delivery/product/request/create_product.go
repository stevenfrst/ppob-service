package request

import "ppob-service/usecase/product"

type CreateProduct struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	CategoryID    uint   `json:"category_id"`
	Price         int    `json:"price"`
	Stocks        int    `json:"stocks"`
	SubCategoryID uint   `json:"sub_category"`
}

func (c *CreateProduct) ToDomain() product.CreateDomain {
	return product.CreateDomain{
		Name:          c.Name,
		Description:   c.Description,
		CategoryID:    c.CategoryID,
		Price:         c.Price,
		Stocks:        c.Stocks,
		SubCategoryID: c.SubCategoryID,
	}
}
