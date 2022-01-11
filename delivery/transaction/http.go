package delivery

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"ppob-service/app/middleware"
	"ppob-service/delivery"
	"ppob-service/delivery/transaction/request"
	"ppob-service/delivery/transaction/response"
	"ppob-service/helpers/errorHelper"
	"ppob-service/usecase/transaction"
	"strconv"
)

type TransactionDelivery struct {
	usecase transaction.ITransactionUsecase
}

func NewTransactionDelivery(usecase transaction.ITransactionUsecase) *TransactionDelivery {
	return &TransactionDelivery{usecase: usecase}
}

func (t *TransactionDelivery) CreatePayment(c echo.Context) error {
	var deliveryModel request.CreatePayment
	if err := c.Bind(&deliveryModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	jwtGetID := middleware.GetUser(c)
	resp, err := t.usecase.GetVirtualAccount(jwtGetID.ID, deliveryModel.ToDomainPayment())
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "error", err)
	}
	return delivery.SuccessResponse(c, resp)
}

func (t *TransactionDelivery) GetNotification(c echo.Context) error {
	var deliveryModel request.NotificationReq
	if err := c.Bind(&deliveryModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := t.usecase.ProcessNotification(deliveryModel.ToDomainNotification())
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "internal errors", err)
	}
	return delivery.SuccessResponse(c, "success")

}

func (t *TransactionDelivery) GetTxUser(c echo.Context) error {
	jwtGetID := middleware.GetUser(c)
	resp, err := t.usecase.GetAllTxUser(jwtGetID.ID)
	if errors.As(err, &errorHelper.ErrRecordNotFound) {
		return delivery.ErrorResponse(c, http.StatusNoContent, "", err)
	} else if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return delivery.SuccessResponse(c, response.FromDomainList(resp))
}

func (t *TransactionDelivery) GetTxByID(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Wrong Params", err)
	}
	resp, err := t.usecase.GetTxByID(idParam)
	if errors.As(err, &errorHelper.ErrRecordNotFound) {
		return delivery.ErrorResponse(c, http.StatusNoContent, "no content", err)
	} else if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "", err)
	}
	return delivery.SuccessResponse(c, response.FromDomain(resp))

}
