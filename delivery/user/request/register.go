package request

import "ppob-service/usecase/user"

type UserRegister struct {
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number"`
}

func (u *UserRegister) ToDomainUser() user.Domain {
	return user.Domain{
		Username:    u.Username,
		Password:    u.Password,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
	}
}