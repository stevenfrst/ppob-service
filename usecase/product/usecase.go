package product

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/minio/minio-go/v7"
	"math/rand"
	"ppob-service/helpers/errorHelper"
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
	} else if resp[0].ID == 0 {
		return []Domain{}, errorHelper.ErrRecordNotFound
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

func (p *ProductUsecase) Create(domain CreateDomain) error {
	err := p.repo.Create(domain)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return errorHelper.DuplicateData
	} else if err != nil {
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
	err := p.repo.CreateCategory(category)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return errorHelper.DuplicateData
	} else if err != nil {
		return err
	}
	return nil
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
		return err
	}

	sub.ImageURL = fmt.Sprintf("http://%v/static/%v", p.url, objName)
	err := p.repo.CreateSubCategory(sub)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return errorHelper.DuplicateData
	} else if err != nil {
		return err
	}
	return nil
}
