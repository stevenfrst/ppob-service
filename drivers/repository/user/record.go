package user

import (
	"gorm.io/gorm"
	"ppob-service/drivers/repository/transaction"
	"ppob-service/usecase/user"
	"time"
)

type User struct {
	ID           uint `gorm:"primarykey"`
	Role         string
	Username     string
	Password     string
	Email        string
	PhoneNumber  string
	IsVerified   bool
	Transactions []transaction.Transaction
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (u *User) ToDomain() user.Domain {
	return user.Domain{
		ID:          u.ID,
		Role:        u.Role,
		Username:    u.Username,
		Password:    u.Password,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		IsVerified:  u.IsVerified,
	}
}

func FromDomain(domain *user.Domain) *User {
	return &User{
		ID:          domain.ID,
		Role:        domain.Role,
		Username:    domain.Username,
		Password:    domain.Password,
		Email:       domain.Email,
		PhoneNumber: domain.PhoneNumber,
		IsVerified:  domain.IsVerified,
	}
}
