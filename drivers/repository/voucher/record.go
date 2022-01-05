package vouchers

import (
	"gorm.io/gorm"
	vouchers "ppob-service/usecase/voucher"
	"time"
)

type Voucher struct {
	ID        uint   `gorm:"primarykey"`
	Code      string `gorm:"unique"`
	Value     int
	Valid     time.Time
	CreatedAt *time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (v *Voucher) ToDomain() vouchers.Domain {

	return vouchers.Domain{
		ID:    v.ID,
		Code:  v.Code,
		Value: v.Value,
		Valid: v.Valid,
	}
}

func ToDomainList(voc []Voucher) []vouchers.Domain {
	var dummyDomain []vouchers.Domain
	for x := range voc {
		dummyProducts := voc[x].ToDomain()
		dummyDomain = append(dummyDomain, dummyProducts)
	}
	return dummyDomain
}

func FromDomain(voc vouchers.Domain) Voucher {
	return Voucher{
		ID:    voc.ID,
		Code:  voc.Code,
		Value: voc.Value,
		Valid: voc.Valid,
	}
}
