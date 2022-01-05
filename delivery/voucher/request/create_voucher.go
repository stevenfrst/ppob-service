package request

import (
	vouchers "ppob-service/usecase/voucher"
	"time"
)

type Voucher struct {
	Code  string    `json:"code"`
	Value int       `json:"value"`
	Valid time.Time `json:"valid"`
}

func (v *Voucher) ToDomain() vouchers.Domain {
	return vouchers.Domain{
		Code:  v.Code,
		Value: v.Value,
		Valid: v.Valid,
	}
}
