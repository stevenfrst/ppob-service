package response

import (
	vouchers "ppob-service/usecase/voucher"
	"time"
)

type Voucher struct {
	ID    uint      `json:"id"`
	Code  string    `json:"code"`
	Value int       `json:"value"`
	Valid time.Time `json:"valid"`
}

func FromDomain(domain vouchers.Domain) Voucher {
	return Voucher{
		ID:    domain.ID,
		Code:  domain.Code,
		Value: domain.Value,
		Valid: domain.Valid,
	}
}

func FromDomainList(voc []vouchers.Domain) []Voucher {
	var dummyProducts []Voucher
	for x := range voc {
		dummyDomain := FromDomain(voc[x])
		dummyProducts = append(dummyProducts, dummyDomain)
	}
	return dummyProducts
}
