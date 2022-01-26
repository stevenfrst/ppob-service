package product

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"ppob-service/delivery/product/response"
	"ppob-service/usecase/product"
	"time"
)

type ProductRepository struct {
	db    *gorm.DB
	cache redis.Conn
}

func NewRepository(gormDB *gorm.DB, cache redis.Conn) product.IProductRepository {
	return &ProductRepository{
		db:    gormDB,
		cache: cache,
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
	err := p.db.Preload("Category").Preload("SubCategory").Where("category_id = ?", 3).Find(&repoModel).Error
	if err != nil {
		return product.Domain{}, err
	}
	log.Println(repoModel[0].Category.Name)
	rand.Seed(time.Now().UTC().UnixNano())
	return repoModel[id].ToDomain(), nil
}

func (p *ProductRepository) GetProduct(id int) ([]product.Domain, error) {
	var repoModel []Product
	err := p.db.Preload("Category").Preload("SubCategory").Where("category_id = ?", id).Find(&repoModel).Error
	if err != nil {
		return ToDomainList([]Product{}), err
	}
	return ToDomainList(repoModel), nil
}

func (p *ProductRepository) GetAllProduct() ([]product.Domain, error) {
	var repoModel []Product
	err := p.db.Preload("Category").Preload("SubCategory").Find(&repoModel).Error
	if err != nil {
		return ToDomainList([]Product{}), err
	}
	return ToDomainList(repoModel), nil
}

func (p *ProductRepository) GetProductByID(id int) Product {
	var repoModel Product
	p.db.Preload("Category").Preload("SubCategory").Where("id = ?", id).First(&repoModel)
	return repoModel
}

func (p *ProductRepository) EditProduct(item product.Domain) error {
	log.Println(item)
	var repoModel Product
	repoModel = p.GetProductByID(int(item.ID))

	repoModel.ID = item.ID
	repoModel.Name = item.Name
	repoModel.Price = item.Price
	repoModel.Description = item.Description
	repoModel.Stocks = item.Stocks

	err := p.db.Save(&repoModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) Delete(id int) error {
	var repoModel Product
	log.Println(id)
	err := p.db.Where("id = ?", id).Delete(&repoModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) GetBestSellerCategory(id, item int) (product.Domain, error) {
	category := fmt.Sprintf("product_%v:%v", item, id)
	//log.Println(category,"+++++")
	rep, err := redis.Values(p.cache.Do("HGETALL", category))
	if err != nil {
		return product.Domain{}, err
	}

	var repoModel product.Domain

	err = redis.ScanStruct(rep, &repoModel)
	if err != nil {
		return product.Domain{}, err
	}
	//log.Println(rep,"++++++")
	log.Println(repoModel)
	return repoModel, err
}

func (p *ProductRepository) GetBestSellerCategorySQL(id int) ([]product.Domain, error) {
	var repoModels []Product
	err := p.db.Preload("Category").Order("sold DESC").Where("category_id = ?", id).Find(&repoModels).Limit(5).Error
	if err != nil {
		return ToDomainList(repoModels), err
	}
	return ToDomainList(repoModels), nil
}

func (p *ProductRepository) Create(input product.CreateDomain) error {
	repoModel := FromDomain(input)
	err := p.db.Create(&repoModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) GetAllProductPagination(offset, pageSize int) ([]product.Domain, error) {
	var repoModels []Product
	err := p.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}).Preload("Category").Preload("SubCategory").Find(&repoModels).Error
	if err != nil {
		return ToDomainList([]Product{}), err
	}
	return ToDomainList(repoModels), nil
}

func (p *ProductRepository) GetAllCategory() []product.Category {
	var repoModels []Category
	p.db.Find(&repoModels)
	return ToCategoryList(repoModels)
}

func (p *ProductRepository) GetAllSubCategory() []product.SubCategory {
	var repoModels []SubCategory
	p.db.Find(&repoModels)
	return ToSubCategoryList(repoModels)
}

func (p *ProductRepository) EditSubCategory(sub product.SubCategory) error {
	repoModels := response.FromSubDomainCategory(sub)
	err := p.db.Save(&repoModels).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) CreateCategory(category product.Category) error {
	var repoModels Category
	repoModels.Name = category.Name
	err := p.db.Create(&repoModels).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) DeleteCategory(id int) error {
	var repoModel Category
	err := p.db.Where("id = ?", id).Delete(&repoModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) DeleteSubCategory(id int) error {
	var repoModel SubCategory
	err := p.db.Where("id = ?", id).Delete(&repoModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) CreateSubCategory(domain product.SubCategory) error {
	repoModel := FromDomainSubCategory(domain)
	err := p.db.Create(&repoModel).Error
	if err != nil {
		return err
	}
	return nil
}