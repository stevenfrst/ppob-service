package transaction

import (
	"gorm.io/gorm"
	"log"
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

func (t *TransactionRepository) CreateDetail(det transaction.DetailDomain) (uint, error) {
	var repoModel = FromDomainDetailTransaction(det)
	err := t.db.Create(&repoModel).Error
	if err != nil {
		return 0, err
	}
	return repoModel.ID, nil
}

func (t *TransactionRepository) CreateTx(tx transaction.Domain) error {
	var repoModel = FromDomainTransaction(tx)
	err := t.db.Create(&repoModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepository) GetTxByID(id int) (transaction.Domain, error) {
	var repoModel Transaction
	err := t.db.Where("id = ?", id).First(&repoModel).Error
	if err != nil {
		return transaction.Domain{}, err
	}
	return repoModel.ToDomain(), nil
}

func (t *TransactionRepository) GetUserTxByID(id int) ([]transaction.HistoryDomain, error) {
	var repoModel []Transaction
	err := t.db.Preload("DetailTransaction").Where("user_id = ?", id).Find(&repoModel).Error
	if err != nil {
		return ToHistoryDomainList([]Transaction{}), err
	}
	return ToHistoryDomainList(repoModel), nil
}

func (t *TransactionRepository) GetTxHistoryByID(id int) (transaction.HistoryDomain, error) {
	var repoModel Transaction
	err := t.db.Preload("DetailTransaction").Where("id = ?", id).First(&repoModel).Error
	if err != nil {
		return repoModel.ToHistoryDomain(), err
	}
	return repoModel.ToHistoryDomain(), nil
}

func (t *TransactionRepository) UpdateTx(tx transaction.Domain) error {
	log.Println(tx)
	return t.db.Save(FromDomainTransaction(tx)).Error
}

func (t *TransactionRepository) GetUserEmail(id int) (string, string) {
	var repoModel User
	t.db.Where("id = ?", id).First(&repoModel)
	return repoModel.Email, repoModel.Username
}

func (t *TransactionRepository) GetNameNTax(id int) (string, int) {
	var repoModel Product
	t.db.Preload("SubCategory").Where("id = ?", id).First(&repoModel)
	return repoModel.Name, repoModel.SubCategory.Tax
}
