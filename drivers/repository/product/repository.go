package product

import (
	"gorm.io/gorm"
	"log"
	"math/rand"
	"ppob-service/usecase/product"
	"time"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewRepository(gormDB *gorm.DB) product.IProductRepository {
	return &ProductRepository{
		db: gormDB,
	}
}

func (p *ProductRepository) CountItem(category int) (int, error) {
	var repoModel []Product
	var count int64
	err := p.db.Model(&repoModel).Where("category_id = ?", category).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (p *ProductRepository) GetTagihanPLN(id int) (product.Domain, error) {
	var repoModel []Product
	err := p.db.Preload("Category").Where("category_id = ?", 3).Find(&repoModel).Error
	if err != nil {
		return product.Domain{}, err
	}
	log.Println(repoModel[0].Category.Name)
	rand.Seed(time.Now().UTC().UnixNano())
	return repoModel[id].ToDomain(), nil
}

func (p *ProductRepository) GetProduct(id int) ([]product.Domain, error) {
	var repoModel []Product
	err := p.db.Preload("Category").Where("category_id = ?", id).Find(&repoModel).Error
	if err != nil {
		return ToDomainList([]Product{}), err
	}
	return ToDomainList(repoModel),nil
}