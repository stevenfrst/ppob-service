package user

type Domain struct {
	ID          uint `gorm:"primarykey"`
	Role        string
	Username    string `gorm:"unique"`
	Password    string
	Email       string `gorm:"unique"`
	PhoneNumber string `gorm:"unique"`
	IsVerified  bool
	Token       string
}

type EmailDriver struct {
	Sender  string
	ToEmail string
	Subject string
}

type IUserUsecase interface {
	Login(username, password string) (Domain, error)
	Register(user Domain) (string, error)
	ChangePassword(id int, oldPassword, newPassword string) (string, error)
	GetCurrentUser(id int) (Domain, error)
	SendPin(id int) error
	Verify(id, pin int) error
	ResetPassword(email string) error
}

type IUserRepository interface {
	CheckLogin(email, password string) (Domain, error)
	Register(users *Domain) (string, error)
	ChangePassword(id int, oldPassword, newPassword string) (string, error)
	DetailUser(id int) (Domain, error)
	GetEmail(id uint) (string, error)
	SavePinToRedis(id int) (string, error)
	ReadPin(id int) (int, error)
	ChangeStatus(id int) error
	ResetPassword(email, password string) error
}
