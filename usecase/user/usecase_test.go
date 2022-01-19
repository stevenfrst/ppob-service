package user_test

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
	"ppob-service/app/config"
	"ppob-service/app/middleware"
	"ppob-service/drivers/email"
	"ppob-service/usecase/user"
	"ppob-service/usecase/user/mocks"
	"testing"
)

var userMockRepo mocks.IUserRepository
var userUseCase user.IUserUsecase
var userDomain user.Domain
var dialer *gomail.Dialer
var emailDummy user.EmailDriver

func Setup() {
	getConfig := config.GetConfigTest()
	configJWT := middleware.ConfigJWT{
		SecretJWT:       getConfig.JWT_SECRET,
		ExpiresDuration: int64(getConfig.JWT_EXPIRED),
	}
	gmail := email.SmtpConfig{
		CONFIG_SMTP_HOST:       getConfig.CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT:       getConfig.CONFIG_SMTP_PORT,
		CONFIG_SMTP_AUTH_EMAIL: getConfig.CONFIG_SMTP_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD:   getConfig.CONFIG_AUTH_PASSWORD,
		CONFIG_SENDER_NAME:     getConfig.CONFIG_SENDER_NAME,
	}
	dialer = email.NewGmailConfig(gmail)
	userUseCase = user.NewUseCase(&userMockRepo, &configJWT, *dialer)
	userDomain = user.Domain{
		ID:          1,
		Role:        "user",
		Username:    "ponta",
		Password:    "Jambu123",
		Email:       "ponta@mail.com",
		PhoneNumber: "082135166117",
		IsVerified:  false,
		Token:       "",
	}
	emailDummy = user.EmailDriver{
		Sender:  "oppaidaisuki363@gmail.com",
		ToEmail: "test@mail.com",
		Subject: "test",
	}
}

func TestResetPassword(t *testing.T) {
	Setup()
	t.Run("Success ChangePassword", func(t *testing.T) {
		userMockRepo.On("ResetPassword",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(nil).Once()

		err := userUseCase.ResetPassword("kafka@mail.com")
		assert.Nil(t, err)
	})

	t.Run("Failed ChangePassword", func(t *testing.T) {
		userMockRepo.On("ResetPassword",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(errors.New("db errors")).Once()

		err := userUseCase.ResetPassword("kafka@mail.com")
		assert.Error(t, err)
	})
}

func TestSendPing(t *testing.T) {
	Setup()
	t.Run("Success SendPin", func(t *testing.T) {
		userMockRepo.On("GetEmail",
			mock.AnythingOfType("uint"),
		).Return("kafka@mail.com", nil).Once()

		userMockRepo.On("SavePinToRedis",
			mock.AnythingOfType("int"),
		).Return("1234", nil).Once()

		err := userUseCase.SendPin(1)
		assert.Nil(t, err)
	})

	t.Run("Failed Get Email", func(t *testing.T) {
		userMockRepo.On("GetEmail",
			mock.AnythingOfType("uint"),
		).Return("", errors.New("some random db errors")).Once()

		err := userUseCase.SendPin(1)
		assert.Error(t, err)
	})

	t.Run("error saving pin to redis", func(t *testing.T) {
		userMockRepo.On("GetEmail",
			mock.AnythingOfType("uint"),
		).Return("kafka@mail.com", nil).Once()

		userMockRepo.On("SavePinToRedis",
			mock.AnythingOfType("int"),
		).Return("", errors.New("redis error")).Once()

		err := userUseCase.SendPin(1)
		assert.Error(t, err)
	})

}

func TestVerify(t *testing.T) {
	Setup()
	t.Run("Verify Success", func(t *testing.T) {
		userMockRepo.On("ReadPin",
			mock.AnythingOfType("int"),
		).Return(1234, nil).Once()

		userMockRepo.On("ChangeStatus",
			mock.AnythingOfType("int"),
		).Return(nil).Once()

		err := userUseCase.Verify(1, 1234)
		assert.Nil(t, err)
	})

	t.Run("Redis Error/NotFound", func(t *testing.T) {
		userMockRepo.On("ReadPin",
			mock.AnythingOfType("int"),
		).Return(0, errors.New("some random errors")).Once()

		err := userUseCase.Verify(1, 1234)
		assert.Error(t, err)
	})

	t.Run("Mysql Error", func(t *testing.T) {
		userMockRepo.On("ReadPin",
			mock.AnythingOfType("int"),
		).Return(1234, nil).Once()

		userMockRepo.On("ChangeStatus",
			mock.AnythingOfType("int"),
		).Return(errors.New("some random database errors")).Once()

		err := userUseCase.Verify(1, 1234)
		assert.Error(t, err)
	})

	t.Run("Pin Not Match", func(t *testing.T) {
		userMockRepo.On("ReadPin",
			mock.AnythingOfType("int"),
		).Return(1234, nil).Once()

		err := userUseCase.Verify(1, 12344)
		assert.Error(t, err)
	})
}

func TestChangePassword(t *testing.T) {
	Setup()
	t.Run("Success Change Passsword", func(t *testing.T) {
		userMockRepo.On("ChangePassword",
			mock.AnythingOfType("int"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return("success", nil).Once()
		resp, err := userUseCase.ChangePassword(1, "haha123", "haha123")
		assert.Nil(t, err)
		assert.Equal(t, "success", resp)
	})
	t.Run("failed Password Mismatch", func(t *testing.T) {
		userMockRepo.On("ChangePassword",
			mock.AnythingOfType("int"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return("", bcrypt.ErrMismatchedHashAndPassword).Once()
		resp, err := userUseCase.ChangePassword(1, "haha123", "hahaa")
		assert.Error(t, err)
		assert.Equal(t, "", resp)
	})
	t.Run("failed internal error", func(t *testing.T) {
		userMockRepo.On("ChangePassword",
			mock.AnythingOfType("int"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return("", errors.New("some db error")).Once()
		resp, err := userUseCase.ChangePassword(1, "hahaa", "hahaa")
		assert.Error(t, err)
		assert.Equal(t, "", resp)
	})
	t.Run("failed user not found", func(t *testing.T) {
		userMockRepo.On("ChangePassword",
			mock.AnythingOfType("int"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return("not found", nil).Once()
		resp, err := userUseCase.ChangePassword(1, "", "hahaa")
		assert.Nil(t, err)
		assert.Equal(t, "User not found", resp)
	})
}

func TestGetCurrentUser(t *testing.T) {
	Setup()
	t.Run("Success Get User", func(t *testing.T) {
		userMockRepo.On("DetailUser",
			mock.AnythingOfType("int"),
		).Return(userDomain, nil).Once()
		resp, err := userUseCase.GetCurrentUser(1)
		assert.Nil(t, err)
		assert.Equal(t, resp, userDomain)
	})
	t.Run("Internal err", func(t *testing.T) {
		userMockRepo.On("DetailUser",
			mock.AnythingOfType("int"),
		).Return(user.Domain{}, errors.New("some random err")).Once()
		resp, err := userUseCase.GetCurrentUser(1)
		assert.Error(t, err)
		assert.Equal(t, resp, user.Domain{})
	})
}

func TestLogin(t *testing.T) {
	Setup()
	t.Run("Success Login", func(t *testing.T) {
		userMockRepo.On("CheckLogin",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(userDomain, nil).Once()
		resp, err := userUseCase.Login("ponta", "Jambu123")
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("user Not Found", func(t *testing.T) {
		userMockRepo.On("CheckLogin",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(user.Domain{}, gorm.ErrRecordNotFound).Once()
		resp, err := userUseCase.Login("ponta", "Jambu123")
		assert.Error(t, err)
		assert.Equal(t, resp, user.Domain{})
	})
	t.Run("password not match", func(t *testing.T) {
		userMockRepo.On("CheckLogin",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(user.Domain{}, nil).Once()
		resp, err := userUseCase.Login("ponta", "Jambu123")
		assert.Error(t, err)
		assert.Equal(t, resp, user.Domain{})
	})
	t.Run("internal error", func(t *testing.T) {
		userMockRepo.On("CheckLogin",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(user.Domain{}, errors.New("unknown err")).Once()
		resp, err := userUseCase.Login("ponta", "Jambu123")
		assert.Error(t, err)
		assert.Equal(t, resp, user.Domain{})
	})
}

func TestRegister(t *testing.T) {
	Setup()
	t.Run("Success Login", func(t *testing.T) {
		userMockRepo.On("Register",
			mock.AnythingOfType("*user.Domain")).Return("success", nil).Once()
		resp, err := userUseCase.Register(userDomain)
		assert.Nil(t, err)
		assert.Equal(t, "success", resp)
	})
	t.Run("Internal error", func(t *testing.T) {
		userMockRepo.On("Register",
			mock.AnythingOfType("*user.Domain")).Return("", errors.New("some unknown error")).Once()
		resp, err := userUseCase.Register(userDomain)
		assert.Error(t, err)
		assert.Equal(t, "", resp)
	})
	t.Run("error duplicate", func(t *testing.T) {
		var mysqlErr mysql.MySQLError
		mysqlErr.Number = 1062
		userMockRepo.On("Register",
			mock.AnythingOfType("*user.Domain")).Return("", &mysqlErr).Once()
		resp, err := userUseCase.Register(userDomain)
		assert.Error(t, err)
		assert.Equal(t, "", resp)
	})
}
