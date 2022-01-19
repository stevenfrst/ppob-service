package transaction_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/gomail.v2"
	"ppob-service/app/config"
	"ppob-service/drivers/email"
	payment "ppob-service/drivers/midtrans"
	_mockPayment "ppob-service/drivers/midtrans/mocks"
	"ppob-service/usecase/transaction"
	"ppob-service/usecase/transaction/mocks"
	"testing"
	"time"
)

var txMockrepo mocks.ITransactionRepository
var txUseCase transaction.ITransactionUsecase
var txDomainMock transaction.Domain
var txNotifMockSuccess transaction.Notification
var txNotifMockFailed transaction.Notification
var historyDomainMock transaction.HistoryDomain
var historyDomainMocks []transaction.HistoryDomain
var paymentMock _mockPayment.MidtransInterface
var dialer *gomail.Dialer
var mockCreateVA transaction.CreateVA
var emailDummy transaction.EmailDriver
var coreApiMock payment.CoreAPIResponse

func Setup() {
	getConfig := config.GetConfigTest()
	configPayment := payment.ConfigMidtrans{
		SERVER_KEY: getConfig.SERVER_KEY,
	}
	emailDummy = transaction.EmailDriver{
		Sender:  "oppaidaisuki363@gmail.com",
		ToEmail: "test@mail.com",
		Subject: "test",
	}
	gmail := email.SmtpConfig{
		CONFIG_SMTP_HOST:       getConfig.CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT:       getConfig.CONFIG_SMTP_PORT,
		CONFIG_SMTP_AUTH_EMAIL: getConfig.CONFIG_SMTP_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD:   getConfig.CONFIG_AUTH_PASSWORD,
		CONFIG_SENDER_NAME:     getConfig.CONFIG_SENDER_NAME,
	}
	dialer = email.NewGmailConfig(gmail)
	configPayment.SetupGlobalMidtransConfig()
	txUseCase = transaction.NewUseCase(&txMockrepo, &paymentMock, *dialer)
	historyDomainMock = transaction.HistoryDomain{
		ID:                1,
		ProductID:         1,
		ProductName:       "tahu",
		Discount:          0,
		Tax:               1000,
		Total:             10000,
		Link:              "123123123",
		TransactionStatus: "settlement",
		FraudStatus:       "accept",
		PaymentType:       "bank_transfer",
		Provider:          "bni",
		CreatedAt:         time.Time{},
	}
	historyDomainMocks = append(historyDomainMocks, historyDomainMock)
	txDomainMock = transaction.Domain{
		ID:                  1,
		UserID:              1,
		DetailTransactionID: 1,
		Total:               1,
		Link:                "123123123",
		TransactionStatus:   "settlement",
		FraudStatus:         "accept",
		PaymentType:         "bank_transfer",
		Provider:            "bni",
	}
	txNotifMockSuccess = transaction.Notification{
		TransactionStatus: "settlement",
		OrderID:           "123456789",
		PaymentType:       "bank_transfer",
		FraudStatus:       "accept",
	}
	txNotifMockFailed = transaction.Notification{
		TransactionStatus: "expire",
		OrderID:           "123456789",
		PaymentType:       "bank_transfer",
		FraudStatus:       "accept",
	}
	mockCreateVA = transaction.CreateVA{
		OrderID:   1213123123,
		ProductID: 1,
		Discount:  0,
		Tax:       1000,
		Subtotal:  10000,
		Bank:      "bni",
	}
	coreApiMock = payment.CoreAPIResponse{
		ID:       1,
		VA:       "12345678",
		Provider: "test",
	}
}

func TestGetVirtualAccount(t *testing.T) {
	Setup()
	t.Run("failed create detail transaction", func(t *testing.T) {
		txMockrepo.On("CreateDetail",
			mock.AnythingOfType("transaction.DetailDomain"),
		).Return(uint(1), errors.New("internal error")).Once()
		_, err := txUseCase.GetVirtualAccount(1, mockCreateVA)
		assert.Error(t, err)
	})
	t.Run("failed create transaction", func(t *testing.T) {
		txMockrepo.On("CreateDetail",
			mock.AnythingOfType("transaction.DetailDomain"),
		).Return(uint(1), nil).Once()

		paymentMock.On("CreateVirtualAccount",
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int"),
			mock.AnythingOfType("string"),
		).Return(coreApiMock).Once()

		txMockrepo.On("CreateTx",
			mock.AnythingOfType("transaction.Domain"),
		).Return(errors.New("some random error")).Once()

		_, err := txUseCase.GetVirtualAccount(1, mockCreateVA)
		assert.Error(t, err)
	})
	t.Run("success create transaction", func(t *testing.T) {
		txMockrepo.On("CreateDetail",
			mock.AnythingOfType("transaction.DetailDomain"),
		).Return(uint(1), nil).Once()

		paymentMock.On("CreateVirtualAccount",
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int"),
			mock.AnythingOfType("string"),
		).Return(coreApiMock).Once()

		txMockrepo.On("CreateTx",
			mock.AnythingOfType("transaction.Domain"),
		).Return(nil).Once()

		_, err := txUseCase.GetVirtualAccount(1, mockCreateVA)
		assert.Nil(t, err)
	})

}

func TestProcessNotification(t *testing.T) {
	Setup()
	t.Run("Failed not found transaction", func(t *testing.T) {
		txMockrepo.On("GetTxByID",
			mock.AnythingOfType("int"),
		).Return(transaction.Domain{}, nil).Once()
		err := txUseCase.ProcessNotification(txNotifMockSuccess)
		assert.Error(t, err)
	})
	t.Run("Failed found transaction - internal", func(t *testing.T) {
		txMockrepo.On("GetTxByID",
			mock.AnythingOfType("int"),
		).Return(transaction.Domain{}, errors.New("some error")).Once()
		err := txUseCase.ProcessNotification(txNotifMockSuccess)
		assert.Error(t, err)
	})

	t.Run("transaction expire", func(t *testing.T) {
		txMockrepo.On("GetTxByID",
			mock.AnythingOfType("int"),
		).Return(txDomainMock, nil).Once()

		txMockrepo.On("UpdateTx",
			mock.AnythingOfType("transaction.Domain"),
		).Return(nil).Once()

		err := txUseCase.ProcessNotification(txNotifMockFailed)
		assert.Nil(t, err)
	})

	t.Run("transaction expire - error update", func(t *testing.T) {
		txMockrepo.On("GetTxByID",
			mock.AnythingOfType("int"),
		).Return(txDomainMock, nil).Once()

		txMockrepo.On("UpdateTx",
			mock.AnythingOfType("transaction.Domain"),
		).Return(errors.New("some db errors")).Once()

		err := txUseCase.ProcessNotification(txNotifMockFailed)
		assert.Error(t, err)
	})

	t.Run("transaction success - error update", func(t *testing.T) {
		txMockrepo.On("GetTxByID",
			mock.AnythingOfType("int"),
		).Return(txDomainMock, nil).Once()

		txMockrepo.On("UpdateTx",
			mock.AnythingOfType("transaction.Domain"),
		).Return(errors.New("some db errors")).Once()

		err := txUseCase.ProcessNotification(txNotifMockSuccess)
		assert.Error(t, err)
	})
	t.Run("transaction success", func(t *testing.T) {
		txMockrepo.On("GetTxByID",
			mock.AnythingOfType("int"),
		).Return(txDomainMock, nil).Once()

		txMockrepo.On("UpdateTx",
			mock.AnythingOfType("transaction.Domain"),
		).Return(nil).Once()

		txMockrepo.On("UpdateStocks",
			mock.AnythingOfType("int"),
		).Once()

		txMockrepo.On("GetUserEmail",
			mock.AnythingOfType("int"),
		).Return("test@mail.com", "test").Once()

		err := txUseCase.ProcessNotification(txNotifMockSuccess)
		assert.Nil(t, err)
	})
}

func TestGetTxID(t *testing.T) {
	Setup()
	t.Run("success get all data", func(t *testing.T) {
		txMockrepo.On("GetTxHistoryByID",
			mock.AnythingOfType("int"),
		).Return(historyDomainMock, nil).Once()

		txMockrepo.On("GetNameNTax",
			mock.AnythingOfType("int"),
		).Return("tahu", 1000).Once()

		resp, err := txUseCase.GetTxByID(1)
		assert.NoError(t, err)
		assert.Equal(t, resp, historyDomainMock)
	})
	t.Run("failed get all data - Internal error", func(t *testing.T) {
		txMockrepo.On("GetTxHistoryByID",
			mock.AnythingOfType("int"),
		).Return(transaction.HistoryDomain{}, errors.New("some db's error")).Once()

		resp, err := txUseCase.GetTxByID(1)
		assert.Error(t, err)
		assert.Equal(t, resp, transaction.HistoryDomain{})
	})
	t.Run("failed get all data - not found", func(t *testing.T) {
		txMockrepo.On("GetTxHistoryByID",
			mock.AnythingOfType("int"),
		).Return(transaction.HistoryDomain{}, nil).Once()

		resp, err := txUseCase.GetTxByID(1)
		assert.Error(t, err)
		assert.Equal(t, resp, transaction.HistoryDomain{})
	})
}

func TestGetAllTxUsername(t *testing.T) {
	Setup()
	t.Run("success get all data", func(t *testing.T) {
		txMockrepo.On("GetUserTxByID",
			mock.AnythingOfType("int"),
		).Return(historyDomainMocks, nil).Once()

		txMockrepo.On("GetNameNTax",
			mock.AnythingOfType("int"),
		).Return("test", 1000)

		resp, err := txUseCase.GetAllTxUser(2)
		assert.NoError(t, err)
		assert.Equal(t, resp, historyDomainMocks)
	})

	t.Run("failed nil response", func(t *testing.T) {
		txMockrepo.On("GetUserTxByID",
			mock.AnythingOfType("int"),
		).Return([]transaction.HistoryDomain{{}}, nil).Once()

		_, err := txUseCase.GetAllTxUser(1)
		assert.Error(t, err)
	})

	t.Run("failed internal error", func(t *testing.T) {
		txMockrepo.On("GetUserTxByID",
			mock.AnythingOfType("int"),
		).Return([]transaction.HistoryDomain{{}}, errors.New("some random err")).Once()

		_, err := txUseCase.GetAllTxUser(1)
		assert.Error(t, err)
	})
}
