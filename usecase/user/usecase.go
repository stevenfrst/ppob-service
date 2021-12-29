package user

import (
	"errors"
	"fmt"
	"gopkg.in/gomail.v2"
	_middleware "ppob-service/app/middleware"
)

type UseCase struct {
	repo IUserRepository
	jwt  *_middleware.ConfigJWT
	mail gomail.Dialer
}

func NewUseCase(userRepo IUserRepository, configJWT *_middleware.ConfigJWT, mail gomail.Dialer) IUserUsecase {
	return &UseCase{
		repo: userRepo,
		jwt:  configJWT,
		mail: mail,
	}
}

func (u *UseCase) Login(username, password string) (Domain, error) {
	user, err := u.repo.CheckLogin(username, password)
	if user.ID == 0 {
		return Domain{}, errors.New("email/password not match")
	} else if err != nil {
		return user, errors.New("internal error")
	}
	token := u.jwt.GenerateToken(int(user.ID), user.Role, user.IsVerified)
	user.Token = token
	return user, err
}

func (u *UseCase) Register(user Domain) (string, error) {
	user.Role = "user"
	resp, err := u.repo.Register(&user)
	if err != nil {
		return "", errors.New("internal error")
	}
	return resp, err
}

func (u *UseCase) ChangePassword(id int, oldPassword, newPassword string) (string, error) {
	resp, err := u.repo.ChangePassword(id, oldPassword, newPassword)
	if err != nil {
		return "", nil
	} else if resp == "not found" {
		return "User not found", nil
	}
	return resp, nil
}

func (u *UseCase) GetCurrentUser(id int) (Domain, error) {
	resp, err := u.repo.DetailUser(id)
	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}

func (u *UseCase) Verify(id,pin int) error {
	resp, err := u.repo.ReadPin(id)
	if err != nil {
		return err
	}
	if resp == pin {
		err = u.repo.ChangeStatus(id)
		if err != nil {
			return err
		}
	} else {
		return errors.New("not match")
	}
	return nil
}
func (u *UseCase) SendPin(id int) error {
	// Get Email By ID
	email, err := u.repo.GetEmail(uint(id))
	if err != nil {
		return err
	}

	// Save random to Redis
	pin, err := u.repo.SavePinToRedis(id)
	if err != nil {
		return err
	}
	// Send Email
	var mailDomain = EmailDriver{
		Sender:  u.mail.Username,
		ToEmail: email,
		Subject: "Kode Verification Pin",
	}
	err = u.mail.DialAndSend(createHeader(mailDomain, pin))
	if err != nil {
		return err
	}

	return nil
}

func createHeader(s EmailDriver, pin string) *gomail.Message {
	var bodyEmail string

	bodyEmail = fmt.Sprintf("Pin Anda Untuk Konfimasi Email <b>%v<b>", pin)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", s.Sender)
	mailer.SetHeader("To", s.ToEmail)
	mailer.SetHeader("Subject", s.Subject)
	mailer.SetBody("text/html", bodyEmail)

	return mailer
}
