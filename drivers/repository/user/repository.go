package user

import (
	"gorm.io/gorm"
	"ppob-service/helpers/encrypt"
	"ppob-service/usecase/user"
)

type UserRepository struct {
	db *gorm.DB
}

func NewRepository(gormDb *gorm.DB) user.IUserRepository {
	return &UserRepository{
		db: gormDb,
	}
}

func (r *UserRepository) GetEmail(id uint) (string, error) {
	var userRepo User
	err := r.db.Where("id = ?", id).First(&userRepo).Error
	if err != nil {
		return "", err
	}
	return userRepo.Email, nil
}

func (r *UserRepository) ChangePassword(id int, oldPassword, newPassword string) (string, error) {
	var userRepo User

	email, err := r.GetEmail(uint(id))
	if err != nil {
		return "", err
	}

	_, err = r.CheckLogin(email, oldPassword)
	if err != nil {
		return "", err
	}

	hashedPassword, _ := encrypt.Hash(newPassword)
	userRepo.Password = hashedPassword
	err = r.db.Save(&userRepo).Error
	if err != nil {
		return "", err
	}
	return "success", nil
}

func (r *UserRepository) CheckLogin(email, password string) (user.Domain, error) {
	var userRepo User

	err := r.db.Where("email = ?", email).First(&userRepo).Error
	if err != nil {
		return user.Domain{}, err
	}
	err = encrypt.CheckPassword(password, userRepo.Password)
	if err != nil {
		return user.Domain{}, nil
	}
	return userRepo.ToDomain(), nil
}

func (r *UserRepository) Register(users *user.Domain) (string, error) {
	userIn := FromDomain(users)
	hashedPassword, err := encrypt.Hash(users.Password)
	if err != nil {
		return "", err
	}
	userIn.Password = hashedPassword
	err = r.db.Create(userIn).Error
	if err != nil {
		return "", err
	}
	return "success", nil
}

func (r *UserRepository) DetailUser(id int) (user.Domain, error) {
	var userRepo User
	err := r.db.Where("id = ? ", id).First(&userRepo).Error
	if err != nil {
		return user.Domain{},err
	}
	return userRepo.ToDomain(),nil
}
