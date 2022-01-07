package payment

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"strconv"
	"time"
)

type ConfigMidtrans struct {
	SERVER_KEY string
}

type MidtransInterface interface {
	SetupGlobalMidtransConfig()
	CancelPayment(orderID string)
	CreateVirtualAccount(userid, nominal int, bank string) CoreAPIResponse
}

type CoreAPIResponse struct {
	ID       int
	VA       string
	Provider string
}

var c coreapi.Client

func InitializeSnapClient() {
	c.New(midtrans.ServerKey, midtrans.Sandbox)
}

func (p *ConfigMidtrans) SetupGlobalMidtransConfig() {
	midtrans.ServerKey = p.SERVER_KEY
	midtrans.Environment = midtrans.Sandbox
}

func (p *ConfigMidtrans) CancelPayment(orderID string) {
	c.ExpireTransaction(orderID)
}

func (p *ConfigMidtrans) CreateVirtualAccount(userid, nominal int, bank string) CoreAPIResponse {
	id := strconv.Itoa(userid)
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  id + generateOrderIdSuffix(),
			GrossAmt: int64(nominal),
		},

		BankTransfer: &coreapi.BankTransferDetails{
			Bank:     midtrans.Bank(bank),
			VaNumber: "",
			Permata:  nil,
			FreeText: nil,
			Bca:      nil,
		},
	}
	res, _ := c.ChargeTransaction(chargeReq)
	orderID, _ := strconv.Atoi(res.OrderID)
	if bank == "permata" {
		return CoreAPIResponse{
			orderID,
			res.PermataVaNumber,
			"permata",
		}

	}
	getVaNum := res.VaNumbers[0].VANumber
	return CoreAPIResponse{
		orderID,
		getVaNum,
		res.VaNumbers[0].Bank,
	}
}

func generateOrderIdSuffix() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
