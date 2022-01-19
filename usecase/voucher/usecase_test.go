package vouchers_test

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	vouchers "ppob-service/usecase/voucher"
	"ppob-service/usecase/voucher/mocks"
	"testing"
	"time"
)

var voucherRepositoryMock mocks.IVoucherRepository
var voucherUseCase vouchers.IVoucherUseCase
var voucherDomain vouchers.Domain
var vouchersDomain []vouchers.Domain
var mysqlErr mysql.MySQLError

func Setup() {
	voucherUseCase = vouchers.NewUseCase(&voucherRepositoryMock)
	voucherDomain = vouchers.Domain{
		ID:    1,
		Code:  "Test123",
		Value: 0,
		Valid: time.Time{},
	}
	mysqlErr.Number = 1062
	vouchersDomain = append(vouchersDomain, voucherDomain)
}

func TestVerify(t *testing.T) {
	Setup()
	t.Run("functional", func(t *testing.T) {
		voucherRepositoryMock.On("Verify",
			mock.AnythingOfType("string")).Return(nil).Once()
		err := voucherUseCase.Verify("Test123")
		assert.Nil(t, err)
	})
}

func TestDeleteByID(t *testing.T) {
	Setup()
	t.Run("functional", func(t *testing.T) {
		voucherRepositoryMock.On("DeleteByID",
			mock.AnythingOfType("int")).Return(nil).Once()
		err := voucherUseCase.DeleteByID(1)
		assert.Nil(t, err)
	})
}

func TestReadALL(t *testing.T) {
	Setup()
	t.Run("Success Read All Data", func(t *testing.T) {
		voucherRepositoryMock.On("ReadALL").Return(vouchersDomain, nil).Once()
		resp, err := voucherUseCase.ReadALL()
		assert.Nil(t, err)
		assert.Equal(t, resp, vouchersDomain)
	})

	t.Run("Failed Read All Data -- Not Found", func(t *testing.T) {
		voucherRepositoryMock.On("ReadALL").Return([]vouchers.Domain{{}}, nil).Once()
		_, err := voucherUseCase.ReadALL()
		assert.Error(t, err)
	})
	t.Run("Failed Read All Data -- Internal Error", func(t *testing.T) {
		voucherRepositoryMock.On("ReadALL").Return([]vouchers.Domain{{}},
			errors.New("some random err")).Once()
		_, err := voucherUseCase.ReadALL()
		assert.Error(t, err)
	})
}

func TestReadByID(t *testing.T) {
	Setup()
	t.Run("Success Read Voucher", func(t *testing.T) {
		voucherRepositoryMock.On("ReadById",
			mock.AnythingOfType("int"),
		).Return(voucherDomain, nil).Once()
		resp, err := voucherUseCase.ReadById(1)
		assert.Nil(t, err)
		assert.Equal(t, resp, voucherDomain)
	})
	t.Run("Failed Read Voucher - Not Found", func(t *testing.T) {
		voucherRepositoryMock.On("ReadById",
			mock.AnythingOfType("int"),
		).Return(vouchers.Domain{}, nil).Once()
		resp, err := voucherUseCase.ReadById(1)
		assert.Error(t, err)
		assert.Equal(t, resp, vouchers.Domain{})
	})
	t.Run("Failed Read Voucher - Internal Errors", func(t *testing.T) {
		voucherRepositoryMock.On("ReadById",
			mock.AnythingOfType("int"),
		).Return(vouchers.Domain{}, errors.New("some randoms error")).Once()
		resp, err := voucherUseCase.ReadById(1)
		assert.Error(t, err)
		assert.Equal(t, resp, vouchers.Domain{})
	})
}

func TestCreate(t *testing.T) {
	Setup()
	t.Run("Success Create Voucher", func(t *testing.T) {
		voucherRepositoryMock.On("Create",
			mock.AnythingOfType("vouchers.Domain"),
		).Return(nil).Once()
		err := voucherUseCase.Create(voucherDomain)
		assert.Nil(t, err)
	})

	t.Run("Failed Duplicate Create Voucher", func(t *testing.T) {
		voucherRepositoryMock.On("Create",
			mock.AnythingOfType("vouchers.Domain"),
		).Return(&mysqlErr).Once()
		err := voucherUseCase.Create(voucherDomain)
		assert.Error(t, err)
	})

	t.Run("Failed Create Voucher | Internal", func(t *testing.T) {
		voucherRepositoryMock.On("Create",
			mock.AnythingOfType("vouchers.Domain"),
		).Return(errors.New("some random error")).Once()
		err := voucherUseCase.Create(voucherDomain)
		assert.Error(t, err)
	})
}
