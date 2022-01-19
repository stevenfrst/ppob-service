package vouchers

import (
	"gorm.io/gorm"
	"ppob-service/helpers/errorHelper"
	vouchers "ppob-service/usecase/voucher"
)

type VoucherRepository struct {
	db *gorm.DB
}

func NewRepository(gormDb *gorm.DB) vouchers.IVoucherRepository {
	return &VoucherRepository{
		db: gormDb,
	}
}

func (r *VoucherRepository) Create(voc vouchers.Domain) error {
	repoModel := FromDomain(voc)
	err := r.db.Create(&repoModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *VoucherRepository) ReadById(id int) (vouchers.Domain, error) {
	var repoModel Voucher
	err := r.db.Where("id = ?", id).First(&repoModel).Error
	if err != nil {
		return vouchers.Domain{}, err
	}
	return repoModel.ToDomain(), nil
}

func (r *VoucherRepository) ReadALL() ([]vouchers.Domain, error) {
	var repoModel []Voucher
	err := r.db.Find(&repoModel).Error
	if err != nil {
		return ToDomainList([]Voucher{}), err
	}
	return ToDomainList(repoModel), nil
}

func (r *VoucherRepository) DeleteByID(id int) error {
	var repoModel Voucher
	err := r.db.Where("id = ?", id).Delete(&repoModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *VoucherRepository) Verify(code string) (int,error) {
	var repoModel Voucher
	err := r.db.Where("code = ?", code).First(&repoModel).Error
	if err != nil {
		return 0,err
	} else if repoModel.ID == 0 {
		return 0,errorHelper.ErrVoucherNotMatch
	}
	return repoModel.Value,nil
}
