package user

import (
	"errors"
	_middleware "ppob-service/app/middleware"
)

type UseCase struct {
	repo IUserRepository
	jwt  *_middleware.ConfigJWT
}

func NewUseCase(userRepo IUserRepository, configJWT *_middleware.ConfigJWT) IUserUsecase {
	return &UseCase{
		repo: userRepo,
		jwt:  configJWT,
	}
}

func (u *UseCase) Login(username, password string) (Domain, error) {
	user, err := u.repo.CheckLogin(username, password)
	if err != nil {
		return user, errors.New("internal error")
	} else if user.ID == 0 {
		return Domain{}, errors.New("email/password not match")
	}
	token := u.jwt.GenerateToken(int(user.ID), user.Role)
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
	resp , err := u.repo.DetailUser(id)
	if err != nil {
		return Domain{}, err
	}
	return resp,nil
}