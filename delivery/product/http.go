package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ppob-service/delivery"
	"ppob-service/delivery/product/request"
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

// GetTagihanPLN godoc
// @Summary Get Random PLN
// @Description Get Random Tagihan PLN
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} delivery.JSONSuccessResult{}
// @Success 400 {object} delivery.JSONBadReqResult{}
// @Success 500 {object} delivery.JSONInternalResult{}
// @Router /v1/product/pln [get]
func (p *ProductDelivery) GetTagihanPLN(c echo.Context) error {
	resp, err := p.usecase.GetTagihanPLN()
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "", err)
	}
	return delivery.SuccessResponse(c, response.FromDomain(resp))
}

// GetProduct godoc
// @Summary GetProduct via Params
// @Description Get Product via Param
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int64 true "ID Category"
// @Success 200 {object} delivery.JSONSuccessResult{}
// @Success 400 {object} delivery.JSONBadReqResult{}
// @Success 500 {object} delivery.JSONInternalResult{}
// @Router /v1/product/{id} [get]
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

// EditProduct godoc
// @Summary Edit Product
// @Description Edit Product via JSON
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param  user body request.EditProduct true "User Data"
// @Success 200 {object} delivery.JSONSuccessResult{}
// @Success 400 {object} delivery.JSONBadReqResult{}
// @Success 500 {object} delivery.JSONInternalResult{}
// @Router /v1/product [put]
func (p *ProductDelivery) EditProduct(c echo.Context) error {
	var productReq request.EditProduct
	err := c.Bind(&productReq)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Failed to Bind Data", err)
	}
	err = c.Validate(&productReq)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "Failed, Wrong Input", err)
	}
	err = p.usecase.EditProduct(productReq.ToDomain())
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Failed to Edit Product", err)
	}
	return delivery.SuccessResponse(c, "success")
}

// DeleteProduct godoc
// @Summary Delete Product
// @Description Delete Product via ID
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID Product"
// @Success 200 {object} delivery.JSONSuccessResult{}
// @Success 400 {object} delivery.JSONBadReqResult{}
// @Success 500 {object} delivery.JSONInternalResult{}
// @Router /v1/product/{id} [delete]
func (p *ProductDelivery) DeleteProduct(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "Wrong Params", err)
	}
	err = p.usecase.Delete(idParam)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Internal error", err)
	}

	return delivery.SuccessResponse(c, "success")
}

// GetBestSellerCategory godoc
// @Summary Get Best Seller
// @Description Get Best Seller each Category
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id category"
// @Success 200 {object} delivery.JSONSuccessResult{}
// @Success 400 {object} delivery.JSONBadReqResult{}
// @Success 500 {object} delivery.JSONInternalResult{}
// @Router /v1/bestseller/{id} [get]
func (p *ProductDelivery) GetBestSellerCategory(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "Wrong Params", err)
	}
	resp, err := p.usecase.GetBestSellerCategory(idParam)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Internal Error", err)
	}
	return delivery.SuccessResponse(c, response.FromDomainList(resp))
}
