package delivery

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"ppob-service/app/middleware"
	"ppob-service/delivery"
	"ppob-service/delivery/user/request"
	"ppob-service/delivery/user/response"
	"ppob-service/usecase/user"
)

type UserDelivery struct {
	usecase user.IUserUsecase
}

func NewUserDelivery(uc user.IUserUsecase) *UserDelivery {
	return &UserDelivery{
		usecase: uc,
	}
}

// Register godoc
// @Summary Register User to server
// @Description register user to server with json.
// @Tags User
// @Accept json
// @Produce json
// @Param  user body request.UserRegister true "User Data"
// @Success 200 {object} delivery.JSONSuccessResult{}
// @Success 400 {object} delivery.JSONBadReqResult{}
// @Success 500 {object} delivery.JSONInternalResult{}
// @Router /v1/register [post]
func (d *UserDelivery) Register(c echo.Context) (err error) {
	var user request.UserRegister
	if err = c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(&user); err != nil {
		return err
	}
	out, err := d.usecase.Register(user.ToDomainUser())
	if err != nil {
		//return delivery.ErrorResponse(c,http.StatusInternalServerError,errorHelper.ERROR_USER_REGISTER,err)
		if fmt.Sprintf("%v", err) == "failed to registering user" {
			return delivery.ErrorResponse(c, http.StatusBadRequest, "error", err)
		} else {
			return delivery.ErrorResponse(c, http.StatusInternalServerError, "error", err)
		}
	}
	return delivery.SuccessResponse(c, out)
}

// Login godoc
// @Summary Login User to server
// @Description login user to server with json.
// @Tags User
// @Accept mpfd
// @Produce json
// @Param email formData string true "email" default(admin)
// @Param password formData string true "password" default(admin)
// @Success 200 {object} delivery.JSONSuccessResult{}
// @Success 400 {object} delivery.JSONBadReqResult{}
// @Success 500 {object} delivery.JSONInternalResult{}
// @Router /v1/login [post]
func (d *UserDelivery) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	var user request.UserLogin
	user.Email = email
	user.Password = password
	if err := c.Validate(&user); err != nil {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "Format Email/Password Salah", err)
	}
	res, err := d.usecase.Login(email, password)
	if err != nil {
		if fmt.Sprintf("%v", err) == "email/password not match" {
			return delivery.ErrorResponse(c, http.StatusUnauthorized, "error", err)
		} else {
			return delivery.ErrorResponse(c, http.StatusInternalServerError, "Internal Error", err)
		}
	}
	return delivery.SuccessResponse(c, response.FromDomainUser(res))
}

// ChangePassword godoc
// @Summary Change Password
// @Description Change Password user to server with json.
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param  user body request.PasswordUpdate true "User Data"
// @Success 200 {object} delivery.JSONSuccessResult{}
// @Success 400 {object} delivery.JSONBadReqResult{}
// @Success 500 {object} delivery.JSONInternalResult{}
// @Router /v1/user/change [post]
func (d *UserDelivery) ChangePassword(c echo.Context) error {
	jwtGetID := middleware.GetUser(c)
	var user request.PasswordUpdate
	err := c.Bind(&user)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Failed to Bind Data", err)
	}
	err = c.Validate(&user)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "Failed, Wrong Input", err)
	}

	id := jwtGetID.ID
	log.Println(id)
	res, err := d.usecase.ChangePassword(id, user.OldPassword, user.NewPassword)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Internal Error", err)
	} else if res == "user not found" {
		return delivery.ErrorResponse(c, http.StatusNoContent, "User not Found", nil)
	}

	return delivery.SuccessResponse(c, res)
}

// GetDetail godoc
// @Summary Get Detail User
// @Description Get Detail User Data.
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} delivery.JSONSuccessResult{}
// @Success 400 {object} delivery.JSONBadReqResult{}
// @Success 500 {object} delivery.JSONInternalResult{}
// @Router /v1/user [get]
func (d *UserDelivery) GetDetail(c echo.Context) error {
	jwtGetID := middleware.GetUser(c)
	resp, err := d.usecase.GetCurrentUser(jwtGetID.ID)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Internal Error", err)
	}
	return delivery.SuccessResponse(c, response.FromDomain(resp))
}

func (d *UserDelivery) JWTTEST(c echo.Context) error {
	jwtGetID := middleware.GetUser(c)
	log.Println(jwtGetID.Role)
	log.Println(jwtGetID.ID)
	log.Println(jwtGetID.IsVerified)
	return delivery.SuccessResponse(c, jwtGetID.ID)
}
