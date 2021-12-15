package routes

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"ppob-service/delivery"
)
import userDelivery "ppob-service/delivery/user"
import productDelivery "ppob-service/delivery/product"
import _middleware "ppob-service/app/middleware"

type RouteControllerList struct {
	UserDelivery userDelivery.UserDelivery
	ProductDelivery productDelivery.ProductDelivery
	JWTConfig    middleware.JWTConfig
}

func (d RouteControllerList) RouteRegister(c *echo.Echo) {
	jwt := middleware.JWTWithConfig(d.JWTConfig)


	//User
	c.POST("/v1/login", d.UserDelivery.Login)
	c.POST("/v1/register", d.UserDelivery.Register)
	c.POST("/v1/user/change",d.UserDelivery.ChangePassword,jwt,RoleValidationUser())
	c.GET("/v1/user",d.UserDelivery.GetDetail,jwt,RoleValidationUser())

	//Product
	c.GET("/v1/product/pln",d.ProductDelivery.GetTagihanPLN)

	c.GET("/test", d.UserDelivery.JWTTEST, jwt,RoleValidationUser())

}

func RoleValidationUser() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := _middleware.GetUser(c)

			if claims.Role == "user" {
				return hf(c)
			} else {
				return delivery.ErrorResponse(c, http.StatusForbidden, "", errors.New("StatusUnauthorized"))
			}
		}
	}
}
