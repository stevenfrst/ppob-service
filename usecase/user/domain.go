package user

type Domain struct {
	ID          uint `gorm:"primarykey"`
	Role        string
	Username    string `gorm:"unique"`
	Password    string
	Email       string `gorm:"unique"`
	PhoneNumber string `gorm:"unique"`
	Pin         int
	Token       string
}

type IUserUsecase interface {
	Login(username, password string) (Domain, error)
	Register(user Domain) (Domain, error)
}

type IUserRepository interface {
	CheckLogin(email, password string) (Domain, error)
	Register(users *Domain) (Domain, error)
}
