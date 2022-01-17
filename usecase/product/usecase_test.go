package product_test

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ppob-service/app/config"
	storagedriver "ppob-service/drivers/s3"
	"ppob-service/usecase/product"
	"ppob-service/usecase/product/mocks"
	"testing"
)

var productMockRepo mocks.IProductRepository
var productUseCase product.IProductUsecase
var productDomain product.Domain
var productsDomain []product.Domain
var createDomain product.CreateDomain
var mysqlErr mysql.MySQLError
var productSubCat product.SubCategory
var productCat product.Category

func Setup() {
	mysqlErr.Number = 1062
	getConfig := config.GetConfigTest()
	s3Config := storagedriver.MinioService{
		Host:     getConfig.STORAGE_URL,
		Username: getConfig.STORAGE_ID,
		Secret:   getConfig.STORAGE_SECRET,
	}
	s3 := s3Config.NewClient()
	productUseCase = product.NewUseCase(&productMockRepo, s3, getConfig.STORAGE_URL)
	productDomain = product.Domain{
		ID:          1,
		Name:        "Test",
		Description: "Test",
		Category:    "Test",
		Price:       1,
		Stocks:      1,
		Tax:         1,
		SubCategory: "Test",
		Link:        "",
	}
	productsDomain = append(productsDomain, productDomain)

	createDomain = product.CreateDomain{
		ID:            1,
		Name:          "test",
		Description:   "test",
		CategoryID:    1,
		Price:         1,
		Stocks:        1,
		SubCategoryID: 1,
	}
	productSubCat = product.SubCategory{
		ID:       1,
		Name:     "test",
		Tax:      0,
		ImageURL: "temp/test.png",
	}
	productCat = product.Category{
		ID:   1,
		Name: "test",
	}
}

func TestCreateSubCategory(t *testing.T) {
	Setup()
	t.Run("errors create", func(t *testing.T) {
		productMockRepo.On("CreateSubCategory",
			mock.AnythingOfType("product.SubCategory"),
		).Return(nil).Once()
		err := productUseCase.CreateSubCategory(productSubCat, "salkomsel")
		assert.Error(t, err)
	})
}

func TestDeleteSubCategory(t *testing.T) {
	Setup()
	t.Run("functional", func(t *testing.T) {
		productMockRepo.On("DeleteSubCategory",
			mock.AnythingOfType("int"),
		).Return(nil).Once()
		err := productUseCase.DeleteSubCategory(1)
		assert.Nil(t, err)
	})
}

func TestDeleteCategory(t *testing.T) {
	Setup()
	t.Run("functional", func(t *testing.T) {
		productMockRepo.On("DeleteCategory",
			mock.AnythingOfType("int"),
		).Return(nil).Once()
		err := productUseCase.DeleteCategory(1)
		assert.Nil(t, err)
	})
}

func TestCreateCategory(t *testing.T) {
	Setup()
	t.Run("Success Create a Data", func(t *testing.T) {
		productMockRepo.On("CreateCategory",
			mock.AnythingOfType("product.Category"),
		).Return(nil).Once()
		err := productUseCase.CreateCategory(productCat)
		assert.Nil(t, err)
	})
	t.Run("failed Create a Data", func(t *testing.T) {
		productMockRepo.On("CreateCategory",
			mock.AnythingOfType("product.Category"),
		).Return(&mysqlErr).Once()
		err := productUseCase.CreateCategory(productCat)
		assert.Error(t, err)
	})
	t.Run("failed db error", func(t *testing.T) {
		productMockRepo.On("CreateCategory",
			mock.AnythingOfType("product.Category"),
		).Return(errors.New("some dbs errors")).Once()
		err := productUseCase.CreateCategory(productCat)
		assert.Error(t, err)
	})
}

func TestEditSubCategory(t *testing.T) {
	Setup()
	t.Run("functional", func(t *testing.T) {
		productMockRepo.On("EditSubCategory",
			mock.AnythingOfType("product.SubCategory"),
		).Return(nil).Once()
		resp := productUseCase.EditSubCategory(productSubCat)
		assert.Nil(t, resp)
	})
}

func TestGetAllSubCategory(t *testing.T) {
	Setup()
	t.Run("functional", func(t *testing.T) {
		productMockRepo.On("GetAllSubCategory").Return([]product.SubCategory{}).Once()
		resp := productUseCase.GetAllSubCategory()
		assert.Equal(t, resp, []product.SubCategory{})
	})
}

func TestGetAllCategory(t *testing.T) {
	Setup()
	t.Run("functional", func(t *testing.T) {
		productMockRepo.On("GetAllCategory").Return([]product.Category{}).Once()
		resp := productUseCase.GetAllCategory()
		assert.Equal(t, resp, []product.Category{})
	})
}

func TestGetAll(t *testing.T) {
	Setup()
	t.Run("Success get data", func(t *testing.T) {
		productMockRepo.On("GetAllProductPagination",
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int"),
		).Return(productsDomain, nil).Once()
		resp, err := productUseCase.GetAll(100, 25)
		assert.Nil(t, err)
		assert.Equal(t, productsDomain, resp)
	})
	t.Run("failed get data", func(t *testing.T) {
		productMockRepo.On("GetAllProductPagination",
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int"),
		).Return([]product.Domain{}, errors.New("some db errors")).Once()
		resp, err := productUseCase.GetAll(100, 25)
		assert.Error(t, err)
		assert.Equal(t, []product.Domain{}, resp)
	})

}

func TestCreate(t *testing.T) {
	Setup()
	t.Run("Success Create", func(t *testing.T) {
		productMockRepo.On("Create",
			mock.AnythingOfType("product.CreateDomain"),
		).Return(nil).Once()
		err := productUseCase.Create(createDomain)
		assert.Nil(t, err)
	})
	t.Run("duplicate data", func(t *testing.T) {
		productMockRepo.On("Create",
			mock.AnythingOfType("product.CreateDomain"),
		).Return(&mysqlErr).Once()
		err := productUseCase.Create(createDomain)
		assert.Error(t, err)
	})
	t.Run("error db", func(t *testing.T) {
		productMockRepo.On("Create",
			mock.AnythingOfType("product.CreateDomain"),
		).Return(errors.New("some random db's errors")).Once()
		err := productUseCase.Create(createDomain)
		assert.Error(t, err)
	})

}

func TestDelete(t *testing.T) {
	Setup()
	t.Run("Success Delete Product", func(t *testing.T) {
		productMockRepo.On("Delete",
			mock.AnythingOfType("int"),
		).Return(nil).Once()
		err := productUseCase.Delete(1)
		assert.Nil(t, err)
	})

	t.Run("Failed Delete Product", func(t *testing.T) {
		productMockRepo.On("Delete",
			mock.AnythingOfType("int"),
		).Return(errors.New("some db err")).Once()
		err := productUseCase.Delete(1)
		assert.Error(t, err)
	})

}

func TestEditProduct(t *testing.T) {
	Setup()
	t.Run("Success Edit Product", func(t *testing.T) {
		productMockRepo.On("EditProduct",
			mock.AnythingOfType("product.Domain"),
		).Return(nil).Once()

		err := productUseCase.EditProduct(productDomain)
		assert.Nil(t, err)
	})

	t.Run("Error Edit Product", func(t *testing.T) {
		productMockRepo.On("EditProduct",
			mock.AnythingOfType("product.Domain"),
		).Return(errors.New("db errors")).Once()

		err := productUseCase.EditProduct(productDomain)
		assert.Error(t, err)
	})
}

func TestGetProduct(t *testing.T) {
	Setup()
	t.Run("Success Getting Product", func(t *testing.T) {
		productMockRepo.On("GetProduct",
			mock.AnythingOfType("int"),
		).Return(productsDomain, nil).Once()

		resp, err := productUseCase.GetProduct(3)
		assert.Nil(t, err)
		assert.Equal(t, resp, productsDomain)
	})
	t.Run("Failed Getting Product | not found", func(t *testing.T) {
		productMockRepo.On("GetProduct",
			mock.AnythingOfType("int"),
		).Return([]product.Domain{{}}, nil).Once()

		_, err := productUseCase.GetProduct(1)
		assert.Error(t, err)
	})
	t.Run("Failed Getting Product | db err", func(t *testing.T) {
		productMockRepo.On("GetProduct",
			mock.AnythingOfType("int"),
		).Return([]product.Domain{{}}, errors.New("some db errors")).Once()

		_, err := productUseCase.GetProduct(3)
		assert.Error(t, err)
	})
}

func TestGetTagihanPLN(t *testing.T) {
	Setup()
	t.Run("Success Getting Product", func(t *testing.T) {
		productMockRepo.On("CountItem",
			mock.AnythingOfType("int"),
		).Return(3, nil).Once()

		productMockRepo.On("GetTagihanPLN",
			mock.AnythingOfType("int"),
		).Return(productDomain, nil).Once()
		resp, err := productUseCase.GetTagihanPLN()
		assert.Nil(t, err)
		assert.Equal(t, resp, productDomain)
	})

	t.Run("Success Getting Product", func(t *testing.T) {
		productMockRepo.On("CountItem",
			mock.AnythingOfType("int"),
		).Return(3, nil).Once()

		productMockRepo.On("GetTagihanPLN",
			mock.AnythingOfType("int"),
		).Return(product.Domain{}, errors.New("some random db's error")).Once()
		resp, err := productUseCase.GetTagihanPLN()
		assert.Error(t, err)
		assert.Equal(t, resp, product.Domain{})
	})
}
