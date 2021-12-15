package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ppob-service/delivery"
	"ppob-service/delivery/product/response"
	"ppob-service/usecase/product"
	"strconv"
)

type ProductDelivery struct {
	usecase product.IProductUsecase
}

func NewProductDelivery(uc product.IProductUsecase) *ProductDelivery {
	return &ProductDelivery{
		usecase: uc,
	}
}

func (p *ProductDelivery) GetTagihanPLN(c echo.Context) error {
	resp, err := p.usecase.GetTagihanPLN()
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "", err)
	}
	return delivery.SuccessResponse(c, response.FromDomain(resp))
}

func (p *ProductDelivery) GetProduct(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Wrong Params", err)
	}
	resp, err := p.usecase.GetProduct(idParam)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "", err)
	}
	return delivery.SuccessResponse(c, response.FromDomainList(resp))
}
