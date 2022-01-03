package request

import "ppob-service/usecase/product"

type SubCategory struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Tax      int    `json:"tax"`
	ImageURL string `json:"image_url"`
}

func (r *SubCategory) ToDomainSubCategory() product.SubCategory {
	return product.SubCategory{
		ID:       r.ID,
		Name:     r.Name,
		Tax:      r.Tax,
		ImageURL: r.ImageURL,
	}
}
