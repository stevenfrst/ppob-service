package delivery

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"ppob-service/delivery"
	"ppob-service/delivery/voucher/request"
	"ppob-service/delivery/voucher/response"
	"ppob-service/helpers/errorHelper"
	vouchers "ppob-service/usecase/voucher"
	"strconv"
)

type VoucherDelivery struct {
	usecase vouchers.IVoucherUseCase
}

func NewProductDelivery(uc vouchers.IVoucherUseCase) *VoucherDelivery {
	return &VoucherDelivery{usecase: uc}
}

func (d *VoucherDelivery) CreteVoucher(c echo.Context) error {
	var deliveryModel request.Voucher
	if err := c.Bind(&deliveryModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Semisal Butuh Validate

	//if err := c.Validate(&deliveryModel); err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	//}
	model := deliveryModel.ToDomain()
	err := d.usecase.Create(model)
	if errors.As(err, &errorHelper.DuplicateVoucher) {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "duplicate", err)
	} else if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "errors", err)
	}
	return delivery.SuccessResponse(c, "success")
}

func (d *VoucherDelivery) ReadByID(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Wrong Params", err)
	}
	resp, err := d.usecase.ReadById(idParam)
	if errors.As(err, &errorHelper.ErrRecordNotFound) {
		return delivery.ErrorResponse(c, http.StatusNoContent, "not found", err)
	} else if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "errors", err)
	}
	return delivery.SuccessResponse(c, response.FromDomain(resp))
}

func (d *VoucherDelivery) ReadAll(c echo.Context) error {
	resp, err := d.usecase.ReadALL()
	if errors.As(err, &errorHelper.ErrRecordNotFound) {
		return delivery.ErrorResponse(c, http.StatusNoContent, "not found", err)
	} else if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "errors", err)
	}
	return delivery.SuccessResponse(c, response.FromDomainList(resp))
}

func (d *VoucherDelivery) DeleteByID(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Wrong Params", err)
	}
	err = d.usecase.DeleteByID(idParam)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "errors", err)
	}
	return delivery.SuccessResponse(c, "success")
}

func (d *VoucherDelivery) Verify(c echo.Context) error {
	voucher := c.Param("voucher")
	err := d.usecase.Verify(voucher)
	if errors.As(err, &errorHelper.ErrVoucherNotMatch) {
		return delivery.ErrorResponse(c,http.StatusBadRequest,"",err)
	} else if err != nil {
		return delivery.ErrorResponse(c,http.StatusInternalServerError,"internal errors",err)
	}
	return delivery.SuccessResponse(c,"voucher match")
}
