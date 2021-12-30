package user

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"ppob-service/helpers/encrypt"
	"ppob-service/usecase/user"
	"strconv"
	"time"
)

type UserRepository struct {
	db    *gorm.DB
	cache redis.Conn
}

func NewRepository(gormDb *gorm.DB, cache redis.Conn) user.IUserRepository {
	return &UserRepository{
		db:    gormDb,
		cache: cache,
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
	//var userRepo User

	email, err := r.GetEmail(uint(id))
	if err != nil {
		return "", err
	}

	resp, err := r.CheckLogin(email, oldPassword)
	log.Println(resp)
	if err != nil {
		return "", err
	}

	userRepo := FromDomain(&resp)

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
		return user.Domain{}, err
	}
	return userRepo.ToDomain(), nil
}

func (r *UserRepository) SavePinToRedis(id int) (string, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	pin := rand.Intn(9999-1000) + 1000
	toRedis := fmt.Sprintf("pin:%v", id)
	_, err := r.cache.Do("SET", toRedis, pin)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(pin), nil
}

func (r *UserRepository) ReadPin(id int) (int, error) {
	pin, err := redis.Int(r.cache.Do("GET", fmt.Sprintf("pin:%v", id)))
	if err != nil {
		return 0, err
	}
	return pin, nil
}

func (r *UserRepository) ChangeStatus(id int) error {
	var repoModel User
	err := r.db.Where("id = ?", id).First(&repoModel).Error
	if err != nil {
		return err
	}

	repoModel.IsVerified = true
	err = r.db.Save(&repoModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) ResetPassword(email, password string) error {
	var repoModel User
	err := r.db.Where("email = ?", email).First(&repoModel).Error
	if err != nil {
		return err
	}
	log.Println(repoModel.ID)
	repoModel.Password, _ = encrypt.Hash(password)
	err = r.db.Save(&repoModel).Error
	if err != nil {
		return err
	}
	log.Println(repoModel.Password)

	return nil
}
