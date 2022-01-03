package product

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
	"math/rand"
	"time"
)

type ProductUsecase struct {
	repo IProductRepository
	s3   *minio.Client
	url  string
}

func NewUseCase(productRepo IProductRepository, s3 *minio.Client, url string) IProductUsecase {
	return &ProductUsecase{
		repo: productRepo,
		s3:   s3,
		url:  url,
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

func (p *ProductUsecase) Create(domain CreateDomain) error {
	err := p.repo.Create(domain)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductUsecase) GetAll(offset, pageSize int) ([]Domain, error) {
	resp, err := p.repo.GetAllProductPagination(offset, pageSize)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (p *ProductUsecase) GetAllCategory() []Category {
	return p.repo.GetAllCategory()
}

func (p *ProductUsecase) GetAllSubCategory() []SubCategory {
	return p.repo.GetAllSubCategory()
}

func (p *ProductUsecase) EditSubCategory(edit SubCategory) error {
	return p.repo.EditSubCategory(edit)
}

func (p *ProductUsecase) CreateCategory(category Category) error {
	return p.repo.CreateCategory(category)
}

func (p *ProductUsecase) DeleteCategory(id int) error {
	return p.repo.DeleteCategory(id)
}

func (p *ProductUsecase) DeleteSubCategory(id int) error {
	return p.repo.DeleteSubCategory(id)
}

func (p *ProductUsecase) CreateSubCategory(sub SubCategory, objName string) error {
	if _, err := p.s3.FPutObject(context.Background(), "static", objName, sub.ImageURL, minio.PutObjectOptions{
		ContentType: "image/png",
	}); err != nil {
		log.Println(err)
		return err
	}

	sub.ImageURL = fmt.Sprintf("http://%v/static/%v", p.url, objName)
	log.Println(sub.ImageURL)
	err := p.repo.CreateSubCategory(sub)
	if err != nil {
		return err
	}
	return nil
}
