package response

import "ppob-service/usecase/user"

type User struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Role        string `json:"role"`
	IsVerified bool `json:"is_verified"`
}

type UserLogin struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

func FromDomainUser(domain user.Domain) UserLogin {
	return UserLogin{
		Token: domain.Token,
		Role:  domain.Role,
	}
}


func FromDomain(domain user.Domain) User {
	return User{
		ID:          domain.ID,
		Username:    domain.Username,
		Email:       domain.Email,
		PhoneNumber: domain.PhoneNumber,
		Role:        domain.Role,
		IsVerified: domain.IsVerified,
	}
}