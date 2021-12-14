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
	Register(user Domain) (string, error)
	ChangePassword(id int, oldPassword, newPassword string) (string, error)
	GetCurrentUser(id int) (Domain, error)
}

type IUserRepository interface {
	CheckLogin(email, password string) (Domain, error)
	Register(users *Domain) (string, error)
	ChangePassword(id int, oldPassword, newPassword string) (string, error)
	DetailUser(id int) (Domain, error)
}
