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
