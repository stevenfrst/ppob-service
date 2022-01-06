package vouchers

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"ppob-service/helpers/errorHelper"
)

type VoucherUsecase struct {
	repo IVoucherRepository
}

func NewUseCase(repo IVoucherRepository) IVoucherUseCase {
	return &VoucherUsecase{
		repo: repo,
	}
}

func (u *VoucherUsecase) Create(voc Domain) error {
	err := u.repo.Create(voc)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return errorHelper.DuplicateData
	} else if err != nil {
		return err
	}
	return nil
}

func (u *VoucherUsecase) ReadById(id int) (Domain, error) {
	resp, err := u.repo.ReadById(id)
	if resp.ID == 0 {
		return Domain{}, errorHelper.ErrRecordNotFound
	} else if err != nil {
		return Domain{}, err
	}
	return resp, nil
}

func (u *VoucherUsecase) ReadALL() ([]Domain, error) {
	resp, err := u.repo.ReadALL()
	if resp[0].ID == 0 {
		return []Domain{}, errorHelper.ErrRecordNotFound
	} else if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (u *VoucherUsecase) DeleteByID(id int) error {
	return u.repo.DeleteByID(id)
}

func (u *VoucherUsecase) Verify(code string) error {
	return u.repo.Verify(code)
}
