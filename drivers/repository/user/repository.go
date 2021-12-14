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

func (r *UserRepository) Register(users *user.Domain) (user.Domain, error) {
	userIn := FromDomain(users)
	hashedPassword, err := encrypt.Hash(users.Password)
	if err != nil {
		return user.Domain{}, err
	}
	userIn.Password = hashedPassword
	err = r.db.Create(userIn).Error
	if err != nil {
		return userIn.ToDomain(), err
	}
	return userIn.ToDomain(), nil

}
