package transaction

import (
	payment "ppob-service/drivers/midtrans"
)

type UseCase struct {
	repo    ITransactionRepository
	payment payment.MidtransInterface
}

func NewUseCase(repo ITransactionRepository, payment payment.MidtransInterface) ITransactionUsecase {
	return &UseCase{
		repo,
		payment,
	}
}

func (t *UseCase) CreateTransaction(productID, userID, nominal int, bank string) (string, error) {
	var transaction Domain
	resp := t.payment.CreateVirtualAccount(userID, nominal, bank)

	transaction.ID = uint(resp.ID)
	transaction.UserID = uint(userID)
	transaction.Total = nominal
	transaction.TransactionStatus = "pending"
	transaction.FraudStatus = "accept"
	transaction.PaymentType = resp.Provider

	err := t.repo.Create(transaction)
	if err != nil {
		return "", err
	}
	return transaction.Link, nil
}
