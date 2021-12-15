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
	res, err := d.usecase.ChangePassword(id, user.OldPassword, user.NewPassword)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Internal Error", err)
	} else if res == "user not found" {
		return delivery.ErrorResponse(c, http.StatusNoContent, "User not Found", nil)
	}

	return delivery.SuccessResponse(c, res)
}

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
	return delivery.SuccessResponse(c, jwtGetID.Role)
}

