package transaction

import (
	"gorm.io/gorm"
	"ppob-service/usecase/transaction"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) transaction.ITransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}
