package vouchers

import (
	"gorm.io/gorm"
	"time"
)

type Voucher struct {
	ID        uint `gorm:"primarykey"`
	Code      string
	Value     int
	Valid     time.Time
	CreatedAt *time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
