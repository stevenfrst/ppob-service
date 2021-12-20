package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseReponse struct {
	Meta struct {
		Status   int      `json:"status"`
		Message  string   `json:"message"`
		Messages []string `json:"messages"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

// Swagger documentation
type JSONSuccessResult struct {
	Meta struct {
		Status   int      `json:"status" example:"200"`
		Message  string   `json:"message" example:"success"`
		Messages []string `json:"messages"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

type JSONBadReqResult struct {
	Meta struct {
		Status   int      `json:"status" example:"400"`
		Message  string   `json:"message" example:"failed"`
		Messages []string `json:"messages"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

type JSONInternalResult struct {
	Meta struct {
		Status   int      `json:"status" example:"500"`
		Message  string   `json:"message" example:"error database"`
		Messages []string `json:"messages"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}


func SuccessResponse(c echo.Context, data interface{}) error {
	response := BaseReponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "success"
	response.Data = data
	return c.JSON(http.StatusOK, response)
}

func ErrorResponse(c echo.Context, status int, err string, errs error) error {
	response := BaseReponse{}
	response.Meta.Status = status
	response.Meta.Messages = []string{errs.Error()}
	response.Meta.Message = err
	return c.JSON(response.Meta.Status, response)
}
