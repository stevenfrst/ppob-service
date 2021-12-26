package routes

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"ppob-service/delivery"
	_ "ppob-service/docs"
)
import userDelivery "ppob-service/delivery/user"
import productDelivery "ppob-service/delivery/product"
import _middleware "ppob-service/app/middleware"

type RouteControllerList struct {
	UserDelivery    userDelivery.UserDelivery
	ProductDelivery productDelivery.ProductDelivery
	JWTConfig       middleware.JWTConfig
}

func (d RouteControllerList) RouteRegister(c *echo.Echo) {
	jwt := middleware.JWTWithConfig(d.JWTConfig)
	c.GET("/swagger/*", echoSwagger.WrapHandler)
	//User
	c.POST("/v1/login", d.UserDelivery.Login)
	c.POST("/v1/register", d.UserDelivery.Register)
	c.POST("/v1/user/change", d.UserDelivery.ChangePassword, jwt, RoleValidationUser())
	c.GET("/v1/user", d.UserDelivery.GetDetail, jwt, RoleValidationUser())

	//Product
	c.POST("/v1/product", d.ProductDelivery.CreateProduct, jwt, RoleValidationAdmin())
	c.GET("/v1/product/pln", d.ProductDelivery.GetTagihanPLN)
	c.GET("/v1/product/:id", d.ProductDelivery.GetProduct)
	c.PUT("/v1/product", d.ProductDelivery.EditProduct, jwt, RoleValidationAdmin())
	c.DELETE("/v1/product/:id", d.ProductDelivery.DeleteProduct, jwt, RoleValidationAdmin())
	c.GET("/v1/product/all", d.ProductDelivery.GetAllProducts)

	// Categories
	c.GET("/v1/category", d.ProductDelivery.GetCategory)
	c.GET("/v1/subcategory", d.ProductDelivery.GetSubCategory)
	c.PUT("/v1/subcategory", d.ProductDelivery.EditSubCategory, jwt, RoleValidationAdmin())
	c.POST("/v1/category", d.ProductDelivery.CreateCategory, jwt, RoleValidationAdmin())

	c.GET("/v1/bestseller/:id", d.ProductDelivery.GetBestSellerCategory)

	c.GET("/test", d.UserDelivery.JWTTEST, jwt, RoleValidationUser())

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

func RoleValidationAdmin() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := _middleware.GetUser(c)

			if claims.Role == "admin" {
				return hf(c)
			} else {
				return delivery.ErrorResponse(c, http.StatusForbidden, "", errors.New("StatusUnauthorized"))
			}
		}
	}
}
