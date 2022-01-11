package transaction

import (
	"fmt"
	"gopkg.in/gomail.v2"
	payment "ppob-service/drivers/midtrans"
	"ppob-service/helpers/errorHelper"
	"strconv"
)

type UseCase struct {
	repo    ITransactionRepository
	payment payment.MidtransInterface
	mail    gomail.Dialer
}

func NewUseCase(repo ITransactionRepository, payment payment.MidtransInterface, mail gomail.Dialer) ITransactionUsecase {
	return &UseCase{
		repo,
		payment,
		mail,
	}
}

func (u *UseCase) GetAllTxUser(id int) ([]HistoryDomain, error) {
	resp, err := u.repo.GetUserTxByID(id)
	if resp[0].ID == 0 {
		return []HistoryDomain{}, errorHelper.ErrRecordNotFound
	} else if err != nil {
		return []HistoryDomain{}, err
	}

	for x := range resp {
		name, tax := u.repo.GetNameNTax(int(resp[x].ProductID))
		resp[x].ProductName = name
		resp[x].Tax = tax
	}

	return resp, nil
}

func (u *UseCase) GetTxByID(id int) (HistoryDomain, error) {
	resp, err := u.repo.GetTxHistoryByID(id)
	if resp.ID == 0 {
		return HistoryDomain{}, errorHelper.ErrRecordNotFound
	} else if err != nil {
		return HistoryDomain{}, err
	}
	name, tax := u.repo.GetNameNTax(int(resp.ProductID))
	resp.ProductName = name
	resp.Tax = tax
	return resp, nil
}

func (u *UseCase) ProcessNotification(input Notification) error {
	txId, _ := strconv.Atoi(input.OrderID)
	tx, err := u.repo.GetTxByID(txId)
	if tx.ID == 0 {
		return errorHelper.ErrRecordNotFound
	} else if err != nil {
		return err
	}

	if input.PaymentType == "bank_transfer" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" || input.TransactionStatus == "settlement" {
		tx.PaymentType = input.PaymentType
		tx.TransactionStatus = input.TransactionStatus
		tx.FraudStatus = input.FraudStatus
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		tx.TransactionStatus = input.TransactionStatus
	}

	email, username := u.repo.GetUserEmail(int(tx.UserID))
	var mailDomain = EmailDriver{
		Sender:  u.mail.Username,
		ToEmail: email,
		Subject: "Transaction Success",
	}
	var bodyEmail string
	bodyEmail = fmt.Sprintf("Hello %v,\nyour recent payment of Rp %v for Invoice ID #%v SUCCESSFULLY VALIDATED. \n\nBest Regards,\nGesek.co", username, tx.Total, txId)
	err = u.mail.DialAndSend(createHeader(mailDomain, bodyEmail))
	if err != nil {
		return err
	}

	err = u.repo.UpdateTx(tx)
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCase) GetVirtualAccount(id int, inVA CreateVA) (string, error) {
	txDetailID, err := u.repo.CreateDetail(DetailDomain{
		ProductID: inVA.ProductID,
		Discount:  inVA.Discount,
		Subtotal:  inVA.Subtotal,
	})
	if err != nil {
		return "", err
	}
	total := inVA.Tax + inVA.Subtotal - inVA.Discount

	resp := u.payment.CreateVirtualAccount(id, total, inVA.Bank)

	err = u.repo.CreateTx(Domain{
		ID:                  uint(resp.ID),
		UserID:              uint(id),
		DetailTransactionID: txDetailID,
		Total:               total,
		Link:                resp.VA,
		TransactionStatus:   "pending",
		FraudStatus:         "accept",
		PaymentType:         "bank transfer",
		Provider:            resp.Provider,
	})
	if err != nil {
		return "", err
	}
	return resp.VA, nil
}

func createHeader(s EmailDriver, header ...string) *gomail.Message {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", s.Sender)
	mailer.SetHeader("To", s.ToEmail)
	mailer.SetHeader("Subject", s.Subject)
	mailer.SetBody("text/html", header[0])

	return mailer
}
