package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ppob-service/app/middleware"
	"ppob-service/delivery"
	"ppob-service/delivery/transaction/request"
	"ppob-service/usecase/transaction"
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
		return delivery.ErrorResponse(c,http.StatusInternalServerError,"internal errors",err)
	}
	return delivery.SuccessResponse(c,"success")

}