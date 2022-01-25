package vouchers

import "time"

type Domain struct {
	ID    uint `gorm:"primarykey"`
	Code  string
	Value int
	Valid time.Time
}

type IVoucherRepository interface {
	Create(voc Domain) error
	ReadById(id int) (Domain, error)
	ReadALL() ([]Domain, error)
	DeleteByID(id int) error
	Verify(code string) (int,error)
}

type IVoucherUseCase interface {
	Create(voc Domain) error
	ReadById(id int) (Domain, error)
	ReadALL() ([]Domain, error)
	DeleteByID(id int) error
	Verify(code string) (int,error)
}
