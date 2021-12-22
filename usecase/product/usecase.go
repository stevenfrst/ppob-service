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

func (p *ProductUsecase) GetProduct(id int) ([]Domain, error) {
	resp, err := p.repo.GetProduct(id)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (p *ProductUsecase) EditProduct(item Domain) error {
	err := p.repo.EditProduct(item)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductUsecase) Delete(id int) error {
	err := p.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductUsecase) GetBestSellerCategory(id int) (resp []Domain, err error) {
	for x := 0; x < 5; x++ {
		rep, _ := p.repo.GetBestSellerCategory(id, x)
		resp = append(resp, rep)
		//log.Println(id, x)
	}
	if err != nil {
		return []Domain{}, err
	} else if resp[0].ID == 0 {
		resp, _ = p.repo.GetBestSellerCategorySQL(id)
	}
	return resp, nil
}
