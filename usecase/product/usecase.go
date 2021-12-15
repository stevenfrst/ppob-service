package product

import (
	"math/rand"
	"time"
)

type ProductUsecase struct {
	repo IProductRepository
}

func NewUseCase(productRepo IProductRepository) IProductUsecase {
	return &ProductUsecase{
		repo: productRepo,
	}
}

func (p *ProductUsecase) GetTagihanPLN() (Domain, error) {
	count, _ := p.repo.CountItem(3)
	rand.Seed(time.Now().UTC().UnixNano())
	id := rand.Intn(count)
	product, err := p.repo.GetTagihanPLN(id)
	if err != nil {
		return Domain{}, err
	}
	return product, nil
}

func (p *ProductUsecase) GetProduct(id int) ([]Domain,error) {
	resp, err := p.repo.GetProduct(id)
	if err != nil {
		return []Domain{},err
	}
	return resp,nil
}