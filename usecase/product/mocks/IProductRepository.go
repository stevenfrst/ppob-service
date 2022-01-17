// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	product "ppob-service/usecase/product"

	mock "github.com/stretchr/testify/mock"
)

// IProductRepository is an autogenerated mock type for the IProductRepository type
type IProductRepository struct {
	mock.Mock
}

// CountItem provides a mock function with given fields: category
func (_m *IProductRepository) CountItem(category int) (int, error) {
	ret := _m.Called(category)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(category)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: input
func (_m *IProductRepository) Create(input product.CreateDomain) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(product.CreateDomain) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateCategory provides a mock function with given fields: category
func (_m *IProductRepository) CreateCategory(category product.Category) error {
	ret := _m.Called(category)

	var r0 error
	if rf, ok := ret.Get(0).(func(product.Category) error); ok {
		r0 = rf(category)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateSubCategory provides a mock function with given fields: domain
func (_m *IProductRepository) CreateSubCategory(domain product.SubCategory) error {
	ret := _m.Called(domain)

	var r0 error
	if rf, ok := ret.Get(0).(func(product.SubCategory) error); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *IProductRepository) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCategory provides a mock function with given fields: id
func (_m *IProductRepository) DeleteCategory(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSubCategory provides a mock function with given fields: id
func (_m *IProductRepository) DeleteSubCategory(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditProduct provides a mock function with given fields: item
func (_m *IProductRepository) EditProduct(item product.Domain) error {
	ret := _m.Called(item)

	var r0 error
	if rf, ok := ret.Get(0).(func(product.Domain) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditSubCategory provides a mock function with given fields: sub
func (_m *IProductRepository) EditSubCategory(sub product.SubCategory) error {
	ret := _m.Called(sub)

	var r0 error
	if rf, ok := ret.Get(0).(func(product.SubCategory) error); ok {
		r0 = rf(sub)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllCategory provides a mock function with given fields:
func (_m *IProductRepository) GetAllCategory() []product.Category {
	ret := _m.Called()

	var r0 []product.Category
	if rf, ok := ret.Get(0).(func() []product.Category); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Category)
		}
	}

	return r0
}

// GetAllProduct provides a mock function with given fields:
func (_m *IProductRepository) GetAllProduct() ([]product.Domain, error) {
	ret := _m.Called()

	var r0 []product.Domain
	if rf, ok := ret.Get(0).(func() []product.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllProductPagination provides a mock function with given fields: offset, pageSize
func (_m *IProductRepository) GetAllProductPagination(offset int, pageSize int) ([]product.Domain, error) {
	ret := _m.Called(offset, pageSize)

	var r0 []product.Domain
	if rf, ok := ret.Get(0).(func(int, int) []product.Domain); ok {
		r0 = rf(offset, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(offset, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllSubCategory provides a mock function with given fields:
func (_m *IProductRepository) GetAllSubCategory() []product.SubCategory {
	ret := _m.Called()

	var r0 []product.SubCategory
	if rf, ok := ret.Get(0).(func() []product.SubCategory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.SubCategory)
		}
	}

	return r0
}

// GetBestSellerCategory provides a mock function with given fields: id, item
func (_m *IProductRepository) GetBestSellerCategory(id int, item int) (product.Domain, error) {
	ret := _m.Called(id, item)

	var r0 product.Domain
	if rf, ok := ret.Get(0).(func(int, int) product.Domain); ok {
		r0 = rf(id, item)
	} else {
		r0 = ret.Get(0).(product.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(id, item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBestSellerCategorySQL provides a mock function with given fields: id
func (_m *IProductRepository) GetBestSellerCategorySQL(id int) ([]product.Domain, error) {
	ret := _m.Called(id)

	var r0 []product.Domain
	if rf, ok := ret.Get(0).(func(int) []product.Domain); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProduct provides a mock function with given fields: id
func (_m *IProductRepository) GetProduct(id int) ([]product.Domain, error) {
	ret := _m.Called(id)

	var r0 []product.Domain
	if rf, ok := ret.Get(0).(func(int) []product.Domain); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTagihanPLN provides a mock function with given fields: id
func (_m *IProductRepository) GetTagihanPLN(id int) (product.Domain, error) {
	ret := _m.Called(id)

	var r0 product.Domain
	if rf, ok := ret.Get(0).(func(int) product.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(product.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}