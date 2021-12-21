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

func (t *TransactionRepository) Create(input transaction.Domain) error {
	repoModel := FromDomainTransaction(input)
	err := t.db.Create(&repoModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepository) GetByID(ID int) (transaction.Domain, error) {
	var repoModel Transaction
	err := t.db.Where("id = ?", ID).Find(&repoModel).Error
	if err != nil {
		return transaction.Domain{}, err
	}
	return repoModel.ToDomain(), nil
}

func (t *TransactionRepository) Update(input transaction.Domain) error {
	repoModel := FromDomainTransaction(input)
	err := t.db.Save(repoModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepository) GetByUserID(id int) ([]transaction.Domain, error) {
	var repoModel []Transaction
	err := t.db.Preload("User").Where("user_id = ?", id).Find(&repoModel).Error
	if err != nil {
		return ToDomainList([]Transaction{}), err
	}
	return ToDomainList(repoModel), nil
}
